# go-fltk fork -- Maintenance Rules

This repo is `mkke/go-fltk`, a fork of `github.com/pwiecz/go-fltk` (the Go
bindings for FLTK). It carries local bindings and fixes the mallorn FLTK
frontends need, plus macOS static libraries built on this machine.

**There is no upstreaming.** Local changes stay local. The fork is therefore
maintained as a small stack of commits rebased onto the latest upstream
release, so that adopting a new upstream version is mechanical.

## The prime directive: keep the stack rebasable

Every local change is judged by what it costs at the next rebase:

1. **Prefer adding a file over modifying one.** A file upstream does not have
   can never conflict. A modified vendored file conflicts on every touch.
2. **Never modify a shared vendored tree.** `include/FL/` must stay
   byte-identical to upstream. See the per-platform include rule below.
3. **Append rather than interleave.** New wrapper functions go at the end of
   the relevant `.cxx` / `.h` section, not woven into existing code.
4. **No incidental whitespace or line-ending changes.** They are invisible in
   review and pure conflict fodder at rebase time.
5. **Only what is necessary.** No debug scaffolding, no unused bindings. If
   no consumer calls it, it does not belong here.

Measure before and after any sizeable change:

```sh
git diff --diff-filter=M --name-only upstream/main main | wc -l   # conflict-capable
git diff --diff-filter=A --name-only upstream/main main | wc -l   # conflict-free
```

As of the 2026-08-03 restack: 26 modified text files, 7 modified binary
archives, 181 added files. Source diff is `+463 / -10`, and only 8 of those
deleted lines are genuine modifications.

## Line endings: `.gitattributes` is load-bearing

`core.autocrlf=input` is set in the user's global git config. It rewrites CRLF
to LF in every file git stages, which silently turns a small edit into a
whole-file diff against upstream. It once inflated a header refresh to ~60000
changed lines when the real delta was ~1500.

`.gitattributes` sets `* -text` to disable all conversion for this repo. **Do
not remove it**, and do not "fix" the stray CRLF lines that exist in a handful
of upstream files (`table.h`, `widget.cxx`, `widget.h`) -- matching upstream's
bytes is the point.

## Per-platform include trees

FLTK headers must match the archive they are compiled against. Two layouts
exist in this repo:

| Platform | Headers | FLTK version |
| --- | --- | --- |
| `darwin/arm64` | `include/darwin/arm64/FL/` | 1.4.5 |
| `linux/amd64` | `include/linux/amd64/FL/` | 1.4.5 (upstream) |
| `windows/amd64` | `include/windows/amd64/FL/` | 1.4.5 (upstream) |
| `darwin/amd64`, `linux/arm`, `linux/arm64`, `openbsd/*` | shared `include/FL/` | 1.4.0 |

The per-platform tree is what `fltk-build.go` already generates
(`CMAKE_INSTALL_INCLUDEDIR=include/$GOOS/$GOARCH`), and it holds `FL/*.H`
plus `FL/fl_config.h` (and `FL/images/` where the platform needs it).

**Rule: when a platform's archives are rebuilt, give that platform its own
include tree.** Never upgrade shared `include/FL/` in place. Doing so breaks
every other platform still on the older archives -- they compile newer
declarations against older objects, which surfaces as link errors such as the
6-vs-7 argument `Fl::set_boxtype` and 8-vs-9 argument `fl_draw`.

Quick consistency check (shared must read 0, darwin/arm64 must read 5):

```sh
grep -n 'define FL_PATCH_VERSION' include/FL/Enumerations.H include/darwin/arm64/FL/Enumerations.H
```

## Rebuilding the macOS libraries

```sh
go run fltk-build.go
```

`fltk-build.go` clones FLTK at `tags/release-1.4.5`, applies
`lib/fltk-1.4.patch`, pins `CMAKE_OSX_DEPLOYMENT_TARGET=13.0`, and installs
into `lib/$GOOS/$GOARCH` and `include/$GOOS/$GOARCH`. The working checkout
lands in the gitignored `fltk_build/`.

`lib/fltk-1.4.patch` patches FLTK itself. It currently carries an upstream
win32 cleanup hunk and a cocoa hunk that calls `-finishLaunching` when `NSApp`
already exists, without which `kAEQuitApplication` (dock Quit, `osascript`
quit) is silently dropped.

Bumping the FLTK version means: change `commit` in `fltk-build.go`, rebuild,
refresh that platform's include tree, and re-check the patch still applies.

## Rebasing onto a new upstream release

```sh
git remote add upstream https://github.com/pwiecz/go-fltk.git   # once
git fetch upstream
git log --oneline upstream/main..main    # the local stack
```

The stack is ordered so conflicts stay isolated: pure-Go and wrapper source
first, then darwin build tooling, then the darwin include tree, then the
binary archives last. Binary conflicts are resolved by taking our side
wholesale.

Before keeping a local commit, check whether upstream has superseded it. The
2026-08-03 restack dropped the local Windows library rebuild for exactly this
reason: upstream shipped 1.4.5 Windows archives *and* headers, so keeping the
local 1.4.4 archives would have recreated the header/archive mismatch the
per-platform rule exists to prevent.

Verify a rebase with a real link, not just a compile -- `go build ./...` on a
library does not exercise the archives:

```sh
go build . && go test .
cd ../gui && go build ./... && go build -o /tmp/tt.bin ./themetest
```

Only `darwin/arm64` is routinely verified here. Linux and Windows are not
built on this machine; say so rather than implying they were checked.

## Consumers

Re-pin after any published change (`git.mallorn.de` repos, `master` branch):

- Pinned by pseudo-version in `go.mod`: **gui**, **mchart**, **mhelp**,
  **hilmir**.
- Local path `replace => ../go-fltk`, so they follow the checkout with no
  pin: **toolbed**, **tradewatch**.

```sh
go mod edit -replace github.com/pwiecz/go-fltk=github.com/mkke/go-fltk@<pseudo-version>
```

Check the `go` directive did not move afterwards -- the shared pre-commit hook
runs `go mod tidy` pinned to the existing directive, but verify anyway.

**gui and mchart compile FLTK headers straight out of this checkout**, via
relative cgo paths like `-I${SRCDIR}/../../go-fltk/include/darwin/arm64`. A
change to the include layout here breaks their build even when nothing about
the Go API changed, and it breaks it in the *working tree*, independently of
which version their `go.mod` pins. Always rebuild gui after touching
`include/`.

Publishing the fork is a force-push (`git push --force-with-lease origin
main`), because the rebase rewrites history that `origin/main` already had.

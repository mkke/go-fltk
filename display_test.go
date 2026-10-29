package fltk

import (
	"os"
	"os/exec"
	"testing"
)

// fontBeforeDisplayEnv names the case a run of TestFontBeforeAnyWindow
// runs.
const fontBeforeDisplayEnv = "GO_FLTK_FONT_BEFORE_DISPLAY"

// firstCalls are what a widget does with the draw font while it is built,
// before any window is shown. A call that reaches a NULL driver ends the
// process, which fails the subtest that ran it.
var firstCalls = map[string]func(t *testing.T){
	// No font is set yet, so the size DrawFont reports differs between
	// drivers. Returning at all is the check.
	"DrawFont": func(t *testing.T) { DrawFont() },
	// The sequence fltkplus measures a label with: keep the face, set the
	// size, measure.
	"MeasureInTheCurrentFace": func(t *testing.T) {
		face, _ := DrawFont()
		SetDrawFont(face, 13)
		if w, h := MeasureText("Reset", false); w <= 0 || h <= 0 {
			t.Errorf("MeasureText(%q) = %d, %d, want both above 0", "Reset", w, h)
		}
	},
}

// TestFontBeforeAnyWindow makes each of firstCalls the first FLTK call of
// a process. FLTK's graphics driver is NULL until the display is open, and
// Windows opens it only for the first window. Each case runs in a fresh
// run of the test binary, because an earlier test in this one may already
// have opened the display.
func TestFontBeforeAnyWindow(t *testing.T) {
	if name := os.Getenv(fontBeforeDisplayEnv); name != "" {
		firstCalls[name](t)
		return
	}
	for name := range firstCalls {
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestFontBeforeAnyWindow$", "-test.count=1") // #nosec G204 -- the test binary itself
			cmd.Env = append(os.Environ(), fontBeforeDisplayEnv+"="+name)
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("%s before any window: %v\n%s", name, err, out)
			}
		})
	}
}

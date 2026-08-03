package fltk

import "testing"

// TestParent_TopLevelWindowReturnsNil pins the nil guard in
// (*widget).Parent(). A top-level window has no FLTK parent, so Parent()
// must return a nil *Group. Before the guard it returned a non-nil
// *Group wrapping a NULL pointer, and any method call on that group
// (e.g. Label()) panicked with "widget is destroyed". Mirrors the
// long-standing nil handling in Window().
func TestParent_TopLevelWindowReturnsNil(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	if p := win.Parent(); p != nil {
		t.Errorf("top-level window Parent() = %v, want nil", p)
	}
}

// TestParent_ChildReturnsContainingGroup keeps the non-nil path honest:
// a widget added to a group reports a usable, non-nil parent.
func TestParent_ChildReturnsContainingGroup(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	NewGroup(0, 0, 200, 200)
	b := NewButton(0, 0, 50, 50, "x")
	win.End()
	if p := b.Parent(); p == nil {
		t.Fatal("child widget Parent() = nil, want the containing group")
	}
}

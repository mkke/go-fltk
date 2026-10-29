package fltk

import "testing"

func TestIsDestroyed_FalseWhileTheWidgetExists(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	b := NewButton(0, 0, 50, 50, "x")
	win.End()
	if win.IsDestroyed() || b.IsDestroyed() {
		t.Errorf("IsDestroyed() = %v (window), %v (button), want false for both", win.IsDestroyed(), b.IsDestroyed())
	}
}

// Destroy schedules the deletion. The widget is deleted, and IsDestroyed
// turns true, when the event loop runs.
func TestIsDestroyed_TrueOnceTheEventLoopDeletesTheWidget(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	b := NewButton(0, 0, 50, 50, "x")
	win.End()
	b.Destroy()
	if b.IsDestroyed() {
		t.Error("IsDestroyed() = true before the event loop ran, want false")
	}
	Check()
	if !b.IsDestroyed() {
		t.Error("IsDestroyed() = false after the event loop ran, want true")
	}
}

// A child is deleted with its window, without a Destroy call of its own.
func TestIsDestroyed_TrueAfterTheWindowIsDeleted(t *testing.T) {
	win := NewWindow(200, 200)
	g := NewGroup(0, 0, 200, 200)
	b := NewButton(0, 0, 50, 50, "x")
	g.End()
	win.End()
	win.Destroy()
	Check()
	if !g.IsDestroyed() || !b.IsDestroyed() {
		t.Errorf("IsDestroyed() = %v (group), %v (button), want true for both", g.IsDestroyed(), b.IsDestroyed())
	}
}

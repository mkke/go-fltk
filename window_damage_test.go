package fltk

import "testing"

// TestDamageArea_DrawsTheWindowBoxUnderTheRectangle shows a window whose
// box counts its draws, with a child that has no box. A redraw of the
// child does not draw the window's box. DamageArea over the child's
// rectangle does.
func TestDamageArea_DrawsTheWindowBoxUnderTheRectangle(t *testing.T) {
	boxDraws := 0
	SetBoxType(FREE_BOXTYPE, func(x, y, w, h int, c Color) {
		boxDraws++
		SetDrawColor(c)
		DrawRectf(x, y, w, h)
	})
	win := NewWindow(200, 100)
	win.SetBox(FREE_BOXTYPE)
	child := NewBox(NO_BOX, 10, 10, 60, 30)
	win.End()
	defer win.Destroy()

	// pump runs the event loop until boxDraws passes from, or for one
	// second, and reports whether it did.
	pump := func(from int) bool {
		for i := 0; i < 50 && boxDraws <= from; i++ {
			Wait(0.02)
		}
		return boxDraws > from
	}

	win.Show()
	if !pump(0) {
		t.Fatal("the window's box was not drawn after Show")
	}
	shown := boxDraws

	child.Redraw()
	if pump(shown) {
		t.Errorf("a redraw of a child without a box drew the window's box %d more times, want 0", boxDraws-shown)
	}
	redrawn := boxDraws

	win.DamageArea(child.X(), child.Y(), child.W(), child.H())
	if !pump(redrawn) {
		t.Error("DamageArea over the child did not draw the window's box")
	}
}

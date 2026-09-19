package fltk

import "testing"

// The getters below exist so a caller can read back what it set. A
// conformance checker that drives an application over an automation
// socket is the case: it reports a widget's tooltip and the size the
// widget draws its own text at, and neither had a getter.

// TestTooltip_ReadsBackWhatWasSet covers the copy_tooltip path.
// Fl_Widget::tooltip() returns the widget's own text and never inherits
// a parent group's, so a widget with no tooltip answers empty.
func TestTooltip_ReadsBackWhatWasSet(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	b := NewButton(0, 0, 50, 50, "x")
	win.End()

	if got := b.Tooltip(); got != "" {
		t.Errorf("a widget with no tooltip reports %q, want the empty string", got)
	}
	b.SetTooltip("Copy the value")
	if got := b.Tooltip(); got != "Copy the value" {
		t.Errorf("Tooltip() = %q, want %q", got, "Copy the value")
	}
}

// TestInputTextSize_IsTheValuesSizeNotTheLabels keeps the two sizes
// apart. An input draws its value at textsize() and its own label
// beside the box at labelsize().
func TestInputTextSize_IsTheValuesSizeNotTheLabels(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	in := NewInput(0, 0, 100, 25)
	win.End()

	in.SetLabelSize(11)
	if got := in.TextSize(); got == 11 {
		t.Errorf("TextSize() = %d, which is the label size rather than the value's", got)
	}
	if got := in.TextSize(); got <= 0 {
		t.Errorf("TextSize() = %d, want the point size the value is drawn at", got)
	}
}

// TestMenuTextSize_AnswersOnAChoice covers the Fl_Menu_ half of the
// same distinction, through the type every menu widget embeds.
func TestMenuTextSize_AnswersOnAChoice(t *testing.T) {
	win := NewWindow(200, 200)
	defer win.Destroy()
	c := NewChoice(0, 0, 100, 25)
	win.End()

	if got := c.TextSize(); got <= 0 {
		t.Errorf("TextSize() = %d, want the point size the item titles are drawn at", got)
	}
}

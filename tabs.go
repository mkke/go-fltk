package fltk

/*
#include "tabs.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text ...string) *Tabs {
	t := &Tabs{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_Tabs(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (t *Tabs) Value() int {
	return int(C.go_fltk_Tabs_value((*C.Fl_Tabs)(t.ptr())))
}

func (t *Tabs) SetValue(value int) {
	C.go_fltk_Tabs_set_value((*C.Fl_Tabs)(t.ptr()), (C.int)(value))
}

type Overflow int

var (
	OverflowCompress = Overflow(C.go_FL_OVERFLOW_COMPRESS)
	OverflowClip     = Overflow(C.go_FL_OVERFLOW_CLIP)
	OverflowPulldown = Overflow(C.go_FL_OVERFLOW_PULLDOWN)
	OverflowDrag     = Overflow(C.go_FL_OVERFLOW_DRAG)
)

func (t *Tabs) SetOverflow(overflow Overflow) {
	C.go_fltk_Tabs_handle_overflow((*C.Fl_Tabs)(t.ptr()), (C.int)(overflow))
}

// Which returns the index of the child whose header tab lies under the
// given event coordinates (window-relative, i.e. fltk.EventX/EventY),
// or -1 when the point is outside the tab bar. Mirrors
// Fl_Tabs::which(). Note that Fl_Tabs only *selects* a tab on
// FL_RELEASE; during a press-and-drag the pressed tab is identified by
// this hit-test, not by Value().
func (t *Tabs) Which(eventX, eventY int) int {
	return int(C.go_fltk_Tabs_which((*C.Fl_Tabs)(t.ptr()), C.int(eventX), C.int(eventY)))
}

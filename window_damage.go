package fltk

/*
#include "window_damage.h"
*/
import "C"

// DamageArea marks a rectangle of the window for a full redraw. x and y
// are in the window's coordinates, which are the coordinates of the
// widgets inside it. The next flush draws the window's box and every
// child that intersects the rectangle, clipped to it.
//
// A widget without a box draws nothing under itself, so a redraw of it
// leaves the pixels below as they were. Damaging its rectangle on its
// window repaints them, which is what Fl_Button does for a button
// without a box. Fl_Widget::damage with a rectangle does this only on a
// window: on any other widget it marks the window FL_DAMAGE_CHILD, and
// the window's box is not drawn.
//
// A window that is not shown ignores the call.
func (w *Window) DamageArea(x, y, width, height int) {
	C.go_fltk_Window_damage_area((*C.Fl_Window)(w.ptr()), C.int(x), C.int(y), C.int(width), C.int(height))
}

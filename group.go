package fltk

/*
#include "group.h"
*/
import "C"
import (
	"unsafe"
)

type Group struct {
	widget
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initWidget(g, unsafe.Pointer(C.go_fltk_new_Group(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
}

func (g *Group) Begin() {
	C.go_fltk_Group_begin((*C.Fl_Group)(g.ptr()))
}
func (g *Group) End() {
	C.go_fltk_Group_end((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Add(w Widget) {
	C.go_fltk_Group_add((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}
func (g *Group) Remove(w Widget) {
	C.go_fltk_Group_remove((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}

func (g *Group) Resizable(w Widget) {
	C.go_fltk_Group_resizable((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}
func (g *Group) DrawChildren() {
	C.go_fltk_Group_draw_children((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Child(index int) *widget {
	child := C.go_fltk_Group_child((*C.Fl_Group)(g.ptr()), C.int(index))
	if child == nil {
		return nil
	}
	widget := &widget{}
	initUnownedWidget(widget, unsafe.Pointer(child))
	return widget
}

func (g *Group) ChildCount() int {
	return int(C.go_fltk_Group_child_count((*C.Fl_Group)(g.ptr())))
}

func (g *Group) Children() []*widget {
	childCount := g.ChildCount()
	children := make([]*widget, 0, childCount)
	for i := 0; i < childCount; i++ {
		children = append(children, g.Child(i))
	}
	return children
}

// CurrentGroup returns the group that newly-created widgets are implicitly
// added to, or nil if no group is currently open. This mirrors
// Fl_Group::current() and is useful for save/restore patterns around
// nested Begin/End pairs that would otherwise leave the current-group
// pointer in an unexpected state.
func CurrentGroup() *Group {
	p := unsafe.Pointer(C.go_fltk_Group_current())
	if p == nil {
		return nil
	}
	g := &Group{}
	initUnownedWidget(g, p)
	return g
}

// SetCurrentGroup sets the group that newly-created widgets are added to.
// Pass nil to clear (no implicit parent). Mirrors Fl_Group::current(g).
func SetCurrentGroup(g *Group) {
	var p *C.Fl_Group
	if g != nil {
		p = (*C.Fl_Group)(g.ptr())
	}
	C.go_fltk_Group_set_current(p)
}

package fltk

/*
#include <stdlib.h>
#include "platform.h"
*/
import "C"

var openCallback func(string)

func SetOpenCallback(newOpenCallback func(string)) {
	C.go_fltk_init_open_callback()
	openCallback = newOpenCallback
}

// openDisplay opens the display unless a graphics driver already exists.
// Windows opens it only for the first window shown. Until then, reading
// the draw font ends in an access violation. macOS opens it on the first
// screen query, so there a widget can read the font while it is built.
// Setting the font needs no call, because FLTK's fl_font(face, size)
// opens the display itself. Measuring with no font set is outside FLTK's
// contract on every platform.
func openDisplay() {
	C.go_fltk_open_display_once()
}

//export _go_openCallbackHandler
func _go_openCallbackHandler(path *C.char) {
	if openCallback != nil {
		openCallback(C.GoString(path))
	}
}

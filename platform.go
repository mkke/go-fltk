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

//export _go_openCallbackHandler
func _go_openCallbackHandler(path *C.char) {
	if openCallback != nil {
		openCallback(C.GoString(path))
	}
}

#include <FL/x.H>
#include <FL/fl_draw.H>

#include "_cgo_export.h"

void go_fltk_init_open_callback() {
  fl_open_callback((void (*)(const char *))(&_go_openCallbackHandler));
}

// The same guard fl_font(face, size) carries in src/fl_draw.cxx. The
// getters fl_font() and fl_size() go straight through fl_graphics_driver,
// which is NULL until the display is open.
void go_fltk_open_display_once() {
  if (!fl_graphics_driver)
    fl_open_display();
}
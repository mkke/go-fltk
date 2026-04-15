#include <FL/x.H>

#include "_cgo_export.h"

void go_fltk_init_open_callback() {
  fl_open_callback((void (*)(const char *))(&_go_openCallbackHandler));
}
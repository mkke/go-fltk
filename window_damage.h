#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Window Fl_Window;

  extern void go_fltk_Window_damage_area(Fl_Window* w, int x, int y, int width, int height);

#ifdef __cplusplus
}
#endif

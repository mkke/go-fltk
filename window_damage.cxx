#include "window_damage.h"

#include <FL/Enumerations.H>
#include <FL/Fl_Window.H>

void go_fltk_Window_damage_area(Fl_Window* w, int x, int y, int width, int height) {
  w->damage(FL_DAMAGE_ALL, x, y, width, height);
}

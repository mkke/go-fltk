package fltk

// TextAttr holds the further attributes of a text style, the attr field of
// Fl_Text_Display::Style_Table_Entry.
type TextAttr uint

// TEXT_ATTR_BGCOLOR draws the text of a style on the style's BgColor
// (Fl_Text_Display::ATTR_BGCOLOR).
const TEXT_ATTR_BGCOLOR TextAttr = 0x0001

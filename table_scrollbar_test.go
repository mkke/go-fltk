package fltk

import "testing"

// newOverflowingTable builds a table whose rows and columns both exceed
// its box, with a column header and a row header, so both scrollbars are
// shown and both have a header to clear. The window's box is flat and
// the table's is none, so the table's own x and y are the inner edges.
func newOverflowingTable(t *testing.T) (*Window, *TableRow) {
	t.Helper()
	win := NewWindow(400, 300)
	win.SetBox(FLAT_BOX)
	tbl := NewTableRow(10, 20, 300, 200)
	tbl.SetBox(NO_BOX)
	tbl.EnableColumnHeaders()
	tbl.SetColumnHeaderHeight(25)
	tbl.EnableRowHeaders()
	tbl.SetRowHeaderWidth(30)
	tbl.SetRowCount(100)
	tbl.SetRowHeightAll(20)
	tbl.SetColumnCount(8)
	tbl.SetColumnWidthAll(100)
	win.End()
	return win, tbl
}

// TestScrollbarBounds_StartWhereTheHeadersEnd pins the placement: the
// vertical scrollbar begins at the column header's lower edge and the
// horizontal one at the row header's right edge, and each ends where the
// other begins.
func TestScrollbarBounds_StartWhereTheHeadersEnd(t *testing.T) {
	win, tbl := newOverflowingTable(t)
	defer win.Destroy()

	vx, vy, vw, vh, vShown := tbl.VerticalScrollbarBounds()
	hx, hy, hw, hh, hShown := tbl.HorizontalScrollbarBounds()
	if !vShown || !hShown {
		t.Fatalf("scrollbars shown = %v, %v; want both, the content overflows on both axes", vShown, hShown)
	}
	ss := tbl.ScrollbarSize()
	if ss == 0 {
		ss = ScrollbarSize()
	}
	if vw != ss || hh != ss {
		t.Errorf("scrollbar thickness = %d, %d; want %d", vw, hh, ss)
	}
	if want := tbl.Y() + tbl.ColumnHeaderHeight(); vy != want {
		t.Errorf("vertical scrollbar y = %d; want %d, the column header's lower edge", vy, want)
	}
	if want := tbl.X() + tbl.RowHeaderWidth(); hx != want {
		t.Errorf("horizontal scrollbar x = %d; want %d, the row header's right edge", hx, want)
	}
	if vx != tbl.X()+tbl.W()-ss {
		t.Errorf("vertical scrollbar x = %d; want %d, the right edge", vx, tbl.X()+tbl.W()-ss)
	}
	if hy != tbl.Y()+tbl.H()-ss {
		t.Errorf("horizontal scrollbar y = %d; want %d, the bottom edge", hy, tbl.Y()+tbl.H()-ss)
	}
	if vy+vh != hy {
		t.Errorf("vertical scrollbar ends at %d; want %d, where the horizontal one begins", vy+vh, hy)
	}
	if hx+hw != vx {
		t.Errorf("horizontal scrollbar ends at %d; want %d, where the vertical one begins", hx+hw, vx)
	}
}

// TestScrollbarBounds_HiddenWhenTheContentFits keeps the visibility
// honest: a table whose content fits reports neither scrollbar shown.
func TestScrollbarBounds_HiddenWhenTheContentFits(t *testing.T) {
	win, tbl := newOverflowingTable(t)
	defer win.Destroy()
	tbl.SetRowCount(3)
	tbl.SetColumnCount(2)

	if _, _, _, _, shown := tbl.VerticalScrollbarBounds(); shown {
		t.Error("vertical scrollbar shown with three rows")
	}
	if _, _, _, _, shown := tbl.HorizontalScrollbarBounds(); shown {
		t.Error("horizontal scrollbar shown with two columns")
	}
}

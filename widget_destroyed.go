package fltk

// IsDestroyed reports whether the FLTK widget behind w has been deleted,
// with its own Destroy or with a group or window that held it. Destroy
// only schedules the deletion, which happens the next time the event loop
// runs, so IsDestroyed reports false until then and every call on w is
// still valid. After the deletion, every other method panics with
// ErrDestroyed, so code that can run after the widget is gone, such as a
// callback queued with Awake, asks this first.
func (w *widget) IsDestroyed() bool {
	return !w.exists()
}

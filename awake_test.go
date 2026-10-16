package fltk

import "testing"

func awakeMapLen() int {
	globalAwakeMap.mutex.Lock()
	defer globalAwakeMap.mutex.Unlock()
	return len(globalAwakeMap.awakeMap)
}

// TestAwake_RefusedCallbackIsNotKept pins that Awake drops the
// registration of a callback FLTK refuses. FLTK holds queued callbacks in
// a fixed ring and refuses one while the ring is full; a refused callback
// never runs, so a registration kept for it would never be released.
func TestAwake_RefusedCallbackIsNotKept(t *testing.T) {
	Lock()
	defer Unlock()

	filled := false
	for i := 0; i < 100000; i++ {
		if !Awake(func() {}) {
			filled = true
			break
		}
	}
	if !filled {
		t.Fatal("FLTK took 100000 callbacks without refusing one; the test cannot fill its ring")
	}

	before := awakeMapLen()
	if Awake(func() {}) {
		t.Fatal("Awake took a callback while FLTK's ring was full")
	}
	if after := awakeMapLen(); after != before {
		t.Errorf("a refused Awake left %d registrations, want %d", after, before)
	}
}

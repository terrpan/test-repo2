package notifycheck

import "testing"

// TestNotify fails on purpose to exercise Scrapyard's failure comment.
func TestNotify(t *testing.T) {
	t.Fatal("deliberate failure")
}

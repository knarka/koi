package koi

import (
	"github.com/nsf/termbox-go"
	"testing"
)

func TestNewKoi(t *testing.T) {
	k := NewKoi()

	if k.filename != "new file" {
		t.Error("incorrect filename")
	}

	if k.scroll != 0 {
		t.Errorf("incorrect scroll value. wanted: %d, got: %d", 0, k.scroll)
	}
}

func TestKeyPresses(t *testing.T) {
	k := NewKoi()

	if k.key(termbox.KeyCtrlC) {
		t.Error("key function incorrectly returned true on <C-c>")
	}

	k.key(termbox.KeyCtrlJ)

	if k.cursor.Y() != 0 {
		t.Error("cursor moved out of bounds")
	}

	k.buffer.AddLine("er was eens een jongen uit katwijk aan zee")
	k.buffer.AddLine("die stopte zijn limericks na regel twee")

	k.key(termbox.KeyCtrlJ)
	k.key(termbox.KeyCtrlJ)
	k.key(termbox.KeyCtrlK)

	if k.cursor.Y() != 1 {
		t.Errorf("incorrect vertical cursor movement. wanted: %d, got: %d", 1, k.cursor.Y())
	}
}

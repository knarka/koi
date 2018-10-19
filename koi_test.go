package main

import (
	"github.com/nsf/termbox-go"
	"testing"
)

func TestNewKoi(t *testing.T) {
	k := newKoi()

	if k.filename != "new file" {
		t.Error("incorrect filename")
	}

	if k.scroll != 0 {
		t.Errorf("incorrect scroll value. wanted: %d, got: %d", 0, k.scroll)
	}
}

func TestKeyPresses(t *testing.T) {
	k := newKoi()

	if k.key(termbox.KeyCtrlC) {
		t.Error("key function incorrectly returned true on <C-c>")
	}

	k.key(termbox.KeyCtrlJ)

	if k.cursor.y != 0 {
		t.Error("cursor moved out of bounds")
	}

	k.buffer.addLine("er was eens een jongen uit katwijk aan zee")
	k.buffer.addLine("die stopte zijn limericks na regel twee")

	k.key(termbox.KeyCtrlJ)
	k.key(termbox.KeyCtrlJ)
	k.key(termbox.KeyCtrlK)

	if k.cursor.y != 1 {
		t.Errorf("incorrect vertical cursor movement. wanted: %d, got: %d", 1, k.cursor.y)
	}
}

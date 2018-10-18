package main

import (
	"testing"
)

func TestNewCursor(t *testing.T) {
	c := newCursor()

	if c.x != 0 || c.y != 0 {
		t.Error("Cursor did not initialize at 0, 0.")
	}
}

func TestCursorMovement(t *testing.T) {
	c := newCursor()

	c.left()
	c.left()
	c.right()
	c.up()
	c.down()
	c.down()

	if c.x != -1 || c.y != 1 {
		t.Error("Cursor did not move correctly.")
	}
}

package cursor

import (
	"testing"
)

func TestNewCursor(t *testing.T) {
	c := NewCursor()

	if c.x != 0 || c.y != 0 {
		t.Error("cursor did not initialize at 0, 0")
	}
}

func TestCursorMovement(t *testing.T) {
	c := NewCursor()

	c.Left()
	c.Left()
	c.Right()
	c.Up()
	c.Down()
	c.Down()

	if c.x != -1 || c.y != 1 {
		t.Error("cursor did not move correctly")
	}
}

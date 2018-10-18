package main

import (
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

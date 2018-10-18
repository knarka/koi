package main

import (
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := newBuffer()

	if len(b.content) != 0 {
		t.Errorf("incorrect buffer length. wanted: %d, got: %d", 0, len(b.content))
	}

	if cap(b.content) != 1024 {
		t.Errorf("incorrect buffer cap. wanted: %d, got: %d", 1024, len(b.content))
	}
}

func TestAddLine(t *testing.T) {
	b := newBuffer()

	b.addLine("hello world")

	if len(b.content) != 1 || b.getLine(1) != "hello world" {
		t.Error("could not succesfully add a line.")
	}
}

func TestLines(t *testing.T) {
	b := newBuffer()

	if b.lines() != 0 {
		t.Error("buffer should have 0 lines right after creation")
	}

	b.addLine("hello world")
	b.addLine("привет мир")

	if b.lines() != 2 {
		t.Error("buffer should have 2 lines after adding 2 lines")
	}
}

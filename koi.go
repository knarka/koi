package main

import (
	"github.com/nsf/termbox-go"
)

type koi struct {
	filename string
	scroll   int
	buffer   *buffer
	cursor   *cursor
}

func newKoi() *koi {
	k := new(koi)

	k.filename = "new file"
	k.scroll = 0

	k.buffer = newBuffer()
	k.buffer.addLine("test")
	k.cursor = newCursor()

	return k
}

func (k *koi) handle(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		return k.key(ev.Key)
	}

	return true
}

func (k *koi) key(key termbox.Key) bool {
	switch key {
	case termbox.KeyCtrlC:
		return false
	case termbox.KeyCtrlK:
		k.cursorUp()
	case termbox.KeyCtrlJ:
		k.cursorDown()
	default:
		return true
	}

	return true
}

func (k *koi) cursorUp() {
	if k.cursor.y > 0 {
		k.cursor.up()
	}
}

func (k *koi) cursorDown() {
	if k.cursor.y < k.buffer.lines()-1 {
		k.cursor.down()
	}
}

func (k *koi) draw() {
	termbox.SetCursor(k.cursor.x, k.cursor.y)
	termbox.Flush()
}

func (k *koi) loop() {
	events := make(chan termbox.Event, 20)

	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-events:
			ok := k.handle(&ev)
			if !ok {
				return
			}
			k.draw()
		}
	}
}

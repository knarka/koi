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
	case termbox.KeyCtrlH:
		k.cursorLeft()
	case termbox.KeyCtrlL:
		k.cursorRight()
	}

	return true
}

func (k *koi) cursorUp() {
	if k.cursor.y > 0 {
		k.cursor.up()
	}
}

func (k *koi) cursorDown() {
	if k.cursor.y < k.buffer.lines()-1 { // TODO: compensate for scroll
		k.cursor.down()
	}
}

func (k *koi) cursorLeft() {
	if k.cursor.x > 0 {
		k.cursor.left()
	}
}

func (k *koi) cursorRight() {
	if k.cursor.x < len(k.buffer.getLine(k.cursor.y, k.scroll)) {
		k.cursor.right()
	}
}

func (k *koi) drawBuffer(x1, x2, y1, y2, xOff, yOff int) {
	for y := y1; y < y2 && y < k.buffer.lines(); y++ {
		line := k.buffer.getLine(y, k.scroll)
		for x := x1; x < x2 && x < len(line); x++ {
			termbox.SetCell(x+xOff, y+yOff, rune(line[x]), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (k *koi) draw() {
	width, height := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	k.drawBuffer(0, width, 0, height, 0, 0)

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

package koi

import (
	"github.com/nsf/termbox-go"

	"github.com/knarka/koi/buffer"
	"github.com/knarka/koi/cursor"
)

type Koi struct {
	filename string
	scroll   int
	buffer   *buffer.TextBuffer
	cursor   *cursor.Cursor
}

func NewKoi() *Koi {
	k := new(Koi)

	k.filename = "new file"
	k.scroll = 0

	k.buffer = buffer.NewBuffer()
	k.buffer.AddLine("test") // XXX
	k.cursor = cursor.NewCursor()

	return k
}

func (k *Koi) handle(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		return k.key(ev.Key)
	}

	return true
}

func (k *Koi) key(key termbox.Key) bool {
	switch key {
	case termbox.KeyCtrlC:
		return false

	case termbox.KeyCtrlK, termbox.KeyArrowUp:
		k.cursorUp()
	case termbox.KeyCtrlJ, termbox.KeyArrowDown:
		k.cursorDown()
	case termbox.KeyCtrlH, termbox.KeyArrowLeft:
		k.cursorLeft()
	case termbox.KeyCtrlL, termbox.KeyArrowRight:
		k.cursorRight()
	}

	return true
}

func (k *Koi) cursorUp() {
	if k.cursor.Y() > 0 {
		k.cursor.Up()
	}
}

func (k *Koi) cursorDown() {
	if k.cursor.Y() < k.buffer.Lines()-1 { // TODO: compensate for scroll
		k.cursor.Down()
	}
}

func (k *Koi) cursorLeft() {
	if k.cursor.X() > 0 {
		k.cursor.Left()
	}
}

func (k *Koi) cursorRight() {
	if k.cursor.X() < len(k.buffer.GetLine(k.cursor.Y(), k.scroll)) {
		k.cursor.Right()
	}
}

func (k *Koi) drawBuffer(x1, x2, y1, y2, xOff, yOff int) {
	for y := y1; y < y2 && y < k.buffer.Lines(); y++ {
		line := k.buffer.GetLine(y, k.scroll)
		for x := x1; x < x2 && x < len(line); x++ {
			termbox.SetCell(x+xOff, y+yOff, rune(line[x]), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (k *Koi) draw() {
	width, height := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	k.drawBuffer(0, width, 0, height, 0, 0)

	termbox.SetCursor(k.cursor.X(), k.cursor.Y())
	termbox.Flush()
}

func (k *Koi) Loop() {
	events := make(chan termbox.Event, 20)

	k.draw()

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

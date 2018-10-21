package buffer

type TextBuffer struct {
	content []string
}

func NewBuffer() *TextBuffer {
	b := new(TextBuffer)

	b.content = make([]string, 1024)
	b.content = b.content[:0]

	return b
}

func (b *TextBuffer) AddLine(line string) {
	b.content = b.content[:len(b.content)+1]
	b.content[len(b.content)-1] = line
}

func (b *TextBuffer) GetLine(i, scroll int) string {
	return b.content[i+scroll]
}

func (b *TextBuffer) Lines() int {
	return len(b.content)
}

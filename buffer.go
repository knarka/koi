package main

type buffer struct {
	content		[]string
}

func newBuffer() *buffer {
	b := new(buffer)

	b.content = make([]string, 1024)
	b.content = b.content[:0]

	return b
}

func (b *buffer) addLine(line string) {
	b.content = b.content[:len(b.content) + 1]
	b.content[len(b.content) - 1] = line
}

func (b *buffer) size() int {
	return len(b.content)
}

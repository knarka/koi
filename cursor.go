package main

type cursor struct {
	x	int
	y	int
}

func newCursor() *cursor {
	c := new(cursor)

	c.x = 0
	c.y = 0

	return c
}

func (c *cursor) left() {
	c.x -= 1
}

func (c *cursor) right() {
	c.x += 1
}

func (c *cursor) up() {
	c.y -= 1
}

func (c *cursor) down() {
	c.y += 1
}

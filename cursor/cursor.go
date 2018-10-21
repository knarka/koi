package cursor

type Cursor struct {
	x int
	y int
}

func NewCursor() *Cursor {
	c := new(Cursor)

	c.x = 0
	c.y = 0

	return c
}

func (c *Cursor) Left() {
	c.x--
}

func (c *Cursor) Right() {
	c.x++
}

func (c *Cursor) Up() {
	c.y--
}

func (c *Cursor) Down() {
	c.y++
}

func (c *Cursor) X() int {
	return c.x
}

func (c *Cursor) Y() int {
	return c.y
}

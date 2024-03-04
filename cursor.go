package main

import "sync"

type cursor struct {
	x  int
	y  int
	mu *sync.Mutex
}

func NewCursor(initX int, initY int) *cursor {
	return &cursor{
		x:  initX,
		y:  initY,
		mu: &sync.Mutex{},
	}
}

func (c *cursor) MoveLeft(maxWidth int) (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.x > 0 {
		c.x--
		return c.x, c.y
	}
	if c.y > 0 {
		c.y--
		c.x = maxWidth - 1
	}
	return c.x, c.y
}

func (c *cursor) MoveRight(maxWidth int, maxHeight int) (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.x < maxWidth-1 {
		c.x++
		return c.x, c.y
	}
	if c.y < maxHeight-1 {
		c.y++
		c.x = 0
		return c.x, c.y
	}
	return c.x, c.y
}

func (c *cursor) Enter() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x = 0
	c.y++
	return c.x, c.y
}

func (c *cursor) CurLocation() (int, int) {
	return c.x, c.y
}

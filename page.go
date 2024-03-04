package main

import "errors"

// TODO: Fix it so that it doesn't have to call the other methods to test a certain method.
// page manages the characters saved in the editor.
type page struct {
	buffer [][]rune
	width  int
	height int
	lastY  int
	lastX  int
}

func NewPage(width int, height int) *page {
	buffer := make([][]rune, height)
	for index := range buffer {
		buffer[index] = make([]rune, width)
	}
	return &page{
		buffer: buffer,
		width:  width,
		height: height,
	}
}

func (p *page) AddChar(x int, y int, ch rune) (err error) {
	if x < 0 {
		return errors.New("x is less than 0")
	}
	if y < 0 {
		return errors.New("y is less than 0")
	}
	if y >= p.height {
		return errors.New("y is greater than height")
	}
	if x >= p.width {
		return errors.New("x is greater than width")
	}

	xyDigit := p.xyToDigit(x, y)
	lastXYDigit := p.xyToDigit(p.lastX, p.lastY)

	if xyDigit == lastXYDigit {
		newX, newY := p.digitToXY(xyDigit + 1)
		oldCh := p.buffer[y][x]
		p.buffer[newY][newX] = oldCh
		p.buffer[y][x] = ch
		return nil
	}

	if xyDigit > lastXYDigit {
		p.lastX = x
		p.lastY = y
		p.buffer[y][x] = ch
		return nil
	}

	for digit := lastXYDigit; digit >= xyDigit; digit-- {
		curX, curY := p.digitToXY(digit)
		newX, newY := p.digitToXY(digit + 1)
		p.buffer[newY][newX] = p.buffer[curY][curX]
	}
	p.buffer[y][x] = ch

	return nil
}

func (p *page) RemoveChar(x int, y int) (err error) {
	if x < 0 {
		return errors.New("x is less than 0")
	}
	if y < 0 {
		return errors.New("y is less than 0")
	}
	if y >= p.height {
		return errors.New("y is greater than height")
	}
	if x >= p.width {
		return errors.New("x is greater than width")
	}

	xyDigit := p.xyToDigit(x, y)
	lastXYDigit := p.xyToDigit(p.lastX, p.lastY)

	if xyDigit == lastXYDigit {
		p.buffer[y][x] = 0
		p.lastX, p.lastY = p.digitToXY(xyDigit - 1)
		return nil
	}
	if xyDigit > lastXYDigit {
		return errors.New("not supposed to reach here")
	}

	for digit := xyDigit + 1; digit <= lastXYDigit; digit++ {
		curX, curY := p.digitToXY(digit)
		newX, newY := p.digitToXY(digit - 1)
		p.buffer[newY][newX] = p.buffer[curY][curX]
	}

	return nil
}

func (p *page) Buffer() [][]rune {
	buffer := make([][]rune, p.height)
	for index := range buffer {
		buffer[index] = make([]rune, p.width)
		copy(buffer[index], p.buffer[index])
	}
	return buffer
}

func (p *page) xyToDigit(x int, y int) int {
	return y*p.width + x
}

func (p *page) digitToXY(digit int) (int, int) {
	return digit % p.width, digit / p.width
}

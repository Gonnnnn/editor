package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChar(t *testing.T) {
	t.Run("Fails if y is less than 0", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(0, -1, 'a')
		assert.EqualError(t, err, "y is less than 0")
	})

	t.Run("Fails if x is less than 0", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(-1, 0, 'a')
		assert.EqualError(t, err, "x is less than 0")
	})

	t.Run("Fails if y is greater than height", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(0, 10, 'a')
		assert.EqualError(t, err, "y is greater than height")
	})

	t.Run("Fails if x is greater than width", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(10, 0, 'a')
		assert.EqualError(t, err, "x is greater than width")
	})

	t.Run("Adds a character at 0, 0 in the empty page", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(0, 0, 'a')
		assert.Nil(t, err)
		assert.Equal(t, 'a', p.buffer[0][0])
	})

	t.Run("Adds a character at 0, 0 in the non-empty page", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(9, 1, 'b')
		fmt.Printf("%v\n", p.buffer)
		assert.Nil(t, err)
		err = p.AddChar(0, 0, 'a')
		assert.Nil(t, err)
		assert.Equal(t, 'a', p.buffer[0][0])
		assert.Equal(t, 'b', p.buffer[2][0])
		fmt.Printf("%v\n", p.buffer)
	})

	t.Run("Adds a character at the same location twice", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.AddChar(0, 0, 'a')
		assert.Nil(t, err)
		err = p.AddChar(0, 0, 'b')
		assert.Nil(t, err)
		assert.Equal(t, 'b', p.buffer[0][0])
		assert.Equal(t, 'a', p.buffer[0][1])
	})

	t.Run("Adds a character at the end of the page when there's already a character", func(t *testing.T) {
		p := NewPage(10, 10)
		p.buffer[9][9] = 'a'
		err := p.AddChar(9, 9, 'b')
		assert.Nil(t, err)
		assert.Equal(t, 'b', p.buffer[9][9])
	})
}

func TestRemoveChar(t *testing.T) {
	t.Run("Fails if y is less than 0", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.RemoveChar(0, -1)
		assert.EqualError(t, err, "y is less than 0")
	})

	t.Run("Fails if x is less than 0", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.RemoveChar(-1, 0)
		assert.EqualError(t, err, "x is less than 0")
	})

	t.Run("Fails if y is greater than height", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.RemoveChar(0, 10)
		assert.EqualError(t, err, "y is greater than height")
	})

	t.Run("Fails if x is greater than width", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.RemoveChar(10, 0)
		assert.EqualError(t, err, "x is greater than width")
	})

	t.Run("Removes a character at 0, 0 in the empty page", func(t *testing.T) {
		p := NewPage(10, 10)
		err := p.RemoveChar(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, p.buffer[0][0])
	})

	t.Run("Removes a character at 0, 0 in the non-empty page", func(t *testing.T) {
		p := NewPage(10, 10)
		_ = p.AddChar(0, 0, 'a')
		err := p.RemoveChar(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, p.buffer[0][0])
	})

	t.Run("Removes a character at the same location twice", func(t *testing.T) {
		p := NewPage(10, 10)
		_ = p.AddChar(0, 0, 'a')
		_ = p.AddChar(1, 0, 'b')
		err := p.RemoveChar(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, p.buffer[0][1])
		assert.Equal(t, 'b', p.buffer[0][0])
		err = p.RemoveChar(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, p.buffer[0][0])
		assert.Equal(t, 0, p.buffer[0][1])
	})

	t.Run("Removes a character at the end of the page when there's already a character", func(t *testing.T) {
		p := NewPage(10, 10)
		_ = p.AddChar(9, 9, 'a')
		err := p.RemoveChar(9, 9)
		assert.Nil(t, err)
		assert.Equal(t, 0, p.buffer[9][9])
	})
}

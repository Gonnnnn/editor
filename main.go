package main

import (
	"editor/golang/array"
	"fmt"

	"github.com/nsf/termbox-go"
)

var ARROW_KEYS = []termbox.Key{
	termbox.KeyArrowUp,
	termbox.KeyArrowDown,
	termbox.KeyArrowLeft,
	termbox.KeyArrowRight,
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	cursor := NewCursor(0 /* =initX */, 0 /* =initY */)
	boxWidth, boxHeight := termbox.Size()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			if array.Contains(ARROW_KEYS, ev.Key) {
				x, y := handleArrowKey(ev, cursor, boxWidth, boxHeight)
				termbox.SetCursor(x, y)
				termbox.Flush()
				continue
			}
			if ev.Key == termbox.KeyEnter {
				x, y := cursor.Enter(boxHeight)
				termbox.SetCursor(x, y)
				termbox.Flush()
				continue
			}
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				curX, curY := cursor.CurLocation()
				termbox.SetCell(curX, curY, ' ', termbox.ColorDefault, termbox.ColorDefault)
				x, y := cursor.MoveLeft(boxWidth)
				termbox.SetCursor(x, y)
				termbox.Flush()
				continue
			}
		
			curX, curY := cursor.CurLocation()
			termbox.SetCell(curX, curY, ev.Ch, termbox.ColorDefault, termbox.ColorDefault)
			newX, newY := cursor.MoveRight(boxWidth, boxHeight)
			termbox.SetCursor(newX, newY)
			termbox.Flush()

		case termbox.EventError:
			fmt.Println(ev.Err)
			break loop
		default:
			fmt.Printf("ev: %+v\n", ev)
		}
	}
}

func handleArrowKey(ev termbox.Event, cursor *cursor, boxWidth int, boxHeight int) (int, int) {
	switch {
	case ev.Key == termbox.KeyArrowUp:
		return cursor.MoveUp()
	case ev.Key == termbox.KeyArrowDown:
		return cursor.MoveDown(boxHeight)
	case ev.Key == termbox.KeyArrowLeft:
		return cursor.MoveLeft(boxWidth)
	case ev.Key == termbox.KeyArrowRight:
		return cursor.MoveRight(boxWidth, boxHeight)
	default:
		return cursor.CurLocation()
	}
}
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

	var buffer []rune
	maxWidth, maxHeight := termbox.Size()
	cursor := NewCursor(0 /* =initX */, 0 /* =initY */)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			if array.Contains(ARROW_KEYS, ev.Key) {
				continue
			}
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				if len(buffer) == 0 {
					continue
				}

				newX, newY := cursor.MoveLeft(maxWidth)
				termbox.SetCell(newX, newY, ' ', termbox.ColorDefault, termbox.ColorDefault)
				termbox.SetCursor(newX, newY)
				termbox.Flush()

				buffer = buffer[:len(buffer)-1]
				continue
			}

			curX, curY := cursor.CurLocation()
			termbox.SetCell(curX, curY, ev.Ch, termbox.ColorDefault, termbox.ColorDefault)
			newX, newY := cursor.MoveRight(maxWidth, maxHeight)
			termbox.SetCursor(newX, newY)
			termbox.Flush()

			buffer = append(buffer, ev.Ch)
			
		case termbox.EventError:
			fmt.Println(ev.Err)
			break loop
		default:
			fmt.Printf("ev: %+v\n", ev)
		}
	}
}

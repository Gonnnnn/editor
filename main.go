package main

import (
	"editor/golang/array"
	"fmt"
	"os"

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

	maxWidth, maxHeight := termbox.Size()
	page := NewPage(maxWidth, maxHeight)
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
			if ev.Key == termbox.KeyEnter {
				curX, curY := cursor.CurLocation()
				err := page.AddChar(curX, curY, '\n')
				if err != nil {
					fmt.Println(err)
					continue
				}
				newX, newY := cursor.Enter()
				termbox.SetCursor(newX, newY)
				termbox.Flush()
				continue
			}
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				newX, newY := cursor.MoveLeft(maxWidth)
				err := page.RemoveChar(newX, newY)
				if err != nil {
					fmt.Println(err)
					continue
				}
				termbox.SetCell(newX, newY, ' ', termbox.ColorDefault, termbox.ColorDefault)
				termbox.SetCursor(newX, newY)
				termbox.Flush()
				continue
			}

			curX, curY := cursor.CurLocation()
			err := page.AddChar(curX, curY, ev.Ch)
			if err != nil {
				fmt.Println(err)
				continue
			}
			termbox.SetCell(curX, curY, ev.Ch, termbox.ColorDefault, termbox.ColorDefault)
			newX, newY := cursor.MoveRight(maxWidth, maxHeight)
			termbox.SetCursor(newX, newY)
			termbox.Flush()

		case termbox.EventError:
			fmt.Println(ev.Err)
			break loop
		default:
			fmt.Printf("ev: %+v\n", ev)
		}
	}

	buffer := page.Buffer()
	fileName := "./output/result.txt"
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, row := range buffer {
		file.WriteString(string(row))
	}
}

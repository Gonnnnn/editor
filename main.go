package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var buffer []rune
	curX, curY := 0, 0
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			termbox.SetCell(curX, curY, ev.Ch, termbox.ColorDefault, termbox.ColorDefault)
			curX++

			boxWidth, _ := termbox.Size()
			if curX >= boxWidth {
				curX = 0
				curY++
			}
			termbox.SetCursor(curX, curY)
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

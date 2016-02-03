package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func printWelcome() {
	Printf(3, 18, termbox.ColorWhite, termbox.ColorBlack, "Welcome to Snake")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	Printf(3, 18, termbox.ColorWhite, termbox.ColorBlack, "To begin press any arrow key")
	termbox.Flush()
keyLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowUp ||
				ev.Key == termbox.KeyArrowDown ||
				ev.Key == termbox.KeyArrowRight ||
				ev.Key == termbox.KeyArrowLeft {
				break keyLoop
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
	ClearWindow()

	Printf(3, 18, termbox.ColorWhite, termbox.ColorBlack, "Game will begin in ...")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	Printf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "3 ...")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	Printf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "2 ...")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	Printf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "1 ...")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	Printf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "0 ...")
	termbox.Flush()
	time.Sleep(1 * time.Second)
	printGo()
}

func printGo() {
	ClearWindow()
	for i := 20; i <= 28; i++ {
		Printf(20, i, termbox.ColorWhite, termbox.ColorWhite, " ")
		if i > 23 && i < 28 {
			Printf(25, i, termbox.ColorWhite, termbox.ColorWhite, " ")
		}

		if i > 20 && i < 26 {
			Printf(i, 28, termbox.ColorWhite, termbox.ColorWhite, " ")
			Printf(i, 20, termbox.ColorWhite, termbox.ColorWhite, " ")
		}

		Printf(27, i, termbox.ColorWhite, termbox.ColorWhite, " ")
		Printf(32, i, termbox.ColorWhite, termbox.ColorWhite, " ")
	}

	Printf(23, 24, termbox.ColorWhite, termbox.ColorWhite, " ")
	Printf(24, 24, termbox.ColorWhite, termbox.ColorWhite, " ")

	x := 28
	for ; x <= 31; x++ {
		Printf(x, 20, termbox.ColorWhite, termbox.ColorWhite, " ")
		Printf(x, 28, termbox.ColorWhite, termbox.ColorWhite, " ")
	}

	termbox.Flush()
	time.Sleep(1 * time.Second)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	printWelcome()
	time.Sleep(1 * time.Second)
	GameStart()
	time.Sleep(2 * time.Second)
}

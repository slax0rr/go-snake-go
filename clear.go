package main

import (
	"github.com/nsf/termbox-go"
)

func ClearWindow() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

package tbhelper

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Print(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func Printfln(x, y int, fg, bg termbox.Attribute, format string) {
	s := fmt.Sprintln(format)
	Print(x, y, fg, bg, s)
}

func Printf(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Print(x, y, fg, bg, s)
}

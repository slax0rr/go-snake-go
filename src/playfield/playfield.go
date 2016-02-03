package playfield

import (
	"github.com/nsf/termbox-go"
	"tbhelper"
)

func DrawField(maxX, maxY int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	x := 5
	for ; x <= maxX; x++ {
		tbhelper.Printf(x, 5, termbox.ColorWhite, termbox.ColorWhite, " ")
		tbhelper.Printf(x, maxY, termbox.ColorWhite, termbox.ColorWhite, " ")
	}
	y := 5
	for ; y <= maxY; y++ {
		tbhelper.Printf(5, y, termbox.ColorWhite, termbox.ColorWhite, " ")
		tbhelper.Printf(maxX, y, termbox.ColorWhite, termbox.ColorWhite, " ")
	}

	termbox.Flush()
}

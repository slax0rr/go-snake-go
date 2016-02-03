package game

import (
	"clear"
	"github.com/nsf/termbox-go"
	"math/rand"
	"tbhelper"
	"time"
)

type BodyPart struct {
	x, y int
}

type Snake struct {
	parts []BodyPart
}

var boundaries []int = []int{5, 5, 105, 55}
var startingLength int = 3
var snake Snake

// directions:
// 1 = left
// 2 = down
// 3 = up
// 4 = right
var direction int = 4
var foodPosition []int
var points int = 0

// in milliseconds
const speed = 100 * time.Millisecond

func Start() {
	clear.ClearWindow()
	drawField(boundaries[2], boundaries[3])

	position := getStartingPosition()
	bodyParts := []BodyPart{
		BodyPart{x: position["x"][0], y: position["y"][0]},
		BodyPart{x: position["x"][1], y: position["y"][0]},
		BodyPart{x: position["x"][2], y: position["y"][0]},
	}

	snake = Snake{parts: bodyParts}
	snake.draw()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

movement:
	for {
		select {
		case ev := <-eventQueue:
			changeDirection(ev.Key)
		default:
			snake.updatePosition()
			if snake.checkColision() == true {
				break movement
			}
			clear.ClearWindow()
			drawField(boundaries[2], boundaries[3])
			snake.draw()
			spawnFood()
			termbox.Flush()
			time.Sleep(speed)
		}
	}

	tbhelper.Printf(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Game over, Points: %d", points)
	termbox.Flush()
}

func (snake *Snake) hasEaten() bool {
	if len(foodPosition) == 0 {
		return false
	}
	head := snake.parts[0]
	if foodPosition[0] == head.x && foodPosition[1] == head.y {
		foodPosition = []int{}
		points++

		return true
	}

	return false
}

func (snake *Snake) checkPosition(x, y int, noHead bool) bool {
	parts := snake.parts
	if noHead == true {
		parts = snake.parts[1:]
	}
	for _, part := range parts {
		if x == part.x && y == part.y {
			return true
		}
	}

	return false
}

func (snake *Snake) checkColision() bool {
	head := snake.parts[0]
	if head.x == boundaries[0] ||
		head.x == boundaries[2] ||
		head.y == boundaries[1] ||
		head.y == boundaries[3] {
		return true
	}

	return snake.checkPosition(head.x, head.y, true)
}

func (snake *Snake) updatePosition() {
	head := snake.parts[0]
	body := snake.parts[1:]
	posPrev := []int{head.x, head.y}
	switch direction {
	case 1:
		head.x -= 1
	case 2:
		head.y += 1
	case 3:
		head.y -= 1
	case 4:
		head.x += 1
	}
	parts := []BodyPart{head}
	for _, part := range body {
		prev := posPrev
		posPrev = []int{part.x, part.y}
		part.x = prev[0]
		part.y = prev[1]
		parts = append(parts, part)
	}

	snake.parts = parts
	if snake.hasEaten() == true {
		snake.parts = append(snake.parts, BodyPart{x: posPrev[0], y: posPrev[1]})
	}
}

func (snake *Snake) draw() {
	for _, part := range snake.parts {
		tbhelper.Printf(part.x, part.y, termbox.ColorWhite, termbox.ColorWhite, " ")
	}
}

func getStartingPosition() map[string][]int {
	xPosOffset := boundaries[0] / 2
	yPosOffset := boundaries[1] / 2
	yPositions := []int{((boundaries[3] - boundaries[1]) / 2) + yPosOffset}

	x := ((boundaries[2] - boundaries[0]) / 2) + xPosOffset
	xPositions := []int{x + 1, x, x - 1}

	return map[string][]int{"x": xPositions, "y": yPositions}
}

func drawField(maxX, maxY int) {
	clear.ClearWindow()
	x := boundaries[0]
	for ; x <= maxX; x++ {
		tbhelper.Printf(x, boundaries[1], termbox.ColorWhite, termbox.ColorWhite, " ")
		tbhelper.Printf(x, maxY, termbox.ColorWhite, termbox.ColorWhite, " ")
	}
	y := boundaries[1]
	for ; y <= maxY; y++ {
		tbhelper.Printf(boundaries[0], y, termbox.ColorWhite, termbox.ColorWhite, " ")
		tbhelper.Printf(maxX, y, termbox.ColorWhite, termbox.ColorWhite, " ")
	}
}

func changeDirection(key termbox.Key) {
	switch key {
	case termbox.KeyArrowLeft:
		if direction != 4 {
			direction = 1
		}
	case termbox.KeyArrowDown:
		if direction != 3 {
			direction = 2
		}
	case termbox.KeyArrowUp:
		if direction != 2 {
			direction = 3
		}
	case termbox.KeyArrowRight:
		if direction != 1 {
			direction = 4
		}
	}
}

func spawnFood() {
	for len(foodPosition) == 0 {
		x := boundaries[2] - boundaries[0] - 1
		y := boundaries[3] - boundaries[1] - 1

		randSource := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(randSource)

		foodX := rnd.Intn(x) + 6
		foodY := rnd.Intn(y) + 6

		if snake.checkPosition(foodX, foodY, false) == false {
			foodPosition = []int{foodX, foodY}
		}
	}

	tbhelper.Printf(foodPosition[0], foodPosition[1], termbox.ColorRed, termbox.ColorRed, " ")
}

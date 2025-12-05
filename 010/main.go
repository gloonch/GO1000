package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200 // 20 frames per second
		speed     = time.Second / 20
	)

	var (
		cell      rune
		py, px    int
		velocityX = 1
		velocityY = 1
	)
	// create a board with width rows and height columns
	// each cell has to be initialized with a byte
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	for i := 0; i < maxFrames; i++ {

		// move the ball
		px += velocityX
		py += velocityY

		if px >= width-1 || px <= 0 {
			velocityX *= -1
		}
		if py >= height-1 || py <= 0 {
			velocityY *= -1
		}

		// remove the previous ball position
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// set the ball position
		board[px][py] = true

		buf = buf[:0]
		// draw the board
		for y := range board[0] {
			for x := range board {
				cell = ' '
				if board[x][y] {
					cell = cellBall

				}
				fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			fmt.Println()
			buf = append(buf, '\n')
		}

		fmt.Print("\033[2J")

		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}

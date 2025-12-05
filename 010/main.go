package main

import "fmt"

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'
	)

	var cell rune

	// create a board with width rows and height columns
	// each cell has to be initialized with a byte
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	board[0][0] = true

	// draw the board
	for y := range board[0] {
		for x := range board {
			cell = ' '
			if board[x][y] {
				cell = cellBall

			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}

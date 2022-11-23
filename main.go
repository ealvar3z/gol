// Copyright 2014 e. alvarez. All rights reserved.
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/smartystreets/assertions"
)

type cell int

const (
	HEIGHT = 5
	WIDTH  = 5
)

const (
	DEAD cell = iota
	ALIVE
)

type board struct {
	cells [HEIGHT][WIDTH]cell
}

// Render renders the initial state of the board
func render_board(b board) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			switch b.cells[y][x] {
			case DEAD:
				fmt.Printf(".")
			case ALIVE:
				fmt.Printf("#")
			default:
				assertions.ShouldBeNil(0, "Unreachable")
			}
		}
		fmt.Println()
	}
}

// Neighbors implements the rules of the game
func neighbors(b board, y0, x0 int) int {
	var counter int
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy != 0 || dx != 0 {
				y := y0 + dy
				x := x0 + dx

				if y < 0 {
					y += HEIGHT
				} else {
					y %= HEIGHT
				}

				if x < 0 {
					x += WIDTH
				} else {
					x %= WIDTH
				}

				if b.cells[y][x] == ALIVE {
					counter++
				}
			}
		}
	}
	assertions.ShouldBeLessThan(counter, 9)
	return counter
}

// Next computes the board state based on the
// result of the neighbors() and returns it
func next(b board) board {
	var result board
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			n := neighbors(b, y, x)
			switch b.cells[y][x] {
			case DEAD:
				if n == 3 {
					result.cells[y][x] = ALIVE
				} else {
					result.cells[y][x] = DEAD
				}
			case ALIVE:
				if n == 2 || n == 3 {
					result.cells[y][x] = ALIVE
				} else {
					result.cells[y][x] = DEAD
				}
			default:
				assertions.ShouldBeNil(0, "Unreachable")
			}
		}
	}
	return result
}

func main() {
	board := board{
		[HEIGHT][WIDTH]cell{
			{DEAD, ALIVE, DEAD, DEAD, DEAD},
			{DEAD, DEAD, ALIVE, DEAD, DEAD},
			{ALIVE, ALIVE, ALIVE, DEAD, DEAD},
		},
	}
	// game loop
	for {
		render_board(board)
		board = next(board)
		time.Sleep(150000000)
		fmt.Printf("\033[%dA", HEIGHT)
		fmt.Printf("\033[%dD", WIDTH)
	}
}

package main

import (
	"math/rand"
	"time"

	"github.com/buger/goterm"
)

const width int = 192
const height int = 52

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func getNextGeneration(grid [][]uint8) [][]uint8 {

	previousState := make([][]uint8, len(grid))
	for i := range grid {
		previousState[i] = make([]uint8, len(grid[i]))
		copy(previousState[i], grid[i])
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			numberOfNeighbours :=
				previousState[mod((y+1), height)][x%width] +
					previousState[mod((y-1), height)][x%width] +
					previousState[y%height][mod((x+1), width)] +
					previousState[y%height][mod((x-1), width)] +
					previousState[mod((y+1), height)][mod((x+1), width)] +
					previousState[mod((y+1), height)][mod((x-1), width)] +
					previousState[mod((y-1), height)][mod((x+1), width)] +
					previousState[mod((y-1), height)][mod((x-1), width)]

			if grid[y][x] == 1 {
				if numberOfNeighbours < 2 || numberOfNeighbours > 3 {
					grid[y][x] = 0
				}
			} else {
				if numberOfNeighbours == 3 {
					grid[y][x] = 1
				}
			}
		}
	}

	return grid
}

func populateRandomly(grid *[][]uint8) {
	for y := 0; y < height; y++ {
		var line []uint8
		for x := 0; x < width; x++ {
			line = append(line, uint8(rand.Intn(2)))
		}
		*grid = append(*grid, line)
	}
}

func main() {

	// Create a grid
	var grid [][]uint8

	populateRandomly(&grid)

	for {

		goterm.MoveCursor(1, 1)

		grid = getNextGeneration(grid)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == 0 {
					goterm.Print(" ")
				} else {
					goterm.Print("*")
				}
			}
			goterm.Print("\n")
		}

		goterm.Flush()
		time.Sleep(time.Millisecond * 75)
	}
}

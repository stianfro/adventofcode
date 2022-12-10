package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GridOptions struct {
	xMotions []int
	yMotions []int
	xMax     int
	yMax     int
}

func Input(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func Grid(motions []string) GridOptions {
	var xMotions, yMotions []int

	x := 1
	y := 1

	for _, v := range motions {
		m := strings.Split(v, " ")

		dir := m[0]
		mov, _ := strconv.Atoi(m[1])

		switch dir {
		case "R":
			x += mov
			xMotions = append(xMotions, x)
			yMotions = append(yMotions, y)
		case "L":
			x -= mov
			xMotions = append(xMotions, x)
			yMotions = append(yMotions, y)
		case "U":
			y += mov
			xMotions = append(xMotions, x)
			yMotions = append(yMotions, y)
		case "D":
			y -= mov
			xMotions = append(xMotions, x)
			yMotions = append(yMotions, y)
		}
	}

	sortedMotionsX := xMotions
	sortedMotionsY := yMotions

	sort.Ints(sortedMotionsX)
	sort.Ints(sortedMotionsY)

	options := GridOptions{
		xMotions: xMotions,
		yMotions: yMotions,
		xMax:     sortedMotionsX[len(sortedMotionsX)-1],
		yMax:     sortedMotionsY[len(sortedMotionsY)-1],
	}

	return options
}

func populateGrid(options GridOptions) {
	grid := [][]string{}

	curPosX := 0 // 0 better?
	curPosY := 0 // 0 better?

	// Something is not looping like I want it here
	for i := 0; i < len(options.xMotions); i++ {
		curPosX = options.xMotions[i] - 1
		curPosY = options.yMotions[i] - 1

		row := make([]string, options.xMax)

		// this must be inside a loop with the instructions
		for y := 0; y < options.yMax; y++ {
			if curPosY == y {
				row[curPosX] = "H"
				break
			} else {
				row[curPosX] = "."
				break
			}
		}
		grid = append(grid, row)
	}
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(grid[i])
	}
}

func main() {
	puzzleInput := Input("input-example.txt")
	gridOptions := Grid(puzzleInput)

	populateGrid(gridOptions)
}

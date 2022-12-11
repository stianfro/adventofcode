package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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

func printGrid(grid [][]string) {
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(grid[i])
	}
}

func populateGrid(options GridOptions) {
	grid := make([][]string, options.yMax)

	row := make([]string, options.xMax)
	for i := 0; i < len(row); i++ {
		row[i] = "."
	}

	for i := 0; i < len(grid); i++ {
		grid[i] = row
	}

	printGrid(grid)

	curPosX := 0 // 0 better?
	curPosY := 0 // 0 better?

	// Something is not looping like I want it here
	for i := 0; i < len(options.xMotions); i++ {
		curPosX = options.xMotions[i] - 1
		curPosY = options.yMotions[i] - 1

		grid[curPosY][curPosX] = "H"

		printGrid(grid) // Something not quite right here
	}
}

func main() {
	// puzzleInput := Input("input-example.txt")
	// gridOptions := Grid(puzzleInput)

	// populateGrid(gridOptions)

	p := plot.New()
	p.Title.Text = "Rope Movement"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, 3)
	pts[0].X = 0
	pts[0].Y = 0

	pts[1].X = 4
	pts[1].Y = 0

	pts[2].X = 4
	pts[2].Y = 4

	err := plotutil.AddLinePoints(p,
		"First", pts,
	)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

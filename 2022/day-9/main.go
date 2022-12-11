package main

import (
	"fmt"
	"os"
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

	xMotions = append(xMotions, 1)
	yMotions = append(yMotions, 1)

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

	_, xMax := MinMax(xMotions)
	_, yMax := MinMax(yMotions)

	options := GridOptions{
		xMotions: xMotions,
		yMotions: yMotions,
		xMax:     xMax,
		yMax:     yMax,
	}

	return options
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func printGrid(grid [][]string) {
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(grid[i])
	}
}

func populateGrid(options GridOptions) {
	p := plot.New()
	p.Title.Text = "Rope Movement"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	pts := make(plotter.XYs, len(options.xMotions))

	// Something is not looping like I want it here
	for i := 0; i < len(options.xMotions); i++ {
		positionX := float64(options.xMotions[i] - 1)
		positionY := float64(options.yMotions[i] - 1)

		pts[i].X = positionX
		pts[i].Y = positionY
	}

	err := plotutil.AddLinePoints(p,
		"Head", pts,
	)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func main() {
	puzzleInput := Input("input-example.txt")
	gridOptions := Grid(puzzleInput)

	populateGrid(gridOptions)
}

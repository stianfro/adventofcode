package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// var xMax []int
// var yMax []int

// type Coordinate struct {
// 		X int
// 		Y int
// }

// var headPosition Coordinate
// var tailPosition Coordinate

var xMotions []int
var yMotions []int

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func generateGrid(motions []string) {
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
		case "L":
			x -= mov
			xMotions = append(xMotions, x)
		case "U":
			y += mov
			yMotions = append(yMotions, y)
		case "D":
			y -= mov
			yMotions = append(yMotions, y)
		}
	}

	sort.Ints(xMotions)
	sort.Ints(yMotions)

	xMax := xMotions[len(xMotions)-1]
	yMax := yMotions[len(yMotions)-1]

	fmt.Println(xMax, yMax)
}

func main() {
	input := getInput("input-example.txt")

	generateGrid(input)

	grid := [][]string{
		{".", ".", ".", ".", ".", "."}, // grid[0]
		{".", ".", ".", ".", ".", "."}, // grid[1]
		{".", ".", ".", ".", ".", "."}, // grid[2]
		{".", ".", ".", ".", ".", "."}, // grid[3]
		{"H", ".", ".", ".", ".", "."}, // grid[4]
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}

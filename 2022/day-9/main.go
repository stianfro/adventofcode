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

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func generateGrid(motions []string) (int, int) {
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

	return xMotions[len(xMotions)-1], yMotions[len(yMotions)-1]
}

func populateGrid(motions []string, xMax, yMax int) {
	grid := [][]string{}

	curPosX := 0 // 0 better?
	curPosY := 0 // 0 better?

	for _, v := range motions {
		m := strings.Split(v, " ")

		dir := m[0]
		mov, _ := strconv.Atoi(m[1])

		switch dir {
		case "R":
			curPosX += mov
		case "L":
			curPosX -= mov
		case "U":
			curPosY += mov
		case "D":
			curPosY -= mov
		}

		var row = make([]string, xMax)

		// this must be inside a loop with the instructions
		for y := 0; y < yMax; y++ { // might need to start on 0
			if curPosY == y {
				row[curPosX] = "H"
				break
			} else {
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
	input := getInput("input-example.txt")

	xMax, yMax := generateGrid(input)
	fmt.Println(xMax, yMax)

	populateGrid(input, xMax, yMax)

	gridTest := [][]string{
		{".", ".", ".", ".", ".", "."}, // grid[0]
		{".", ".", ".", ".", ".", "."}, // grid[1]
		{".", ".", ".", ".", ".", "."}, // grid[2]
		{".", ".", ".", ".", ".", "."}, // grid[3]
		{"H", ".", ".", ".", ".", "."}, // grid[4]
	}

	for _, rowTest := range gridTest {
		fmt.Println(rowTest)
	}
}

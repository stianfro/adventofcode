package main

import (
	"os"
	"strings"
)

var headPosition Coordinate
var tailPosition Coordinate

type Coordinate struct {
	X int
	Y int
}

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func generateGrid(motions []string) {

}

func main() {
	input := getInput("input-example.txt")

	generateGrid(input)
}

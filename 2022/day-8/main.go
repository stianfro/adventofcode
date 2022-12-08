package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func main() {
	input := getInput("input-example.txt")

	// 30373 rowA
	// 25512 rowB
	// 65332 rowC

	for i := range input {
		if i >= len(input)-2 {
			break
		}

		// Could (and should) make a map here and then loop over it below (with k, v)

		rowA := strings.Split(input[i], "")
		rowB := strings.Split(input[i+1], "")
		rowC := strings.Split(input[i+2], "")

		fmt.Println(rowA)
		fmt.Println(rowB)
		fmt.Println(rowC)

		visibleTrees := 0

		for i := range rowA {
			// Don't check left tree if i is the first element in row
			if i != 0 {
				// Check left tree
				if rowA[i-1] < rowA[i] {
					visibleTrees += 1
					// isVisible = true
					break
				}
			} else if i != len(rowA) { // Don't check right tree if i is the last element in row
				// Check right tree
				if rowA[i+1] < rowA[i] {
					visibleTrees += 1
					break
				}
			} else {
				// Check tree below
				if rowA[i] > rowB[i] {
					visibleTrees += 1
					break
				}
			}
		}

		for i := range rowB {
			if i != 0 {
				if rowB[i-1] < rowB[i] {
					visibleTrees += 1
					break
				}
			} else if i != len(rowB) {
				if rowB[i+1] < rowB[i] {
					visibleTrees += 1
					break
				}
			} else {
				// Check tree above

				// Check tree below
				if rowA[i] > rowB[i] {
					visibleTrees += 1
					break
				}
			}
		}
	}
}

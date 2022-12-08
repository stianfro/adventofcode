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

	visibleTrees := 0

	for i := range input {
		if i >= len(input)-2 {
			break
		}

		// Could (and should) make a map here and then loop over it below (with k, v)

		row := make(map[string][]string)

		row["A"] = strings.Split(input[i], "")
		row["B"] = strings.Split(input[i+1], "")
		row["C"] = strings.Split(input[i+2], "")

		for k := range row {
			fmt.Println(k, row[k])
		}

		for i := range row["A"] {
			// Don't check left tree if i is the first element in row
			if i != 0 {
				// Check left tree
				if row["A"][i-1] < row["A"][i] {
					visibleTrees++
					break
				}
			} else if i != len(row["A"]) { // Don't check right tree if i is the last element in row
				// Check right tree
				if row["A"][i+1] < row["A"][i] {
					visibleTrees++
					break
				}
			} else {
				// Check tree below
				if row["A"][i] > row["B"][i] {
					visibleTrees++
					break
				}
			}
		}

		for i := range row["B"] {
			if i != 0 {
				if row["B"][i-1] < row["B"][i] {
					visibleTrees++
					break
				}
			} else if i != len(row["B"]) {
				if row["B"][i+1] < row["B"][i] {
					visibleTrees++
					break
				}
			} else {
				// Check tree above
				if row["B"][i] > row["A"][i] {
					visibleTrees++
					break
				}
				// Check tree below
				if row["A"][i] > row["B"][i] {
					visibleTrees++
					break
				}
			}
		}

		for i := range row["C"] {
			if i != 0 {
				if row["C"][i-1] < row["C"][i] {
					visibleTrees++
					break
				}
			} else if i != len(row["C"]) {
				if row["C"][i+1] < row["C"][i] {
					visibleTrees++
					break
				}
			} else {
				if row["C"][i] > row["A"][i] {
					visibleTrees++
					break
				}
			}
		}
	}

	fmt.Println(visibleTrees)
}

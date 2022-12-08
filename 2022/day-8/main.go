package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

var visibleT, visibleB, visibleL, visibleR bool

// for every tree (t) in every row, check if all trees until edge (e) is lower than t
// if all tress are lower until edge, add 1 to visibleTrees
func main() {
	input := getInput("input-example.txt")

	visibleTrees := 0
	allRows := [][]string{}

	for _, row := range input {
		r := strings.Split(row, "")
		allRows = append(allRows, r)
	}

	edgeT := 0
	edgeB := len(allRows) - 1

	for i, row := range allRows {
		fmt.Println(allRows[i])

		if i == edgeT {
			visibleTrees++
			continue
		} else if i == edgeB {
			visibleTrees++
			continue
		}

		edgeL := 0
		edgeR := len(row) - 1

		for i, tree := range row {
			t, _ := strconv.Atoi(tree)

			if i == edgeL {
				visibleTrees++
				continue
			} else if i == edgeR {
				visibleTrees++
				continue
			}

			distanceToEdgeR := len(row) - i
			distanceToEdgeL := i

			// l->r
			for a := 0; a < distanceToEdgeR; a++ {
				T, _ := strconv.Atoi(row[a])
				if t > T {
					fmt.Println(row[a])
					break
				} else {
					continue
				}
			}
			// l<-r
			for b := distanceToEdgeL - 1; b >= 0; b-- {
				T, _ := strconv.Atoi(row[b])
				if t > T {
					fmt.Println(row[b])
					break
				} else {
					continue
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println(visibleTrees)
}

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

func makeRange(min, max string) []int {
	minInt, _ := strconv.Atoi(min)
	maxInt, _ := strconv.Atoi(max)

	a := make([]int, maxInt-minInt+1)
	for i := range a {
		a[i] = minInt + i
	}
	return a
}

func main() {
	partOne()
}

func partOne() {
	answer := 0
	pairRgAll := [][][]int{}

	for _, pair := range getInput("input-example.txt") {
		pairRg := [][]int{}
		for _, sect := range strings.Split(pair, ",") {
			sectId := strings.Split(sect, "-")
			sectRg := makeRange(sectId[0], sectId[1])
			pairRg = append(pairRg, sectRg)
		}
		pairRgAll = append(pairRgAll, pairRg)
	}

	for _, pair := range pairRgAll {
		assignmentsA := pair[0]
		assignmentsB := pair[1]

		for _, sectionA := range assignmentsA {
			//   sectionA = 3
			//   assignmentsA = [2,3,4,5,6,7,8] (7)
			for idx, sectionB := range assignmentsB {
				// sectionB = 3
				// assignmentsB =   [3,4,5,6,7] (5)
				if sectionA == sectionB {
					if idx+1 == len(assignmentsB) {
						answer += 1
					}
				}
			}
		}
	}

	fmt.Println(answer)
}

// Guesses
// 491 (too low)
// 982 (too high)
// 494 (too low)
// 497

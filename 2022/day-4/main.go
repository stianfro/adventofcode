package main

import (
	"fmt"
	"os"
	"reflect"
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
	pairRgAll := [][][]int{}

	for _, pair := range getInput("input.txt") {
		pairRg := [][]int{}
		for _, sect := range strings.Split(pair, ",") {
			sectId := strings.Split(sect, "-")
			sectRg := makeRange(sectId[0], sectId[1])
			pairRg = append(pairRg, sectRg)
		}
		pairRgAll = append(pairRgAll, pairRg)
	}

	pairDupTot := [][][]int{}

	for _, pair := range pairRgAll {
		assignmentsA := pair[0]
		assignmentsB := pair[1]

		for _, sectionA := range assignmentsA {
			for idx, sectionB := range assignmentsB {
				if sectionA == sectionB {
					if idx+1 == len(assignmentsB) {
						pairDup := [][]int{assignmentsA, assignmentsB}
						pairDupTot = append(pairDupTot, pairDup)
					}
				}
			}
		}

		for _, sectionB := range assignmentsB {
			for idx, sectionA := range assignmentsA {
				if sectionB == sectionA {
					if idx+1 == len(assignmentsA) && len(assignmentsA) < len(assignmentsB) {
						pairDup := [][]int{assignmentsA, assignmentsB}
						if len(pairDupTot) >= 1 {
							if !reflect.DeepEqual(pairDupTot[len(pairDupTot)-1], pairDup) {
								pairDupTot = append(pairDupTot, pairDup)
							}
						}
					}
				}
			}
		}
	}

	answer := 0

	for _, dup := range pairDupTot {
		fmt.Println(dup)
		answer += 1
		if reflect.DeepEqual(dup[0], dup[1]) {
			answer += 1
		}
	}

	fmt.Println(answer)
}

// Guesses
// 491 (too low)
// 982 (too high)
// 494 (too low)
// 497
// 757
// 760

// 1.. 1-1
// 123 1-3
// PROBLEM HERE

// answer = 3

// 1-3,1-1
// 1-1,1-3
// 1-3,1-3

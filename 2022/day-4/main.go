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

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	partOne()
}

func partOne() {
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

	for _, a := range pairRgAll {
		overlap := Intersection(a[0], a[1])

		if len(overlap) < len(a[0]) || len(overlap) < len(a[1]) {
			fmt.Println(overlap)
		}
	}
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

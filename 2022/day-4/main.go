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
	p1 := 0

	for _, line := range getInput("input.txt") {
		a := strings.Split(line, ",")
		one, two := a[0], a[1]

		b := strings.Split(one, "-")
		c := strings.Split(two, "-")

		s1, e1 := b[0], b[1]
		s2, e2 := c[0], c[1]

		s1i, _ := strconv.Atoi(s1)
		e1i, _ := strconv.Atoi(e1)
		s2i, _ := strconv.Atoi(s2)
		e2i, _ := strconv.Atoi(e2)

		if s1i <= s2i && e2i <= e1i || s2i <= s1i && e1i <= e2i {
			p1 += 1
		}
	}

	fmt.Println(p1)
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

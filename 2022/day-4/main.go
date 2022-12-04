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
	p1 := 0
	p2 := 0

	for _, line := range getInput("input.txt") {
		a := strings.Split(line, ",")
		one, two := a[0], a[1]

		b := strings.Split(one, "-")
		c := strings.Split(two, "-")

		s1s, e1s := b[0], b[1]
		s2s, e2s := c[0], c[1]

		s1, _ := strconv.Atoi(s1s)
		e1, _ := strconv.Atoi(e1s)
		s2, _ := strconv.Atoi(s2s)
		e2, _ := strconv.Atoi(e2s)

		if s1 <= s2 && e2 <= e1 || s2 <= s1 && e1 <= e2 {
			p1 += 1
		}
		if !(e1 < s2 || s1 > e2) {
			p2 += 1
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

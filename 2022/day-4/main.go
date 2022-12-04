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
	answer := 0
	pairRg := [][]int{}

	for _, pair := range getInput("input.txt") {
		for _, sect := range strings.Split(pair, ",") {
			sectId := strings.Split(sect, "-")
			sectRg := makeRange(sectId[0], sectId[1])
			pairRg = append(pairRg, sectRg)
		}
	}

	for a, x := range pairRg {
		if a%2 != 0 {
			continue
		}
		for _, y := range x {
			for b, z := range pairRg[a+1] {
				if y == z {
					if b+1 == len(pairRg[a+1]) {
						answer += 1
						if reflect.DeepEqual(pairRg[a], pairRg[a+1]) {
							answer += 1
						}
					}
				} else {
					continue
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

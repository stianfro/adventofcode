// Shameless copy of https://github.com/jonathanpaulson/AdventOfCode/blob/master/2022/9.py
// After several attempts of trying to solve it on my own
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Input(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}
	return strings.Split(string(inputBytes), "\n")
}

func adjust(H, T []int) []int {
	tail := []int{0, 0}

	dr := H[0] - T[0]
	dc := H[1] - T[1]

	if math.Abs(float64(dr)) <= 1 && math.Abs(float64(dc)) <= 1 {
		// pass
	} else if math.Abs(float64(dr)) >= 2 && math.Abs(float64(dc)) >= 2 {
		if T[0]-1 < H[0] {
			tail[0] = H[0] - 1
		} else {
			tail[0] = H[0] + 1
		}

		if T[1] < H[1] {
			tail[1] = H[1] - 1
		} else {
			tail[1] = H[1] + 1
		}
	} else if math.Abs(float64(dr)) >= 2 {
		if T[0] < H[0] {
			tail[0] = H[0] - 1
		} else {
			tail[0] = H[0] + 1
		}

		tail[1] = H[1]
	} else if math.Abs(float64(dc)) >= 2 {
		tail[0] = H[0]

		if T[1] < H[1] {
			tail[1] = H[1] - 1
		} else {
			tail[1] = H[1] + 1
		}
	}

	return tail
}

func main() {
	puzzleInput := Input("input-example.txt")

	H := []int{0, 0}
	T := [][]int{{0, 0}}

	DR := map[string]int{
		"L": 0,
		"U": -1,
		"R": 0,
		"D": 1,
	}
	DC := map[string]int{
		"L": -1,
		"U": 0,
		"R": 1,
		"D": 0,
	}
	P1 := [][]int{{0, 0}}
	// P2 := T[8]
	for _, line := range puzzleInput {
		motion := strings.Split(line, " ")

		direction := motion[0]
		amount, _ := strconv.Atoi(motion[1])

		for i := 0; i < amount; i++ {
			H = []int{H[0] + DR[direction], H[1] + DC[direction]}
			T[0] = adjust(H, T[0])
			for i := 1; i == 9; i++ {
				T[i] = adjust(T[i-1], T[i])
			}
			P1 = append(P1, T[0])
			fmt.Println(T[0])
		}
	}

	fmt.Println(len(P1))
}

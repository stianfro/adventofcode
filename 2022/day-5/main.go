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

func printStacks(stacks [][]string) {
	for i, v := range stacks {
		fmt.Println(i+1, v)
	}
}

func printStacksTop(stacks [][]string) string {
	top := ""
	for _, v := range stacks {
		if len(v) != 0 {
			top += v[len(v)-1]
		}
	}
	return top
}

func stackSum(stacks [][]string) int {
	sum := 0
	for _, v := range stacks {
		sum += len(v)
	}
	return sum
}

func main() {
	var I []string
	var stacks [][]string

	example := false

	if example {
		I = getInput("input-example.txt")
		stacks = [][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		}
	} else {
		I = getInput("input.txt")
		stacks = [][]string{
			{"S", "L", "W"},
			{"J", "T", "N", "Q"},
			{"S", "C", "H", "F", "J"},
			{"T", "R", "M", "W", "N", "G", "B"},
			{"T", "R", "L", "S", "D", "H", "Q", "B"},
			{"M", "J", "B", "V", "F", "H", "R", "L"},
			{"D", "W", "R", "N", "J", "M"},
			{"B", "Z", "T", "F", "H", "N", "D", "J"},
			{"H", "L", "Q", "N", "B", "F", "T"},
		}
	}

	printStacks(stacks)
	fmt.Println("")

	fmt.Println("Start moving crates:")
	for _, i := range I {
		instruction := strings.Split(i, " ")

		mov, _ := strconv.Atoi(instruction[1])
		src, _ := strconv.Atoi(instruction[3])
		dst, _ := strconv.Atoi(instruction[5])

		src = src - 1
		dst = dst - 1

		fmt.Println(i)

		toMove := mov // e.g 4

		for i := 0; i < mov; i++ {
			if !(toMove <= 0) {
				x := toMove
				y := toMove
				for i := 0; i < y; i++ {
					stacks[dst] = append(stacks[dst], stacks[src][len(stacks[src])-x])
					x -= 1
				}
				stacks[src] = stacks[src][:len(stacks[src])-y]
				toMove -= y
			}
		}
		printStacks(stacks)
		fmt.Println("")
	}

	answer := printStacksTop(stacks)
	fmt.Println(answer)
}

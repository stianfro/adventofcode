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

func main() {
	var I []string
	var stacks [][]string

	example := true

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
			if toMove >= 1 && toMove <= 3 {
				for i := 0; i < toMove; i++ {
					stacks[dst] = append(stacks[dst], stacks[src][len(stacks[src])-(toMove-i)])
				}
				stacks[src] = stacks[src][:len(stacks[src])-toMove]
				toMove -= 3
			} else {
				if len(stacks[src]) != 0 {
					stacks[dst] = append(stacks[dst], stacks[src][len(stacks[src])-1])
					stacks[src] = stacks[src][:len(stacks[src])-1]
				}
			}
		}
		printStacks(stacks)
		fmt.Println("")
	}

	answer := printStacksTop(stacks)
	fmt.Println(answer)
}

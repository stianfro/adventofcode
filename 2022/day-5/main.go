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
		top += v[len(v)-1]
	}
	return top
}

func main() {
	// stacks := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	stacks := [][]string{
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

	fmt.Println("--- Part One ---")
	printStacks(stacks)
	fmt.Println("")

	fmt.Println("Start moving crates:")
	for _, i := range getInput("input.txt") {
		instruction := strings.Split(i, " ")

		mov, _ := strconv.Atoi(instruction[1])
		src, _ := strconv.Atoi(instruction[3])
		dst, _ := strconv.Atoi(instruction[5])

		fmt.Println(i)

		for i := 0; i < mov; i++ {
			stacks[dst-1] = append(stacks[dst-1], stacks[src-1][len(stacks[src-1])-1])
			stacks[src-1] = stacks[src-1][:len(stacks[src-1])-1]
		}
		printStacks(stacks)
		fmt.Println("")
	}

	p1 := printStacksTop(stacks)
	fmt.Println(p1)
}

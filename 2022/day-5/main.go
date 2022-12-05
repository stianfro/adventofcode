package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printStacks(stacks [][]string) {
	for _, v := range stacks {
		fmt.Println(v)
	}
}

func printStacksTop(stacks [][]string) {
	for _, v := range stacks {
		fmt.Println(v[len(v)-1])
	}
}

func main() {
	stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	// stacks := [][]string{
	// 	{"S", "L", "W"},
	// 	{"J", "T", "N", "Q"},
	// 	{"S", "C", "H", "F", "J"},
	// 	{"T", "R", "M", "W", "N", "G", "B"},
	// 	{"T", "R", "L", "S", "D", "H", "Q", "B"},
	// 	{"M", "J", "B", "V", "F", "H", "R", "L"},
	// 	{"D", "W", "R", "N", "J", "M"},
	// 	{"B", "Z", "T", "F", "H", "N", "D", "J"},
	// 	{"H", "L", "Q", "N", "B", "F", "T"},
	// }

	I := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	for _, i := range I {
		instruction := strings.Split(i, " ")

		mov, _ := strconv.Atoi(instruction[1])
		src, _ := strconv.Atoi(instruction[3])
		dst, _ := strconv.Atoi(instruction[5])

		fmt.Println(mov, src, dst)
	}

	printStacks(stacks)
	printStacksTop(stacks)
}

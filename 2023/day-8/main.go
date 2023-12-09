package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n\n")

	x, y := data[0], data[1]
	steps := strings.Split(x, "")
	nodes := strings.Split(y, "\n")

	network := make(map[string][]string)

	for _, line := range nodes {
		if line == "" {
			continue
		}
		s := strings.Split(line, " = ")
		pos, targets := s[0], s[1]
		targets = strings.Trim(targets, "()")
		network[pos] = strings.Split(targets, ", ")
	}

	stepCount := 0
	current := "AAA"

	for current != "ZZZ" {
		for _, step := range steps {
			stepCount++

			var direction int

			if step == "L" {
				direction = 0
			} else {
				direction = 1
			}

			current = network[current][direction]
		}
	}

	fmt.Println(stepCount)
}

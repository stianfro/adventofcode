package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input-test.txt")
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
		network[pos] = strings.Split(targets, ", ")
	}

	fmt.Println(steps)
	fmt.Println(network)
}

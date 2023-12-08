package main

import (
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Left  string
	Right string
}

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n\n")

	elements := make(map[string]Element)

	x, y := data[0], data[1]
	instructions := strings.Split(x, "")
	nodes := strings.Split(y, "\n")

	for _, line := range nodes {
		var element Element

		if line == "" {
			continue
		}

		s := strings.Split(line, " = ")
		lr := strings.Trim(s[1], "()")
		lrSlice := strings.Split(lr, ", ")

		element.Left = lrSlice[0]
		element.Right = lrSlice[1]

		elements[s[0]] = element
	}

	next := "AAA"
	goal := false
	step := 0

	for goal == false {
		for _, direction := range instructions {
			step++

			element := elements[next]

			switch direction {
			case "L":
				next = element.Left
			case "R":
				next = element.Right
			}

			if next == "ZZZ" {
				fmt.Println("Goal! Steps:", step)
				goal = true
			}
		}
	}
}

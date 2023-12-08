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

	var startingNodes []string

	for _, line := range nodes {
		var element Element

		if line == "" {
			continue
		}

		s := strings.Split(line, " = ")
		id := s[0]
		lr := strings.Trim(s[1], "()")
		lrSlice := strings.Split(lr, ", ")

		element.Left = lrSlice[0]
		element.Right = lrSlice[1]

		elements[id] = element

		end := id[len(id)-1]

		if string(end) == "A" {
			startingNodes = append(startingNodes, id)
		}
	}

	goal := false
	step := 0

	goals := make(map[string]bool)

	for _, n := range startingNodes {
		goals[n] = false
	}

	var nexts []string

	for _, n := range startingNodes {
		nexts = append(nexts, n)
	}

	for goal == false {

		for _, direction := range instructions {
			step++

			for a := range nexts {
				element := elements[nexts[a]]

				switch direction {
				case "L":
					nexts[a] = element.Left
					fmt.Println("Step:", step, "Direction: left", "Leading to:", nexts[a])
				case "R":
					nexts[a] = element.Right
					fmt.Println("Step:", step, "Direction: right", "Leading to:", nexts[a])
				}

				for b := range nexts {
					end := string(nexts[b][len(nexts[b])-1])

					if end == "Z" {
						goals[startingNodes[a]] = true
					} else {
						goals[startingNodes[a]] = false
					}
				}
			}
			fmt.Println("")

			for _, v := range goals {
				if v != true {
					continue
				}
				goal = true
			}
		}
	}

	fmt.Println(step)
}

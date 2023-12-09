package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n\n")

	steps, y := data[0], data[1]
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

	var positions []string
	for key := range network {
		if strings.HasSuffix(key, "A") {
			positions = append(positions, key)
		}
	}

	cycles := [][]int{}

	for _, current := range positions {
		cycle := []int{}

		currentSteps := steps
		stepCount := 0
		firstZ := ""

		for true {
			for stepCount == 0 || !strings.HasSuffix(current, "Z") {
				stepCount++

				var direction int
				if string(currentSteps[0]) == "L" {
					direction = 0
				} else {
					direction = 1
				}

				current = network[current][direction]
				currentSteps = currentSteps[1:] + string(currentSteps[0])
			}

			cycle = append(cycle, stepCount)

			if firstZ == "" {
				firstZ = current
				stepCount = 0
			} else if current == firstZ {
				break
			}
		}

		cycles = append(cycles, cycle)
	}

	var nums []int

	for _, cycle := range cycles {
		nums = append(nums, cycle[0])
	}

	sort.Ints(nums)

	lcm := nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	for _, num := range nums {
		lcm = lcm * num / gcd(lcm, num)
	}

	fmt.Println(lcm)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

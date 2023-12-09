package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input-example.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var sequenceArray [][]int

		// loop until sumDiffs == 0 ?
		var history []int
		var diffs []int
		var diffsBelow []int

		if line == "" {
			continue
		}

		s := strings.Split(line, " ")
		for _, v := range s {
			vI, _ := strconv.Atoi(v)
			history = append(history, vI)
		}
		sequenceArray = append(sequenceArray, history)

		for i := range history {
			if i == 0 {
				continue
			}

			curr := history[i]
			prev := history[i-1]

			diff := curr - prev
			diffs = append(diffs, diff)
		}
		sequenceArray = append(sequenceArray, diffs)

		for i := range diffs {
			if i == 0 {
				continue
			}

			curr := diffs[i]
			prev := diffs[i-1]

			diff := curr - prev
			diffsBelow = append(diffsBelow, diff)
		}
		sequenceArray = append(sequenceArray, diffsBelow)

		fmt.Println(history)
		fmt.Println(diffs)
		fmt.Println(diffsBelow)
		fmt.Println("")

		lastHistory := history[len(history)-1]
		lastDiff := diffs[len(diffs)-1]
		lastDiffBelow := diffsBelow[len(diffsBelow)-1]

		a := lastDiff + lastDiffBelow
		b := lastHistory + a

		history = append(history, b)
		diffs = append(diffs, a)
		diffsBelow = append(diffsBelow, lastDiffBelow)

		fmt.Println(history)
		fmt.Println(diffs)
		fmt.Println(diffsBelow)
		fmt.Println("")

		var sumDiffsBelow int
		for _, v := range diffsBelow {
			sumDiffsBelow += v
		}

		if sumDiffsBelow == 0 {
			// next
			fmt.Println(sequenceArray)
		} else {
			// add new sequence to sequencearray
			fmt.Println("too bad")
		}

		// os.Exit(0)
	}
}

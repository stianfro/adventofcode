package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	Lists []List
}

type List struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func main() {
	var seeds []int
	var allMaps []Map

	input, _ := os.ReadFile("input-example.txt")
	data := strings.Split(string(input), "\n\n")

	for i, line := range data {
		if i == 0 {
			numbers := strings.Split(line, ": ")
			numberSlice := strings.Split(numbers[1], " ")
			for i := range numberSlice {
				seed, _ := strconv.Atoi(numberSlice[i])
				seeds = append(seeds, seed)
			}
			continue
		}

		var m Map

		numbers := strings.Split(line, ":\n")
		numberSlice := strings.Split(numbers[1], "\n")

		for _, entry := range numberSlice {
			var list List

			if entry == "" {
				continue
			}

			ranges := strings.Split(entry, " ")

			list.DestinationRangeStart, _ = strconv.Atoi(ranges[0])
			list.SourceRangeStart, _ = strconv.Atoi(ranges[1])
			list.RangeLength, _ = strconv.Atoi(ranges[2])

			m.Lists = append(m.Lists, list)
		}

		allMaps = append(allMaps, m)
	}

	for index, m := range allMaps {
		fmt.Println("---")
		fmt.Println(index, m)

		for _, list := range m.Lists {
			var sourceRange []int
			var destRange []int

			for i := 0; i < list.RangeLength; i++ {
				sourceRange = append(sourceRange, list.SourceRangeStart+i)
				destRange = append(destRange, list.DestinationRangeStart+i)
			}

			fmt.Println("src:", sourceRange)
			fmt.Println("dst:", destRange)
			fmt.Println("")
		}

	}

	println("finished")
}

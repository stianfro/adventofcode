package main

import (
	"fmt"
	"os"
	"sort"
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

type Number struct {
	Current []int
	Next    []int
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

	x := Number{
		Current: seeds,
	}

	fmt.Println(x.Current)

	for _, m := range allMaps {
		x.Next = findNext(m, x.Current)
		fmt.Println(x.Current, x.Next)
		x.Current = x.Next
	}

	sort.Ints(x.Current)
	fmt.Println(x.Current[0])
}

func findNext(m Map, currentNumbers []int) []int {
	var sourceRange []int
	var destRange []int

	for a := 0; a < len(m.Lists); a++ {
		fmt.Println(a, "/", len(m.Lists))

		for b := 0; b < m.Lists[a].RangeLength; b++ {
			sourceRange = append(sourceRange, m.Lists[a].SourceRangeStart+b)
			destRange = append(destRange, m.Lists[a].DestinationRangeStart+b)
		}
	}

	sourceDestMap := make(map[int]int)

	for i := range sourceRange {
		sourceDestMap[sourceRange[i]] = destRange[i]
	}

	var nextNumbers []int

	for _, s := range currentNumbers {
		correspondingNumber := sourceDestMap[s]

		if correspondingNumber == 0 {
			correspondingNumber = s
		}

		nextNumbers = append(nextNumbers, correspondingNumber)
	}

	return nextNumbers
}

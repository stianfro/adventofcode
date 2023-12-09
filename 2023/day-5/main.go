package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n\n")

	seeds := strings.Split(strings.Split(data[0], ": ")[1], " ")

	blocks := data[1:]
	for _, block := range blocks {
		ranges := [][]string{}
		for _, line := range strings.Split(block, "\n")[1:] {
			if line == "" {
				continue
			}
			s := strings.Split(line, " ")
			a, b, c := s[0], s[1], s[2]
			ranges = append(ranges, []string{a, b, c})
		}
		new := []int{}
		for _, x := range seeds {
			xI, _ := strconv.Atoi(x)
			var found bool
			for _, r := range ranges {
				a, b, c := r[0], r[1], r[2]
				aI, _ := strconv.Atoi(a)
				bI, _ := strconv.Atoi(b)
				cI, _ := strconv.Atoi(c)

				if inRange(xI, bI, cI) {
					new = append(new, xI-bI+aI)
					found = true
					break
				}
			}
			if !found {
				new = append(new, xI)
			}
		}
		newSeed := []string{}
		for _, n := range new {
			newSeed = append(newSeed, strconv.Itoa(n))
		}
		seeds = newSeed
	}

	seedsInt := []int{}
	for _, s := range seeds {
		i, _ := strconv.Atoi(s)
		seedsInt = append(seedsInt, i)
	}
	sort.Ints(seedsInt)
	fmt.Println(seedsInt[0])
}

func inRange(x, b, c int) bool {
	if b <= x && x < b+c {
		return true
	} else {
		return false
	}
}

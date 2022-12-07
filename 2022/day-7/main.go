package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}

	return strings.Split(string(inputBytes), "\n")
}

func main() {
	path := []string{}
	SZ := make(map[string]int)
	for _, v := range getInput("input.txt") {
		words := strings.Split(v, " ")
		if words[1] == "cd" {
			if words[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, words[2])
			}
		} else if words[1] == "ls" {
			continue
		} else if words[0] == "dir" {
			continue
		} else {
			size, _ := strconv.Atoi(words[0])
			for i := 1; i < len(path)+1; i++ {
				p := strings.Join(path[:i], "/")
				SZ[p] += size
			}
		}
	}

	p1 := 0
	p2 := 0

	spaceTotal := 70000000
	spaceNeeded := 30000000
	spaceUsed := SZ["/"]
	spaceAvailable := spaceTotal - spaceUsed
	spaceMinimum := spaceNeeded - spaceAvailable

	fmt.Println("Total:    ", spaceTotal)
	fmt.Println("Needed:   ", spaceNeeded)
	fmt.Println("Used:     ", spaceUsed)
	fmt.Println("Available:", spaceAvailable)
	fmt.Println("Minimum:  ", spaceMinimum)

	candidates := []int{}

	for _, v := range SZ {
		if v >= spaceMinimum {
			candidates = append(candidates, v)
		}
		if v <= 100000 {
			p1 += v
		}
	}

	sort.Ints(candidates)

	p2 = candidates[0]

	fmt.Println(p1)
	fmt.Println(p2)
}

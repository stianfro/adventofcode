package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Calories []int
}

func main() {
	var elves []Elf
	var totalCalories []int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		println("failed to read input")
	}

	data := strings.Split(string(input), "\n")

	var foodSlice []int

	for _, food := range data {
		if food == "" {
			elf := Elf{Calories: foodSlice}
			elves = append(elves, elf)

			foodSlice = []int{}
			continue
		}
		foodValue, _ := strconv.Atoi(food)
		foodSlice = append(foodSlice, foodValue)
	}

	for _, elf := range elves {
		sumCalories := 0

		for _, food := range elf.Calories {
			sumCalories += food
		}

		totalCalories = append(totalCalories, sumCalories)
	}

	sort.Ints(totalCalories)

	println(totalCalories[len(totalCalories)-1])
}

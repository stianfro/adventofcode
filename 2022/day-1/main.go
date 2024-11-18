package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Elf struct {
	Calories []int
}

func main() {
	var elves []Elf
	var totalCalories []int

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

	finalSum := 0

	for i := 1; i <= 3; i++ {
		finalSum += totalCalories[len(totalCalories)-i]
	}

	println(finalSum)
}

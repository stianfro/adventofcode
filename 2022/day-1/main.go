package main

import "sort"

type Elf struct {
	Calories []int
}

func main() {
	elf1 := Elf{
		Calories: []int{
			1000,
			2000,
			3000,
		},
	}
	elf2 := Elf{
		Calories: []int{
			4000,
		},
	}
	elf3 := Elf{
		Calories: []int{
			5000,
			6000,
		},
	}
	elf4 := Elf{
		Calories: []int{
			7000,
			8000,
			9000,
		},
	}
	elf5 := Elf{
		Calories: []int{
			10000,
		},
	}

	elves := []Elf{
		elf1,
		elf2,
		elf3,
		elf4,
		elf5,
	}

	var totalCalories []int

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

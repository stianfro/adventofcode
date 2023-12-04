package main

import (
	"os"
	"slices"
	"strings"
)

type Card struct {
	WinningNumbers []string
	NumbersYouHave []string
}

func main() {
	var allCards []Card

	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")

	for _, entry := range data {
		lists := strings.Split(entry, ": ")[1]
		listSlice := strings.Split(lists, " | ")
		listA := strings.Split(listSlice[0], " ")
		listB := strings.Split(listSlice[1], " ")

		listA = deleteEmpty(listA)
		listB = deleteEmpty(listB)

		card := Card{
			WinningNumbers: listA,
			NumbersYouHave: listB,
		}

		allCards = append(allCards, card)
	}

	var totalPoints int

	for _, card := range allCards {
		var points int
		var matchCount int

		for _, entry := range card.NumbersYouHave {

			if slices.Contains(card.WinningNumbers, entry) {
				matchCount += 1
			}

		}

		for i := 0; i < matchCount; i++ {
			if i == 0 {
				points += 1
			} else {
				points = points * 2
			}
		}

		totalPoints += points
	}

	println(totalPoints)
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Card struct {
	Instances      int
	WinningNumbers []string
	NumbersYouHave []string
}

func main() {
	var allCards []Card
	var sumCards int

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
			Instances:      1,
		}

		allCards = append(allCards, card)
	}

	for n, card := range allCards {
		for i := 0; i < card.Instances; i++ {
			var matchingNumbers int

			for _, entry := range card.NumbersYouHave {
				if slices.Contains(card.WinningNumbers, entry) {
					matchingNumbers += 1
				}
			}

			for i := 0; i < matchingNumbers; i++ {
				allCards[n+i+1].Instances += 1
			}
		}
		fmt.Println("Card "+fmt.Sprint(n+1), "instances:", card.Instances)
		sumCards += card.Instances
	}

	fmt.Println("Answer:", sumCards)
}

func PartOne(allCards []Card) {
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

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Bid      int
	Strength int
	Rank     int
	Type     string
	Cards    []string
}

func main() {
	var allHands []Hand
	// var allHandsStrength []Hand
	var fiveKind, fourKind, fullHouse, threeKind, twoPair, onePair, highCard []Hand

	strength := make(map[string]int)
	strength["A"] = 14
	strength["K"] = 13
	strength["Q"] = 12
	strength["J"] = 11
	strength["T"] = 10
	strength["9"] = 9
	strength["8"] = 8
	strength["7"] = 7
	strength["6"] = 6
	strength["5"] = 5
	strength["4"] = 4
	strength["3"] = 3
	strength["2"] = 2

	input, _ := os.ReadFile("input-example.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var hand Hand

		if line == "" {
			continue
		}

		s := strings.Split(line, " ")
		cards, bid := s[0], s[1]

		hand.Cards = strings.Split(cards, "")
		hand.Bid, _ = strconv.Atoi(bid)

		allHands = append(allHands, hand)
	}

	// determine type
	for index, hand := range allHands {
		count := make(map[string]int)
		for _, card := range hand.Cards {
			count[card]++
		}

		var occurences []int

		for _, v := range count {
			occurences = append(occurences, v)
		}

		allHands[index].Type = determineType(occurences)

		switch allHands[index].Type {
		case "Five of a kind":
			fiveKind = append(fiveKind, allHands[index])
		case "Four of a kind":
			fourKind = append(fourKind, allHands[index])
		case "Full house":
			fullHouse = append(fullHouse, allHands[index])
		case "Three of a kind":
			threeKind = append(threeKind, allHands[index])
		case "Two pair":
			twoPair = append(twoPair, allHands[index])
		case "One pair":
			onePair = append(onePair, allHands[index])
		case "High card":
			highCard = append(highCard, allHands[index])
		}
	}

	// pit hands up against eachother
	// not necessary?
	for a, hand := range twoPair {
		var intHand []int
		for _, card := range hand.Cards {
			intHand = append(intHand, strength[card])
		}

		for b, versusHand := range twoPair {
			if b == a {
				continue
			}

			if hand.Type == "" {
				continue
			}

			var intVersusHand []int
			for _, versusCard := range versusHand.Cards {
				intVersusHand = append(intVersusHand, strength[versusCard])
			}

			fmt.Println("Fight!")
			fmt.Println(hand.Cards, "vs", versusHand.Cards)
			fmt.Println(intHand, "vs", intVersusHand)

			// todo: find out how to handle rank
			for i := 0; i < 5; i++ {
				if intHand[i] > intVersusHand[i] {
					fmt.Println(hand.Cards, "wins!")
					break
				}
				if intVersusHand[i] > intHand[i] {
					fmt.Println(versusHand.Cards, "wins!")
					break
				}
			}
			fmt.Println("")
		}
	}
}

func determineType(i []int) string {
	sort.Ints(i)

	if len(i) == 1 {
		return "Five of a kind"
	}
	if len(i) == 2 {
		if i[1] == 4 {
			return "Four of a kind"
		} else {
			return "Full house"
		}
	}
	if len(i) == 3 {
		if i[2] == 3 {
			return "Three of a kind"
		} else {
			return "Two pair"
		}
	}
	if len(i) == 4 {
		return "One pair"
	}

	return "High card"
}

package main

import (
	"fmt"
	"os"
	"strings"
)

// Rock > Scissors
// Scissors > Paper
// Paper > Rock

// A/X = Rock     (1)
// B/Y = Paper    (2)
// C/Z = Scissors (3)

// Lose = 0
// Draw = 3
// Won  = 6

var result string
var score int

func rockPaperScissors(opponent string, player string) (result string, score int) {
	switch opponent {
	case "A":
		switch player {
		case "X":
			// Draw
			score = 3 + 1
			result = "Draw"
		case "Y":
			// Win
			score = 6 + 2
			result = "Player wins, opponent loses"
		case "Z":
			// Lose
			score = 0 + 3
			result = "Player loses, opponent wins"
		}
	case "B":
		switch player {
		case "X":
			// Lose
			score = 0 + 1
			result = "Player loses, opponent wins"
		case "Y":
			// Draw
			score = 3 + 2
			result = "Draw"
		case "Z":
			// Win
			score = 6 + 3
			result = "Player wins, opponent loses"
		}
	case "C":
		switch player {
		case "X":
			// Win
			score = 6 + 1
			result = "Player wins, opponent loses"
		case "Y":
			// Lose
			score = 0 + 2
			result = "Player loses, opponent wins"
		case "Z":
			// Draw
			score = 3 + 3
			result = "Draw"
		}
	}

	return result, score
}

func main() {
	totalScore := 0
	input := [][]string{}

	inputBytes, err := os.ReadFile("input-example.txt")
	if err != nil {
		println("failed to read input")
	}

	inputData := strings.Split(string(inputBytes), "\n")

	for _, v := range inputData {
		round := strings.Split(v, " ")

		input = append(input, round)
	}

	for i := range input {
		result, score = rockPaperScissors(input[i][0], input[i][1])
		fmt.Println(result+". Score:", score)

		totalScore += score
	}

	fmt.Println("Total score:", totalScore)
}

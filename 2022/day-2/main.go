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

// X = Need to lose
// Y = Need to draw
// Z = Need to win

var score int

func rockPaperScissors(opponent string, player string) (score int) {
	switch opponent {
	case "A":
		switch player {
		case "X":
			// Draw
			score = 3 + 1
		case "Y":
			// Win
			score = 6 + 2
		case "Z":
			// Lose
			score = 0 + 3
		}
	case "B":
		switch player {
		case "X":
			// Lose
			score = 0 + 1
		case "Y":
			// Draw
			score = 3 + 2
		case "Z":
			// Win
			score = 6 + 3
		}
	case "C":
		switch player {
		case "X":
			// Win
			score = 6 + 1
		case "Y":
			// Lose
			score = 0 + 2
		case "Z":
			// Draw
			score = 3 + 3
		}
	}

	return score
}

func rockPaperScissorsCorrect(opponent string, player string) (score int) {
	switch player {
	case "X": // If should lose
		switch opponent {
		case "A":
			// Rock/Scissors
			score = 0 + 3
		case "B":
			// Paper/Rock
			score = 0 + 1
		case "C":
			// Scissors/Paper
			score = 0 + 2
		}
	case "Y": // If should draw
		switch opponent {
		case "A":
			// Rock/Rock
			score = 3 + 1
		case "B":
			// Paper/Paper
			score = 3 + 2
		case "C":
			// Scissors/Scissors
			score = 3 + 3
		}
	case "Z": // If should win
		switch opponent {
		case "A":
			// Rock/Paper
			score = 6 + 2
		case "B":
			// Paper/Scissors
			score = 6 + 3
		case "C":
			// Scissors/Rock
			score = 6 + 1
		}
	}
	return score
}

func main() {
	var totalScore int
	input := [][]string{}

	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		println("failed to read input")
	}

	inputData := strings.Split(string(inputBytes), "\n")

	for _, v := range inputData {
		round := strings.Split(v, " ")

		input = append(input, round)
	}

	fmt.Println("--- Part One ---")
	totalScore = 0
	for i := range input {
		score = rockPaperScissors(input[i][0], input[i][1])
		totalScore += score
	}
	fmt.Println("Total score:", totalScore)
	fmt.Println("")

	fmt.Println("--- Part Two ---")
	totalScore = 0
	for i := range input {
		score = rockPaperScissorsCorrect(input[i][0], input[i][1])
		totalScore += score
	}
	fmt.Println("Correct total score:", totalScore)
	fmt.Println("")
}

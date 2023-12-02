package main

import (
	"os"
	"strconv"
	"strings"
)

var letterDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	var allDigits int

	input, _ := os.ReadFile("input.txt")
	calibrationDocument := strings.Split(string(input), "\n")

	for _, line := range calibrationDocument {
		var digitSlice []string

		convertedLine := convertLetters(line)

		data := strings.Split(convertedLine, "")
		for _, character := range data {
			if isDigit(character) {
				digitSlice = append(digitSlice, character)
			}
		}

		digitFirst := digitSlice[0]
		digitLast := ""
		if len(digitSlice) == 1 {
			digitLast = digitFirst
		} else {
			digitLast = digitSlice[len(digitSlice)-1]
		}

		digitString := digitFirst + digitLast

		digitInteger, _ := strconv.Atoi(digitString)
		allDigits += digitInteger
	}
	println(allDigits)
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func letterDigit(s string) string {
	switch s {
	case "one":
		return "o1e"
	case "two":
		return "t2o"
	case "three":
		return "th3ee"
	case "four":
		return "fo4r"
	case "five":
		return "fi5e"
	case "six":
		return "s6x"
	case "seven":
		return "se7en"
	case "eight":
		return "ei8ht"
	case "nine":
		return "ni9e"
	}

	return ""
}

func convertLetters(s string) string {
	for _, letter := range letterDigits {
		if strings.Contains(s, letter) {
			digit := letterDigit(letter)
			s = strings.ReplaceAll(s, letter, digit)
		}
	}

	return s
}

package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	var allDigits int

	input, _ := os.ReadFile("input.txt")
	calibrationDocument := strings.Split(string(input), "\n")

	for _, line := range calibrationDocument {
		var digitSlice []string

		data := strings.Split(line, "")
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

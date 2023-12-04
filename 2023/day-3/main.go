package main

import (
	"os"
	"regexp"
	"strings"
)

var symbols = "*&-%+=@#/$"

func main() {
	var longLine []string

	input, _ := os.ReadFile("input-test.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var number string

		for i := range line {
			character := string(line[i])

			if i == len(line)-1 {
				continue
			}

			nextCharacter := string(line[i+1])

			if strings.ContainsAny(character, symbols) || character == "." {
				longLine = append(longLine, character)
				continue
			}

			number += character

			if strings.ContainsAny(nextCharacter, symbols) || nextCharacter == "." {
				longLine = append(longLine, number)
				number = ""
				continue
			}
		}
	}

	for i, entry := range longLine {
		isNumber, _ := regexp.MatchString(`\d`, entry)
		isSymbol := strings.ContainsAny(entry, symbols)

		if isSymbol || entry == "." {
			continue
		}

		if isNumber {
			println("num: ", entry)

			println(" +8: ", longLine[i+8])
			println(" +9: ", longLine[i+9])
			println("+10: ", longLine[i+10])

			println("---")

			// if strings.ContainsAny(longLine[i+10], symbols) {
			// 	println(entry + " is adjacent to symbol " + longLine[i+10])
			// }
			// if strings.ContainsAny(longLine[i+10], symbols) {
			// 	println(entry + " is adjacent to symbol " + longLine[i+10])
			// }
		}
	}
}

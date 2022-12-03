package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}

	return strings.Split(string(inputBytes), "\n")
}

func genAlphabet() []string {
	alphabet := []string{}

	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	return alphabet
}

func main() {
	inputData := getInput("input-example.txt")
	// priority := genAlphabet()

	for _, v := range inputData {
		r := strings.Split(v, "")
		fmt.Println(r)
		c1 := r[0 : len(r)-len(r)/2]
		fmt.Println(c1)
		c2 := r[len(r)-len(r)/2:]
		fmt.Println(c2)

		d := []string{}

		for _, i1 := range c1 {
			if len(d) > 0 {
				break
			}
			for _, i2 := range c2 {
				if i1 == i2 {
					d = append(d, i1)
					break
				}
			}
		}

		fmt.Println(d)
	}
}

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
	d := getInput("input.txt")
	a := genAlphabet()

	fmt.Println("--- Part One ---")
	t := 0

	for _, v := range d {
		r := strings.Split(v, "")

		c1 := r[0 : len(r)-len(r)/2]
		c2 := r[len(r)-len(r)/2:]

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

		s := 0

		for i, v := range a {
			p := i + 1

			if v == d[0] {
				s += p
			}
		}

		fmt.Println(s)

		t += s
	}
	fmt.Println(t)

	fmt.Println("--- Part Two ---")
	x := 0
	g := []string{}
	G := [][]string{}

	for _, v := range d {
		x += 1

		g = append(g, v)

		if x >= 3 {
			x = 0
			G = append(G, g)
			g = []string{}
		}
	}

	i := []string{}

	for _, g := range G {
		r1 := strings.Split(g[0], "")
		r2 := strings.Split(g[1], "")
		r3 := strings.Split(g[2], "")

		for _, a := range r1 {
			for _, b := range r2 {
				if a == b {
					for _, c := range r3 {
						if b == c {
							x += 1
							break
						}
					}
					break
				}
			}
			if x > 0 {
				x = 0
				i = append(i, a)
				break
			}
		}
	}

	t = 0

	for _, x := range i {
		s := 0

		for idx, y := range a {
			p := idx + 1

			if y == x {
				s += p
			}
		}
		t += s
	}

	fmt.Println(t)
}

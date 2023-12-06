package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input-example.txt")
	data := strings.Split(string(input), "\n\n")

	seeds := []int{79, 14, 55, 13}
	fmt.Println(seeds)

	for i, line := range data {
		if i == 0 {
			continue
		}

		x := strings.Split(line, ":\n")
		y := strings.Split(x[1], "\n")

		for _, v := range y {
			var dst, src, sz int

			if v == "" {
				continue
			}

			z := strings.Split(v, " ")

			dst, _ = strconv.Atoi(z[0])
			src, _ = strconv.Atoi(z[1])
			sz, _ = strconv.Atoi(z[2])

			fmt.Println(dst, src, sz)
		}
	}
}

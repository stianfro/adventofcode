package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// var input = []string{
// 	"2x3x4",
// 	"1x1x10",
// }

func main() {
	var wrappingPaperTotal int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("error reading")
	}
	data := strings.Split(string(input), "\n")

	for i, box := range data {
		if len(box) <= 0 {
			continue
		}

		var smallestSide int
		var wrappingPaper int

		lwh := strings.Split(box, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		sideLW := l * w
		sideWH := w * h
		sideHL := h * l

		sides := []int{
			sideLW,
			sideWH,
			sideHL,
		}

		slices.Sort(sides)
		smallestSide = sides[0]

		wrappingPaper = 2*sideLW + 2*sideWH + 2*sideHL + smallestSide
		log.Println(i, wrappingPaper)

		wrappingPaperTotal += wrappingPaper
	}

	log.Println("total:", wrappingPaperTotal)
}

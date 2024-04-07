package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("error reading")
	}
	data := strings.Split(string(input), "\n")
	// var data = []string{
	// 	"2x3x4",
	// 	"1x1x10",
	// }

	partOne(data)
	partTwo(data)
}

func partTwo(boxes []string) {
	var ribbonTotal int

	for _, box := range boxes {
		if len(box) <= 0 {
			continue
		}

		lwh := strings.Split(box, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		lwhSorted := []int{l, w, h}
		slices.Sort(lwhSorted)

		volume := l * w * h
		perimeter := 2*lwhSorted[0] + 2*lwhSorted[1]

		ribbon := perimeter + volume
		ribbonTotal += ribbon
	}

	log.Println("p2:", ribbonTotal)
}

func partOne(boxes []string) {
	var wrappingPaperTotal int

	for _, box := range boxes {
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
		wrappingPaperTotal += wrappingPaper
	}

	log.Println("p1:", wrappingPaperTotal)
}

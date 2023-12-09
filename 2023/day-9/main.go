package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type History struct {
	Sequences [][]int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")

	var allHistory []History

	for _, line := range data {
		var history History
		var seq []int

		if line == "" {
			continue
		}

		s := strings.Split(line, " ")
		for _, v := range s {
			vI, _ := strconv.Atoi(v)
			seq = append(seq, vI)
		}
		history.Sequences = append(history.Sequences, seq)

		for i := 0; i < len(history.Sequences); i++ {
			var sumSeq int
			var diffs []int

			for b := range history.Sequences[i] {
				if b == 0 {
					continue
				}
				diff := history.Sequences[i][b] - history.Sequences[i][b-1]
				sumSeq += diff
				diffs = append(diffs, diff)
			}

			history.Sequences = append(history.Sequences, diffs)

			if sumSeq == 0 {
				break
			}
		}
		allHistory = append(allHistory, parseHistory(history))
	}

	var answer int
	for _, h := range allHistory {
		answer += h.Sequences[0][len(h.Sequences[0])-1]
	}
	fmt.Println(answer)
}

func parseHistory(h History) History {
	seqs := h.Sequences
	seqs[len(seqs)-1] = append(seqs[len(seqs)-1], 0)

	for i := range seqs {
		if i == len(seqs)-1 {
			continue
		}
		x := seqs[len(seqs)-(i+1)]
		y := seqs[len(seqs)-(i+2)]
		a := x[len(x)-1] + y[len(y)-1]

		seqs[len(seqs)-(i+2)] = append(seqs[len(seqs)-(i+2)], a)
	}

	return h
}

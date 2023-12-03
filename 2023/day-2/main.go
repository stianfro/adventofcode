package main

import (
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

const (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func main() {
	var possibleGames []Game
	var sumIDs int

	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var game Game

		g := strings.Split(line, ": ")
		id := strings.Split(g[0], " ")
		sets := g[1]
		game.ID, _ = strconv.Atoi(id[len(id)-1])
		setSplit := strings.Split(sets, "; ")

		for i := range setSplit {
			setParsed := parseSet(setSplit[i])
			game.Sets = append(game.Sets, setParsed)
		}

		if possible(game) {
			possibleGames = append(possibleGames, game)
		}
	}

	for _, game := range possibleGames {
		sumIDs += game.ID
	}
	println(sumIDs)
}

func parseSet(s string) Set {
	var set Set
	sets := strings.Split(s, ", ")

	for i := range sets {
		cube := strings.Split(sets[i], " ")
		cubeCount, _ := strconv.Atoi(cube[0])

		switch cube[1] {
		case "red":
			set.Red = cubeCount
		case "green":
			set.Green = cubeCount
		case "blue":
			set.Blue = cubeCount
		}
	}

	return set
}

func possible(g Game) bool {
	for _, set := range g.Sets {
		if set.Red > redLimit {
			return false
		}
		if set.Green > greenLimit {
			return false
		}
		if set.Blue > blueLimit {
			return false
		}
	}

	return true
}

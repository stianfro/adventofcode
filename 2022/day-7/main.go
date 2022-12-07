package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) []string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		println("failed to read input")
	}

	return strings.Split(string(inputBytes), "\n")
}

type Filesystem struct {
	Directories []Directory
}

type Directory struct {
	Name  string
	Level int
	Size  int
}

func main() {
	var dir []Directory
	var lvl int
	var cur int
	var dirSize int

	input := getInput("input-example.txt")

	for _, v := range input {
		output := strings.Split(v, " ")

		if output[0] == "$" {
			switch output[1] {
			case "cd":
				if output[2] == ".." {
					dir[cur] = Directory{
						Name:  "foo",
						Level: lvl,
						Size:  dirSize,
					}
					lvl -= 1
				} else if output[2] != "/" {
					cur += 1
					lvl += 1
				}
			case "ls": // $ ls
				continue
			}
		} else if output[0] == "dir" {
			continue
		} else { // files
			size, _ := strconv.Atoi(output[0])
			dirSize += size
		}
	}

	fmt.Println(dir)
}

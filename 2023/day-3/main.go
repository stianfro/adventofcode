package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// var symbols = "*&-%+=@#/$"
var symbols = "*"

func main() {
	var allLines [][]string

	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var lineSlice []string

		for _, character := range line {
			lineSlice = append(lineSlice, string(character))
		}

		allLines = append(allLines, lineSlice)
	}

	var sumRatios int

	for a, line := range allLines {
		for b, char := range line {
			if strings.ContainsAny(char, symbols) {
				var partNumbers []string
				var partCount int

				// left
				if ok, _ := regexp.MatchString(`\d`, line[b-1]); ok {
					var numberToAdd string
					numberToAdd += line[b-1]

					if ok, _ := regexp.MatchString(`\d`, line[b-2]); ok {
						numberToAdd += line[b-2]

						if ok, _ := regexp.MatchString(`\d`, line[b-3]); ok {
							numberToAdd += line[b-3]
						}
					}
					partNumbers = append(partNumbers, reverse(numberToAdd))
					partCount += 1
				}
				// right
				if ok, _ := regexp.MatchString(`\d`, line[b+1]); ok {
					var numberToAdd string
					numberToAdd += line[b+1]

					if ok, _ := regexp.MatchString(`\d`, line[b+2]); ok {
						numberToAdd += line[b+2]

						if ok, _ := regexp.MatchString(`\d`, line[b+3]); ok {
							numberToAdd += line[b+3]
						}
					}
					partNumbers = append(partNumbers, numberToAdd)
					partCount += 1
				}

				// up
				if a >= 1 {
					// up left
					if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-1]); ok {
						if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b]); !ok {
							var numberToAdd string
							numberToAdd += allLines[a-1][b-1]

							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-2]); ok {
								numberToAdd += allLines[a-1][b-2]

								if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-3]); ok {
									numberToAdd += allLines[a-1][b-3]
								}
							}
							partNumbers = append(partNumbers, reverse(numberToAdd))
							partCount += 1
						}
					}
					// up middle
					if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b]); ok {
						var numberToAdd string
						numberToAdd = allLines[a-1][b]
						if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-1]); ok {
							numberToAdd = allLines[a-1][b-1] + allLines[a-1][b]
							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+1]); ok {
								numberToAdd = allLines[a-1][b-1] + allLines[a-1][b] + allLines[a-1][b+1]
							}
							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-2]); ok {
								numberToAdd = allLines[a-1][b-2] + allLines[a-1][b-1] + allLines[a-1][b]
							}
						}
						if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+1]); ok {
							numberToAdd = allLines[a-1][b] + allLines[a-1][b+1]
							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b-1]); ok {
								numberToAdd = allLines[a-1][b-1] + allLines[a-1][b] + allLines[a-1][b+1]
							}
							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+2]); ok {
								numberToAdd = allLines[a-1][b] + allLines[a-1][b+1] + allLines[a-1][b+2]
							}
						}
						partNumbers = append(partNumbers, numberToAdd)
						partCount += 1
					}
					// up right
					if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+1]); ok {
						if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b]); !ok {
							var numberToAdd string
							numberToAdd += allLines[a-1][b+1]

							if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+2]); ok {
								numberToAdd += allLines[a-1][b+2]

								if ok, _ := regexp.MatchString(`\d`, allLines[a-1][b+3]); ok {
									numberToAdd += allLines[a-1][b+3]
								}
							}
							partNumbers = append(partNumbers, numberToAdd)
							partCount += 1
						}
					}
				}

				// down
				if a <= len(allLines)-1 {
					// down left
					if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-1]); ok {
						if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b]); !ok {
							var numberToAdd string
							numberToAdd += allLines[a+1][b-1]

							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-2]); ok {
								numberToAdd += allLines[a+1][b-2]

								if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-3]); ok {
									numberToAdd += allLines[a+1][b-3]
								}
							}
							partNumbers = append(partNumbers, reverse(numberToAdd))
							partCount += 1
						}
					}
					// down middle
					if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b]); ok {
						var numberToAdd string
						numberToAdd = allLines[a+1][b]
						if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-1]); ok {
							numberToAdd = allLines[a+1][b-1] + allLines[a+1][b]
							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+1]); ok {
								numberToAdd = allLines[a+1][b-1] + allLines[a+1][b] + allLines[a+1][b+1]
							}
							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-2]); ok {
								numberToAdd = allLines[a+1][b-2] + allLines[a+1][b-1] + allLines[a+1][b]
							}
						}
						if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+1]); ok {
							numberToAdd = allLines[a+1][b] + allLines[a+1][b+1]
							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b-1]); ok {
								numberToAdd = allLines[a+1][b-1] + allLines[a+1][b] + allLines[a+1][b+1]
							}
							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+2]); ok {
								numberToAdd = allLines[a+1][b] + allLines[a+1][b+1] + allLines[a+1][b+2]
							}
						}

						partNumbers = append(partNumbers, numberToAdd)
						partCount += 1
					}
					// down right
					if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+1]); ok {
						if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b]); !ok {
							var numberToAdd string
							numberToAdd += allLines[a+1][b+1]

							if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+2]); ok {
								numberToAdd += allLines[a+1][b+2]

								if ok, _ := regexp.MatchString(`\d`, allLines[a+1][b+3]); ok {
									numberToAdd += allLines[a+1][b+3]
								}
							}
							partNumbers = append(partNumbers, numberToAdd)
							partCount += 1
						}
					}
				}

				if partCount == 2 {
					partA, _ := strconv.Atoi(partNumbers[0])
					partB, _ := strconv.Atoi(partNumbers[1])
					ratio := partA * partB
					sumRatios += ratio
				}
			}
		}
	}

	println(sumRatios)
}

func reverse(s string) string {
	r := make([]byte, 0, len(s))
	for len(s) > 0 {
		_, n := utf8.DecodeLastRuneInString(s)
		i := len(s) - n
		r = append(r, s[i:]...)
		s = s[:i]
	}
	return string(r)
}

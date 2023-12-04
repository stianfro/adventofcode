package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var symbols = "*&-%+=@#/$"

func main() {
	var allLines [][]string
	var partNumbers []string

	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n")

	for _, line := range data {
		var lineSlice []string

		for _, character := range line {
			lineSlice = append(lineSlice, string(character))
		}

		allLines = append(allLines, lineSlice)
	}

	for a, line := range allLines {
		for b, char := range line {
			if strings.ContainsAny(char, symbols) {
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
						}
					}
				}
			}
		}
	}

	var answer int

	for _, n := range partNumbers {
		println(n)

		number, _ := strconv.Atoi(n)
		answer += number
	}

	println(answer)
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

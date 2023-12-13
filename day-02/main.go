package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func count(re *regexp.Regexp, s string) int {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return 0
	}

	value, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}

	return value
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	blueRegex := regexp.MustCompile(`(\d+) blue`)
	redRegex := regexp.MustCompile(`(\d+) red`)
	greenRegex := regexp.MustCompile(`(\d+) green`)

	lines := strings.Split(string(dat), "\n")
	sumPart1 := 0
	sumPart2 := 0
	for lineIdx, line := range lines[:len(lines)-1] {
		lineValidPart1 := true
		maxPerColor := map[string]int{"blue": 0, "red": 0, "green": 0}
		for _, game := range strings.Split(line, ";") {
			red := count(redRegex, game)
			green := count(greenRegex, game)
			blue := count(blueRegex, game)

			if red > 12 || green > 13 || blue > 14 {
				lineValidPart1 = false
			}

			maxPerColor["red"] = max(red, maxPerColor["red"])
			maxPerColor["green"] = max(green, maxPerColor["green"])
			maxPerColor["blue"] = max(blue, maxPerColor["blue"])
		}

		if lineValidPart1 {
			sumPart1 += lineIdx + 1
		}

		sumPart2 += maxPerColor["red"] * maxPerColor["green"] * maxPerColor["blue"]
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", sumPart1, sumPart2)
}

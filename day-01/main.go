package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// part one
	digitRegex := regexp.MustCompile(`\d`)
	sumPartOne := 0
	for _, line := range lines[:len(lines)-1] {
		matches := digitRegex.FindAllString(line, -1)
		value, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			panic(err)
		}
		sumPartOne += value
	}
	fmt.Printf("Part 1: %d\n", sumPartOne)

	// part two
	digitSpelledRegex := regexp.MustCompile(`\d|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	sumPartTwo := 0
	for _, line := range lines[:len(lines)-1] {
		matches := digitSpelledRegex.FindAllString(line, -1)
		value, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			panic(err)
		}
		sumPartTwo += value
	}
	fmt.Printf("Part 2: %d\n", sumPartTwo)
}

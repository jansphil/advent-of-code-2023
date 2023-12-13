package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseNumbers(input []string) []int {
	out := make([]int, len(input))
	for i, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		out[i] = num
	}
	return out
}

func parseLine(line string) ([]int, []int) {
	digitsRegexp := regexp.MustCompile(`\d+`)

	parts := strings.Split(line, "|")
	numbers := parseNumbers(digitsRegexp.FindAllString(parts[1], -1))
	winningNumbers := parseNumbers(digitsRegexp.FindAllString(strings.Split(parts[0], ":")[1], -1))

	return winningNumbers, numbers
}

func isIn[t comparable](probe t, array []t) bool {
	for _, v := range array {
		if v == probe {
			return true
		}
	}

	return false
}

func main() {
	dat, err := os.ReadFile("input.txt")
	// dat, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		count := -1
		winningNumbers, numbers := parseLine(line)
		for _, num := range numbers {
			if isIn(num, winningNumbers) {
				count++
			}
		}
		sum += int(math.Pow(2, float64(count)))
	}
	fmt.Printf("Part 1: %d\n", sum)

	cardCounts := make(map[int]int)
	sumPartTwo := 0
	for i, line := range lines[:len(lines)-1] {
		cardCounts[i] = cardCounts[i] + 1

		count := 0
		winningNumbers, numbers := parseLine(line)
		for _, num := range numbers {
			if isIn(num, winningNumbers) {
				count++
			}
		}

		for j := 1; j <= count; j++ {
			cardCounts[i+j] = cardCounts[i+j] + cardCounts[i]
		}

		sumPartTwo += cardCounts[i]
	}
	fmt.Println(cardCounts)
	fmt.Printf("Part 2: %d\n", sumPartTwo)
}

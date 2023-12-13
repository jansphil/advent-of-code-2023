package utils

import "strconv"

// parse string to int
func ParseInt(input string) int {
	val, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return val
}

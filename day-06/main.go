package main

import (
	"fmt"
	"math"
)

func calc(record float64, time float64) int {
	min := time/2 - math.Sqrt(math.Pow(time/2, 2)-record)
	max := time/2 + math.Sqrt(math.Pow(time/2, 2)-record)

	return int(math.Floor(max)) - int(math.Floor(min))
}

func main() {
	fmt.Println(calc(478, 58) *
		calc(2232, 99) *
		calc(1019, 64) *
		calc(1071, 69))
	fmt.Println(calc(478223210191071, 58996469))
}

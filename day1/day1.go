package day1

import (
	"advent-of-code-2019/shared"
	"math"
	"strconv"
)

func Problem1_Answer() int {
	inputs := shared.GetInput("day1/input")
	var sum int
	for _, element := range inputs {
		parsedInt, _ := strconv.ParseInt(element, 10, 32)
		sum += CalcFuel(int(parsedInt))
	}
	return sum
}

func CalcFuel(mass int) int {
	divided := float64(mass) / 3
	floored := math.Floor(divided)
	subtracted := floored - 2
	return int(subtracted)
}

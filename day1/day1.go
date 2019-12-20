package day1

import (
	"advent-of-code-2019/shared"
	"math"
	"strconv"
)

func Problem1Answer() int {
	return CreateAnswer(CalcFuel)
}

func Problem2Answer() int {
	return CreateAnswer(CalcFuelCompound)
}

func CalcFuel(mass int) int {
	divided := float64(mass) / 3
	floored := math.Floor(divided)
	subtracted := floored - 2
	return int(subtracted)
}

func CalcFuelCompound(mass int) int {
	var fuelCompound int

	initialFuel := CalcFuel(mass)

	for remainingFuel := CalcFuel(initialFuel); remainingFuel > 0; remainingFuel = CalcFuel(remainingFuel) {
		fuelCompound += remainingFuel
	}

	return initialFuel + fuelCompound
}

func CreateAnswer(f func(mass int) int) int {
	inputs := shared.GetInput("day1/input")
	var sum int
	for _, element := range inputs {
		parsedInt, _ := strconv.ParseInt(element, 10, 64)
		sum += f(int(parsedInt))
	}
	return sum
}

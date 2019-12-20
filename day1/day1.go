package day1

import (
	"advent-of-code-2019/shared"
	"math"
	"strconv"
)

func Problem1Answer() int64 {
	return CreateAnswer(CalcFuel)
}

func Problem2Answer() int64 {
	return CreateAnswer(CalcFuelCompound)
}

func CalcFuel(mass int64) int64 {
	divided := float64(mass) / 3
	floored := math.Floor(divided)
	subtracted := floored - 2
	return int64(subtracted)
}

func CalcFuelCompound(mass int64) int64 {
	var fuelCompound int64

	initialFuel := CalcFuel(mass)

	for remainingFuel := CalcFuel(initialFuel); remainingFuel > 0; remainingFuel = CalcFuel(remainingFuel) {
		fuelCompound += remainingFuel
	}

	return initialFuel + fuelCompound
}

func CreateAnswer(f func(mass int64) int64) int64 {
	inputs := shared.GetInputLines("day1/input")
	var sum int64
	for _, element := range inputs {
		parsedint64, _ := strconv.ParseInt(element, 10, 64)
		sum += f(int64(parsedint64))
	}
	return sum
}

package day2

import (
	"advent-of-code-2019/shared"
	"strconv"
	"strings"
)

func Problem1Answer() int64 {
	input := shared.GetInput("day2/input")

	initial := DeserializeIntcode(input)
	initial[1] = 12
	initial[2] = 2

	final := RunIntcode(initial)

	return final[0]
}

func RunIntcode(initial []int64) []int64 {
	state := initial
loop:
	for i := 0; ; i += 4 {
		var op func(x, y int64) int64
		switch state[i] {
		case 99:
			break loop
		case 1:
			op = func(x, y int64) int64 { return x + y }
		case 2:
			op = func(x, y int64) int64 { return x * y }
		}
		px := state[i+1]
		py := state[i+2]
		x := state[px]
		y := state[py]
		ps := state[i+3]
		state[ps] = op(x, y)
	}

	return state
}

func DeserializeIntcode(intcode string) []int64 {
	split := strings.Split(intcode, ",")

	numbers := make([]int64, len(split))
	for i, e := range split {
		number, _ := strconv.ParseInt(e, 10, 64)
		numbers[i] = number
	}

	return numbers
}

func SerializeIntcode(intcode []int64) string {
	raws := make([]string, len(intcode))
	for i, e := range intcode {
		raws[i] = strconv.FormatInt(e, 10)
	}

	joined := strings.Join(raws, ",")

	return joined
}

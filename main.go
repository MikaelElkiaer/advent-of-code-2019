package main

import (
	"advent-of-code-2019/day1"
	"advent-of-code-2019/day2"
	"advent-of-code-2019/day3"
	"advent-of-code-2019/day4"
	"fmt"
	"os"
)

func main() {
	var ans int64

	day, problem := os.Args[1], os.Args[2]
	switch day {
	case "1":
		switch problem {
		case "1":
			ans = day1.Problem1Answer()
		case "2":
			ans = day1.Problem2Answer()
		}
	case "2":
		switch problem {
		case "1":
			ans = day2.Problem1Answer()
		case "2":
			ans = day2.Problem2Answer()
		}
	case "3":
		switch problem {
		case "1":
			ans = day3.Problem1Answer()
		case "2":
			ans = day3.Problem2Answer()
		}
	case "4":
		switch problem {
		case "1":
			ans = day4.Problem1Answer()
		case "2":
			ans = day4.Problem2Answer()
		}
	}

	fmt.Printf("Day %s problem %s answer: %d\n", day, problem, ans)
}

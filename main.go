package main

import (
	"advent-of-code-2019/day1"
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
	}

	fmt.Printf("Day %s problem %s answer: %d\n", day, problem, ans)
}

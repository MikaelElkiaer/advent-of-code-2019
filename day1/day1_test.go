package day1

import (
	"fmt"
	"testing"
)

func TestCalcFuel(t *testing.T) {
	var tests = []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.mass)
		t.Run(testname, func(t *testing.T) {
			ans := CalcFuel(tt.mass)
			if ans != tt.want {
				t.Errorf("Got %d, want %d", ans, tt.want)
			}
		})
	}
}
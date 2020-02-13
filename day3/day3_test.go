package day3

import (
	"fmt"
	"testing"
)

func TestDistanceToClosestWireIntersectFromCentralPort(t *testing.T) {
	var tests = []struct {
		initial string
		want    int64
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.initial)
		t.Run(testname, func(t *testing.T) {
			wire1, wire2 := ParseWirePathsInput(tt.initial)
			ans := DistanceToClosestWireIntersectFromCentralPort(wire1, wire2)
			if ans != tt.want {
				t.Errorf("Got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestFewestStepsToWireIntersectFromCentralPort(t *testing.T) {
	var tests = []struct {
		initial string
		want    int64
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.initial)
		t.Run(testname, func(t *testing.T) {
			wire1, wire2 := ParseWirePathsInput(tt.initial)
			ans := FewestStepsToWireIntersectFromCentralPort(wire1, wire2)
			if ans != tt.want {
				t.Errorf("Got %d, want %d", ans, tt.want)
			}
		})
	}
}

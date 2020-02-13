package day3

import (
	"advent-of-code-2019/shared"
	"math"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type WirePathStep struct {
	Directio string
	Distance int64
}

type Position struct {
	X int64
	Y int64
}

func Problem1Answer() int64 {
	input := shared.GetInput("day3/input")
	wire1, wire2 := ParseWirePathsInput(input)

	distance := DistanceToClosestWireIntersectFromCentralPort(wire1, wire2)

	return distance
}

func Problem2Answer() int64 {
	input := shared.GetInput("day3/input")
	wire1, wire2 := ParseWirePathsInput(input)

	distance := FewestStepsToWireIntersectFromCentralPort(wire1, wire2)

	return distance
}

func FewestStepsToWireIntersectFromCentralPort(wire1 []WirePathStep, wire2 []WirePathStep) int64 {
	wire1Positions := CalcPositions(wire1)
	var fewestSteps int64 = math.MaxInt64
	wire2LatestPosition := Position{0, 0}
	for i := 0; i < len(wire2); i++ {
		wire2PreviousPosition := wire2LatestPosition
		wire2LatestPosition = CalcPosition(wire2[i], wire2LatestPosition)
		for j := 1; j < len(wire1Positions); j++ {
			wire1PreviousPosition := wire1Positions[j-1]
			wire1LatestPosition := wire1Positions[j]

			intersection, success := CalcIntersection(wire1PreviousPosition, wire1LatestPosition, wire2PreviousPosition, wire2LatestPosition)

			if success && !cmp.Equal(intersection, Position{0, 0}) {
				var steps int64 = 0

				for _, w := range wire2[:i] {
					steps += w.Distance
				}
				for _, w := range wire1[:j-1] {
					steps += w.Distance
				}

				steps += CalcDistance(wire2PreviousPosition, intersection)
				steps += CalcDistance(wire1PreviousPosition, intersection)

				if steps < fewestSteps {
					fewestSteps = steps
				}
			}
		}
	}

	return fewestSteps
}

func DistanceToClosestWireIntersectFromCentralPort(wire1 []WirePathStep, wire2 []WirePathStep) int64 {
	wire1Positions := CalcPositions(wire1)
	var closestIntersectDistance int64 = math.MaxInt64
	wire2LatestPosition := Position{0, 0}
	for i := 0; i < len(wire2); i++ {
		wire2PreviousPosition := wire2LatestPosition
		wire2LatestPosition = CalcPosition(wire2[i], wire2LatestPosition)
		for j := 1; j < len(wire1Positions); j++ {
			wire1PreviousPosition := wire1Positions[j-1]
			wire1LatestPosition := wire1Positions[j]

			distance := CalcIntersectionDistance(wire1PreviousPosition, wire1LatestPosition, wire2PreviousPosition, wire2LatestPosition)

			if 0 < distance && distance < closestIntersectDistance {
				closestIntersectDistance = distance
			}
		}
	}

	return closestIntersectDistance
}

func CalcIntersectionDistance(l1p1, l1p2, l2p1, l2p2 Position) int64 {
	intersection, success := CalcIntersection(l1p1, l1p2, l2p1, l2p2)

	if success {
		return CalcDistance(Position{0, 0}, intersection)
	}

	return math.MaxInt64
}

func CalcIntersection(l1p1, l1p2, l2p1, l2p2 Position) (Position, bool) {
	if l1p1.X == l1p2.X && l2p1.X == l2p2.X {

	} else if l1p1.Y == l1p2.Y && l2p1.Y == l2p2.Y {

	} else if l1p1.X == l1p2.X && l2p1.Y == l2p2.Y {
		if Min(l1p1.Y, l1p2.Y) <= l2p1.Y && l2p1.Y <= Max(l1p1.Y, l1p2.Y) && Min(l2p1.X, l2p2.X) <= l1p1.X && l1p1.X <= Max(l2p1.X, l2p2.X) {
			return Position{l1p1.X, l2p1.Y}, true
		}
	} else if l1p1.Y == l1p2.Y && l2p1.X == l2p2.X {
		if Min(l1p1.X, l1p2.X) <= l2p1.X && l2p1.X <= Max(l1p1.X, l1p2.X) && Min(l2p1.Y, l2p2.Y) <= l1p1.Y && l1p1.Y <= Max(l2p1.Y, l2p2.Y) {
			return Position{l2p1.X, l1p1.Y}, true
		}
	}

	return Position{0, 0}, false
}

func CalcDistance(p1, p2 Position) int64 {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if b < a {
		return a
	}
	return b
}

func CalcPositions(wireSteps []WirePathStep) []Position {
	wirePositions := make([]Position, len(wireSteps)+1)
	wirePositions[0] = Position{0, 0}

	for i := 0; i < len(wireSteps); i++ {
		wirePositions[i+1] = CalcPosition(wireSteps[i], wirePositions[i])
	}

	return wirePositions
}

func CalcPosition(wireStep WirePathStep, prevPosition Position) Position {
	x := prevPosition.X
	y := prevPosition.Y

	switch wireStep.Directio {
	case right:
		x += wireStep.Distance
	case left:
		x -= wireStep.Distance
	case up:
		y += wireStep.Distance
	case down:
		y -= wireStep.Distance
	}

	return Position{x, y}
}

func ParseWirePaths(wire string) []WirePathStep {
	split := strings.Split(wire, ",")
	deserialized := make([]WirePathStep, len(split))
	for i, e := range split {
		direction := e[0:1]
		distance, _ := strconv.ParseInt(e[1:], 10, 64)
		deserialized[i] = WirePathStep{direction, distance}
	}

	return deserialized
}

func ParseWirePathsInput(wires string) ([]WirePathStep, []WirePathStep) {
	lineSplit := strings.Split(wires, "\n")
	wire1 := ParseWirePaths(lineSplit[0])
	wire2 := ParseWirePaths(lineSplit[1])

	return wire1, wire2
}

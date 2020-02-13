package day3

import (
	"advent-of-code-2019/shared"
	"math"
	"strconv"
	"strings"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type WirePathStep struct {
	direction string
	distance  int64
}

type Position struct {
	x int64
	y int64
}

func Problem1Answer() int64 {
	input := shared.GetInput("day3/input")
	wire1, wire2 := ParseWirePathsInput(input)

	distance := DistanceToClosestWireIntersectFromCentralPort(wire1, wire2)

	return distance
}

func FewestStepsToWireIntersectFromCentralPort(wire1 []WirePathStep, wire2 []WirePathStep) int64 {
	panic("not implemented")
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
	if l1p1.x == l1p2.x && l2p1.x == l2p2.x {

	} else if l1p1.y == l1p2.y && l2p1.y == l2p2.y {

	} else if l1p1.x == l1p2.x && l2p1.y == l2p2.y {
		if Min(l1p1.y, l1p2.y) <= l2p1.y && l2p1.y <= Max(l1p1.y, l1p2.y) && Min(l2p1.x, l2p2.x) <= l1p1.x && l1p1.x <= Max(l2p1.x, l2p2.x) {
			return CalcDistance(Position{0, 0}, Position{l1p1.x, l2p1.y})
		}
	} else if l1p1.y == l1p2.y && l2p1.x == l2p2.x {
		if Min(l1p1.x, l1p2.x) <= l2p1.x && l2p1.x <= Max(l1p1.x, l1p2.x) && Min(l2p1.y, l2p2.y) <= l1p1.y && l1p1.y <= Max(l2p1.y, l2p2.y) {
			return CalcDistance(Position{0, 0}, Position{l2p1.x, l1p1.y})
		}
	}

	return math.MaxInt64
}

func CalcDistance(p1, p2 Position) int64 {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
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
	x := prevPosition.x
	y := prevPosition.y

	switch wireStep.direction {
	case right:
		x += wireStep.distance
	case left:
		x -= wireStep.distance
	case up:
		y += wireStep.distance
	case down:
		y -= wireStep.distance
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

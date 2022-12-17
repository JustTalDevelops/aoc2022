package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// targetRow represents the target row for problem one.
	targetRow = 2_000_000
	// multiplyFactor represents the factor to multiply with for problem two.
	multiplyFactor = 4_000_000
	// positionCap represents the maximum value in a position axis for problem two.
	positionCap = 4_000_000
)

// main solves problems one and two for Advent of Code (Day Fifteen).
func main() {
	entryRegex := regexp.MustCompile("(-?\\d+)")

	var (
		sensors []sensor
		beacons []beacon
	)
	for _, entry := range strings.Split(strings.ReplaceAll(string(lo.Must(os.ReadFile("input.txt"))), "\r", ""), "\n") {
		coordinates := lo.Map(entryRegex.FindAllString(entry, 4), func(s string, _ int) int {
			return lo.Must(strconv.Atoi(s))
		})

		sensorPos := position{coordinates[0], coordinates[1]}
		beaconPos := position{coordinates[2], coordinates[3]}

		sensors = append(sensors, sensor{
			pos:    sensorPos,
			radius: sensorPos.distance(beaconPos),
		})
		beacons = append(beacons, beacon{pos: beaconPos})
	}

	var (
		frequency int
		index     = make(map[int]struct{})
	)
	for _, s := range sensors {
		if frequency < 1 {
			topY, bottomY := s.pos[1]-s.radius-1, s.pos[1]+s.radius+1
			for y := topY; y <= bottomY; y++ {
				centerDist := abs(s.pos[1] - y)

				left := position{s.pos[0] - s.radius - 1 + centerDist, y}
				right := position{s.pos[0] + s.radius + 1 - centerDist, y}
				if left[0] < 0 || left[0] > positionCap || right[0] < 0 || right[0] > positionCap {
					// Either the left or right position is out of bounds.
					continue
				}
				if y < 0 || y > positionCap {
					// The current row is out of bounds.
					continue
				}
				if !lo.ContainsBy(sensors, func(s sensor) bool {
					return s.radius >= s.pos.distance(left)
				}) {
					frequency = (left[0] * multiplyFactor) + left[1]
					break
				}
				if !lo.ContainsBy(sensors, func(s sensor) bool {
					return s.radius >= s.pos.distance(right)
				}) {
					frequency = (right[0] * multiplyFactor) + right[1]
					break
				}
			}
		}

		centerDist := abs(s.pos[1] - targetRow)
		start, end := s.pos[0]-s.radius+centerDist, s.pos[0]+s.radius-centerDist
		for x := start; x <= end; x++ {
			target := position{x, targetRow}
			if !lo.ContainsBy(beacons, func(b beacon) bool {
				return b.pos == target
			}) {
				index[x] = struct{}{}
			}
		}
	}

	fmt.Printf("Part One: %d\n", len(index))
	fmt.Printf("Part Two: %d\n", frequency)
}

// position represents a position in a 2D space.
type position [2]int

// distance returns the Manhattan distance between two positions.
func (p position) distance(other position) int {
	return abs(p[0]-other[0]) + abs(p[1]-other[1])
}

// sensor represents a sensor as described by AOC.
type sensor struct {
	pos    position
	radius int
}

// beacon represents a beacon as described by AOC.
type beacon struct {
	pos position
}

// abs returns the absolute value of the given integer.
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

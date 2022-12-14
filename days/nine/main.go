package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Nine).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		knotsOne = make([]pos, 2)
		knotsTwo = make([]pos, 10)

		visitedOne = make(map[pos]struct{})
		visitedTwo = make(map[pos]struct{})
	)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction, steps := line[0], lo.Must(strconv.Atoi(line[1]))
		for i := 0; i < steps; i++ {
			for _, knots := range [][]pos{knotsOne, knotsTwo} {
				target := &knots[0]
				switch direction {
				case "U":
					target[1]++
				case "D":
					target[1]--
				case "L":
					target[0]--
				case "R":
					target[0]++
				}

				for k := 1; k < len(knots); k++ {
					this, before := knots[k], knots[k-1]
					if diff := before.sub(this); abs(diff[0]) > 1 || abs(diff[1]) > 1 {
						knots[k] = this.add(diff.delta())
					}
				}
			}

			visitedOne[knotsOne[len(knotsOne)-1]] = struct{}{}
			visitedTwo[knotsTwo[len(knotsTwo)-1]] = struct{}{}
		}
	}

	fmt.Printf("Part One: %d\n", len(visitedOne))
	fmt.Printf("Part Two: %d\n", len(visitedTwo))

	lo.Must0(scanner.Err())
}

// pos represents a position in a 2D space.
type pos [2]int

// add adds a given position to the current position.
func (p pos) add(other pos) pos {
	return pos{p[0] + other[0], p[1] + other[1]}
}

// sub subtracts a given position from the current position.
func (p pos) sub(other pos) pos {
	return pos{p[0] - other[0], p[1] - other[1]}
}

// delta returns the delta of the given position.
func (p pos) delta() pos {
	return pos{delta(p[0]), delta(p[1])}
}

// abs returns the absolute value of the given integer.
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// delta returns the delta of the given integer.
func delta(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

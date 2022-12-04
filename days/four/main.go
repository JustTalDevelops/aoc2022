package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Four).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	var (
		fullyContainingRanges int
		overlappingRanges     int
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ranges := lo.Map(strings.Split(scanner.Text(), ","), func(s string, _ int) []int {
			return lo.Map(strings.Split(s, "-"), func(s string, _ int) int {
				return lo.Must(strconv.Atoi(s))
			})
		})

		rangeOne, rangeTwo := ranges[0], ranges[1]
		if (rangeOne[0] <= rangeTwo[0] && rangeOne[1] >= rangeTwo[1]) ||
			(rangeTwo[0] <= rangeOne[0] && rangeTwo[1] >= rangeOne[1]) {

			fullyContainingRanges++
		}
		if (rangeOne[0] <= rangeTwo[0] && rangeOne[1] >= rangeTwo[0]) ||
			(rangeTwo[0] <= rangeOne[0] && rangeTwo[1] >= rangeOne[0]) {

			overlappingRanges++
		}
	}

	lo.Must0(scanner.Err())

	fmt.Printf("Part One: %d\n", fullyContainingRanges)
	fmt.Printf("Part Two: %d\n", overlappingRanges)
}

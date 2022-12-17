package main

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"os"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Fourteen).
func main() {
	grid, bottom := make(map[[2]int]struct{}), 0
	for _, path := range strings.Split(strings.ReplaceAll(string(lo.Must(os.ReadFile("input.txt"))), "\r", ""), "\n") {
		borders := strings.Split(path, " -> ")
		for i := 0; i < len(borders)-1; i++ {
			f := lo.Map(strings.Split(borders[i], ","), func(s string, _ int) int {
				return lo.Must(strconv.Atoi(s))
			})
			t := lo.Map(strings.Split(borders[i+1], ","), func(s string, _ int) int {
				return lo.Must(strconv.Atoi(s))
			})

			bottom = lo.Max([]int{bottom, t[1]})

			fromX, fromY := lo.Min([]int{f[0], t[0]}), lo.Min([]int{f[1], t[1]})
			toX, toY := lo.Max([]int{f[0], t[0]}), lo.Max([]int{f[1], t[1]})
			for x := fromX; x <= toX; x++ {
				for y := fromY; y <= toY; y++ {
					grid[[2]int{x, y}] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("Part One: %d\n", simulate(maps.Clone(grid), bottom, true))
	fmt.Printf("Part Two: %d\n", simulate(maps.Clone(grid), bottom, false))
}

// simulate performs a sand-grain falling simulation on the grid given and returns the result.
func simulate(grid map[[2]int]struct{}, bottom int, void bool) int {
	formations := len(grid)
	queue := [][2]int{{500, 0}}
	for len(queue) > 0 {
		entry := queue[len(queue)-1]
		if void && entry[1] >= bottom {
			// We're in the void, so we can stop here.
			break
		}
		if entry[1] > bottom {
			grid[entry] = struct{}{}
			queue = queue[:len(queue)-1]
			continue
		}

		moves := lo.Filter([][2]int{
			{entry[0] + 1, entry[1] + 1},
			{entry[0] - 1, entry[1] + 1},
			{entry[0], entry[1] + 1},
		}, func(next [2]int, _ int) bool {
			_, ok := grid[next]
			return !ok
		})
		if len(moves) == 0 {
			grid[entry] = struct{}{}
			queue = queue[:len(queue)-1]
			continue
		}
		queue = append(queue, moves...)
	}
	return len(grid) - formations
}

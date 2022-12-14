package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Twelve).
func main() {
	input := string(lo.Must(os.ReadFile("input.txt")))
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")

	var (
		start, end [2]int
		grid       [][]int
	)
	for y, row := range lo.Map(lines, func(s string, _ int) []rune {
		return []rune(s)
	}) {
		grid = append(grid, make([]int, len(row)))
		for x, cell := range row {
			if cell == 'S' {
				start = [2]int{x, y}
				grid[y][x] = lo.IndexOf(lo.LowerCaseLettersCharset, 'a')
			} else if cell == 'E' {
				end = [2]int{x, y}
				grid[y][x] = lo.IndexOf(lo.LowerCaseLettersCharset, 'z')
			} else {
				grid[y][x] = lo.IndexOf(lo.LowerCaseLettersCharset, cell)
			}
		}
	}

	fmt.Printf("Part One: %d\n", stepsToPath(grid, start, func(pos [2]int, val int) bool {
		return pos == end
	}, func(val, nVal int) bool {
		return nVal <= val+1
	}))
	fmt.Printf("Part Two: %d\n", stepsToPath(grid, end, func(pos [2]int, val int) bool {
		return val == 0
	}, func(val, nVal int) bool {
		return nVal >= val-1
	}))
}

// entry represents an entry in the queue.
type entry struct {
	pos   [2]int
	steps int
}

// stepsToPath returns the amount of steps it takes to get from the start to the target on the grid passed.
func stepsToPath(grid [][]int, start [2]int, target func(pos [2]int, val int) bool, neighbour func(val, nVal int) bool) int {
	visited := make(map[[2]int]struct{})
	queue := []entry{{pos: start}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if _, ok := visited[cur.pos]; ok {
			// We've already visited this position, so we can skip it.
			continue
		}

		visited[cur.pos] = struct{}{}

		val := grid[cur.pos[1]][cur.pos[0]]
		if target(cur.pos, val) {
			// We've found the target, so we can return the amount of steps it took to get here.
			return cur.steps
		}

		for _, n := range neighbours(grid, cur.pos) {
			nVal := grid[n[1]][n[0]]
			if _, ok := visited[n]; !ok && neighbour(val, nVal) {
				queue = append(queue, entry{
					pos:   n,
					steps: cur.steps + 1,
				})
			}
		}
	}
	return -1
}

// neighbours returns all neighbours of the given position on the grid.
func neighbours(grid [][]int, pos [2]int) (positions [][2]int) {
	return lo.Filter(lo.Map([][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}, func(offset [2]int, index int) [2]int {
		return [2]int{pos[0] + offset[0], pos[1] + offset[1]}
	}), func(pos [2]int, _ int) bool {
		return pos[0] >= 0 && pos[0] < len(grid[0]) && pos[1] >= 0 && pos[1] < len(grid)
	})
}

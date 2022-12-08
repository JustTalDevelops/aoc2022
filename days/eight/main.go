package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Eight).
func main() {
	input := strings.ReplaceAll(string(lo.Must(os.ReadFile("input.txt"))), "\r", "")
	grid := lo.Map(strings.Split(input, "\n"), func(l string, _ int) []int {
		return lo.Map([]rune(l), func(r rune, _ int) int {
			return int(r)
		})
	})

	var (
		visibleTrees    int
		bestScenicScore int
	)
	for row := 0; row < len(grid); row++ {
		rows, columns := len(grid), len(grid[row])
		for column := 0; column < len(grid[row]); column++ {
			if column == 0 || column == columns-1 || row == 0 || row == rows-1 {
				visibleTrees++
				continue
			}

			val := grid[row][column]

			// Above.
			aboveVisible, aboveValue := true, 0
			for i := row - 1; i >= 0; i-- {
				aboveValue++
				if grid[i][column] >= val {
					aboveVisible = false
					break
				}
			}

			// Below.
			belowVisible, belowValue := true, 0
			for i := row + 1; i < rows; i++ {
				belowValue++
				if grid[i][column] >= val {
					belowVisible = false
					break
				}
			}

			// Left.
			leftVisible, leftValue := true, 0
			for i := column - 1; i >= 0; i-- {
				leftValue++
				if grid[row][i] >= val {
					leftVisible = false
					break
				}
			}

			// Right.
			rightVisible, rightValue := true, 0
			for i := column + 1; i < columns; i++ {
				rightValue++
				if grid[row][i] >= val {
					rightVisible = false
					break
				}
			}

			bestScenicScore = lo.Max([]int{bestScenicScore, aboveValue * belowValue * leftValue * rightValue})
			if aboveVisible || belowVisible || leftVisible || rightVisible {
				visibleTrees++
			}
		}
	}

	fmt.Printf("Part One: %d\n", visibleTrees)
	fmt.Printf("Part Two: %d\n", bestScenicScore)
}

package main

import (
	"fmt"
	"github.com/samber/lo"
	"os"
)

// main solves problems one and two for Advent of Code (Day Six).
func main() {
	input := []rune(string(lo.Must(os.ReadFile("input.txt"))))

	fmt.Printf("Part One: %d\n", indexOfUniqueSelection(input, 4))
	fmt.Printf("Part Two: %d\n", indexOfUniqueSelection(input, 14))
}

// indexOfUniqueSelection returns the index of the first unique selection of runes based on the given length.
func indexOfUniqueSelection[T comparable](collection []T, length int) int {
	for i := 0; i < len(collection); i++ {
		if len(lo.FindDuplicates(collection[i:i+length])) < 1 {
			return i + length
		}
	}
	panic("should never happen`")
}

package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Three).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		prioritySumOne int
		prioritySumTwo int

		group []string
	)
	for scanner.Scan() {
		sack := scanner.Text()
		group = append(group, sack)

		center := len(sack) / 2
		firstCompartment, secondCompartment := sack[:center], sack[center:]

		for i, r := range lo.LettersCharset {
			if strings.ContainsRune(firstCompartment, r) && strings.ContainsRune(secondCompartment, r) {
				prioritySumOne += i + 1
			}

			if len(group) == 3 && len(group) == lo.CountBy(group, func(s string) bool {
				return strings.ContainsRune(s, r)
			}) {
				prioritySumTwo += i + 1
				group = nil
			}
		}
	}

	lo.Must0(scanner.Err())

	fmt.Printf("Part One: %d\n", prioritySumOne)
	fmt.Printf("Part Two: %d\n", prioritySumTwo)
}

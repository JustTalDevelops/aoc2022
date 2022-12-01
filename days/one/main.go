package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"sort"
	"strconv"
)

// main solves problems one and two for Advent of Code (Day One).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		elves []int
		elf   []int
	)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			elves, elf = append(elves, lo.Sum(elf)), nil
			continue
		}
		elf = append(elf, lo.Must(strconv.Atoi(line)))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("Part One: %d\n", elves[0])
	fmt.Printf("Part Two: %d\n", lo.Sum(elves[:3]))
}

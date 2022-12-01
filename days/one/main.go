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
		sum  int
		sums []int
	)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			sums, sum = append(sums, sum), 0
			continue
		}
		sum += lo.Must(strconv.Atoi(line))
	}

	lo.Must0(scanner.Err())

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Printf("Part One: %d\n", sums[0])
	fmt.Printf("Part Two: %d\n", lo.Sum(sums[:3]))
}

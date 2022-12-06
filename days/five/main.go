package main

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Five).
func main() {
	letterRegex := regexp.MustCompile(`\[([A-Z])\]`)
	moveRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	input := lo.Must(os.ReadFile("input.txt"))

	sections := strings.Split(strings.ReplaceAll(string(input), "\r", ""), "\n\n")
	graphSection, moveSection := strings.Split(sections[0], "\n"), strings.Split(sections[1], "\n")

	initialGraph := make(map[int][]string)
	for _, line := range graphSection[:len(graphSection)-1] {
		for _, entry := range strings.Fields(line) {
			letter := letterRegex.FindStringSubmatch(entry)[1]

			ind := strings.Index(line, letter)
			line = strings.Replace(line, letter, " ", 1)

			column := (ind + 3) / 4
			if _, ok := initialGraph[column]; !ok {
				initialGraph[column] = make([]string, 0)
			}
			initialGraph[column] = append(initialGraph[column], letter)
		}
	}

	partOneGraph := maps.Clone(initialGraph)
	partTwoGraph := maps.Clone(initialGraph)
	for _, rawMove := range moveSection {
		move := lo.Map(moveRegex.FindStringSubmatch(rawMove)[1:], func(s string, _ int) int {
			return lo.Must(strconv.Atoi(s))
		})

		amount, from, to := move[0], move[1], move[2]
		for i := 0; i < amount; i++ {
			crate := partOneGraph[from][0]
			partOneGraph[from] = partOneGraph[from][1:]
			partOneGraph[to] = append([]string{crate}, partOneGraph[to]...)
		}

		crates := slices.Clone(partTwoGraph[from][:amount])
		partTwoGraph[from] = partTwoGraph[from][amount:]
		partTwoGraph[to] = append(crates, partTwoGraph[to]...)
	}

	indexes := lo.Keys(initialGraph)
	sort.Ints(indexes)

	partOne, partTwo := &strings.Builder{}, &strings.Builder{}
	for _, index := range indexes {
		partOne.WriteString(partOneGraph[index][0])
		partTwo.WriteString(partTwoGraph[index][0])
	}

	fmt.Printf("Part One: %s\n", partOne.String())
	fmt.Printf("Part Two: %s\n", partTwo.String())
}

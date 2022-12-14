package main

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// monkey represents a brainless animal throwing items around that I have to deal with.
type monkey struct {
	items []int

	operation rune
	operand   string

	divisor int

	monkeyIfTrue  int
	monkeyIfFalse int

	inspections int
}

// main solves problems one and two for Advent of Code (Day Eleven).
func main() {
	monkeyRegex := regexp.MustCompile(`Monkey (\d+):\n` +
		`\s+Starting items: ((\d+(, )?)+)\n` +
		`\s+Operation: new = old ([+-/*]) (\d+|old)\n` +
		`\s+Test: divisible by (\d+)\n` +
		`\s+If true: throw to monkey (\d+)\n` +
		`\s+If false: throw to monkey (\d+)`)

	input := lo.Must(os.ReadFile("input.txt"))
	sections := strings.Split(strings.ReplaceAll(string(input), "\r", ""), "\n\n")

	monkeys := make([]monkey, len(sections))
	for _, section := range sections {
		matches := monkeyRegex.FindStringSubmatch(section)
		monkeys[lo.Must(strconv.Atoi(matches[1]))] = monkey{
			operation:     []rune(matches[5])[0],
			operand:       matches[6],
			divisor:       lo.Must(strconv.Atoi(matches[7])),
			monkeyIfTrue:  lo.Must(strconv.Atoi(matches[8])),
			monkeyIfFalse: lo.Must(strconv.Atoi(matches[9])),
			items: lo.Map(strings.Split(matches[2], ", "), func(s string, index int) int {
				return lo.Must(strconv.Atoi(s))
			}),
		}
	}

	fmt.Printf("Part One: %d\n", monkeyBusiness(slices.Clone(monkeys), 20, true))
	fmt.Printf("Part Two: %d\n", monkeyBusiness(slices.Clone(monkeys), 10_000, false))
}

// monkeyBusiness calculates the amount of monkey related shenanigans that occur within a given number of rounds.
func monkeyBusiness(monkeys []monkey, totalRounds int, ableToCope bool) int {
	divisors := lo.Map(monkeys, func(m monkey, _ int) int {
		return m.divisor
	})
	lcm := lo.Reduce(divisors, func(a, b int, _ int) int {
		return a * b
	}, 1)

	var (
		worryLevel     int
		monkeyToSendTo int
	)
	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			m := &monkeys[i]
			m.inspections += len(m.items)

			for _, item := range m.items {
				operand, err := strconv.Atoi(m.operand)
				if err != nil {
					operand = item
				}

				switch m.operation {
				case '+':
					worryLevel = item + operand
				case '-':
					worryLevel = item - operand
				case '*':
					worryLevel = item * operand
				case '/':
					worryLevel = item / operand
				}
				if ableToCope {
					worryLevel /= 3
				} else {
					worryLevel %= lcm
				}

				if worryLevel%m.divisor == 0 {
					monkeyToSendTo = m.monkeyIfTrue
				} else {
					monkeyToSendTo = m.monkeyIfFalse
				}

				target := &monkeys[monkeyToSendTo]
				target.items = append(target.items, worryLevel)
			}

			// We've thrown all the items away, so we can clear the slice.
			m.items = nil
		}
	}

	business := lo.Map(monkeys, func(m monkey, _ int) int { return m.inspections })
	sort.Ints(business)

	return business[len(business)-1] * business[len(business)-2]
}

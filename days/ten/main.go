package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Ten).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	display, x := make([][]rune, 6), 1
	for i := 0; i < len(display); i++ {
		display[i] = make([]rune, 40)
	}

	var (
		cycle int
		sum   int
	)
	for scanner.Scan() {
		// Every operation is guaranteed to have at least a single tick.
		tick(display, &cycle, &sum, &x)

		operation := scanner.Text()
		if operation == "noop" {
			// No operation.
			continue
		}

		// We need to simulate an extra tick for the adding operation.
		tick(display, &cycle, &sum, &x)

		x += lo.Must(strconv.Atoi(strings.TrimPrefix(operation, "addx ")))
	}

	fmt.Printf("Part One: %d\n", sum)
	fmt.Printf("Part Two:\n")
	for _, row := range display {
		fmt.Println(string(row))
	}

	lo.Must0(scanner.Err())
}

// tick performs a single tick on the CPU.
func tick(display [][]rune, cycle, sum, x *int) {
	row := *cycle / 40
	column := *cycle % 40
	if column >= *x-1 && column <= *x+1 {
		display[row][column] = '█'
	} else {
		display[row][column] = '░'
	}

	if *cycle++; *cycle%40 == 20 {
		*sum += *cycle * *x
	}
}

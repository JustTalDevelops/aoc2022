package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"path"
	"strconv"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Seven).
func main() {
	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		working     string
		directories = make(map[string]int)
	)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd .." {
			working = path.Dir(working)
			continue
		}
		if dir := strings.TrimPrefix(line, "$ cd "); line != dir {
			working = path.Join(working, dir)
			continue
		}
		size, err := strconv.Atoi(strings.Split(line, " ")[0])
		if err != nil {
			continue
		}

		directories["/"] += size
		for dir := working; len(dir) > 1; dir = path.Dir(dir) {
			directories[dir] += size
		}
	}

	fmt.Printf("Part One: %d\n", lo.Sum(lo.Filter(lo.Values(directories), func(size, _ int) bool {
		return size <= 100000
	})))

	const (
		driveSize   = 70000000
		minimumSize = 30000000
	)

	target := minimumSize - (driveSize - directories["/"])
	bestSize := driveSize
	for _, size := range directories {
		if size > target && size < bestSize {
			bestSize = size
		}
	}

	fmt.Printf("Part Two: %d\n", bestSize)

	lo.Must0(scanner.Err())
}

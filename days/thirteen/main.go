package main

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"os"
	"sort"
	"strings"
)

// main solves problems one and two for Advent of Code (Day Thirteen).
func main() {
	input := lo.Must(os.ReadFile("input.txt"))

	batches := lo.Map(strings.Split(strings.ReplaceAll(string(input), "\r", ""), "\n\n"), func(s string, _ int) [2]any {
		unparsedPackets := strings.Split(s, "\n")

		var firstPacket, secondPacket any
		lo.Must0(json.Unmarshal([]byte(unparsedPackets[0]), &firstPacket))
		lo.Must0(json.Unmarshal([]byte(unparsedPackets[1]), &secondPacket))

		return [2]any{firstPacket, secondPacket}
	})

	fmt.Printf("Part One: %d\n", lo.Sum(lo.FilterMap(batches, func(batch [2]any, i int) (int, bool) {
		if packetComparator(batch[0], batch[1]) > 0 {
			return 0, false
		}
		return i + 1, true
	})))

	separatorPackets := []any{[]any{[]any{2.0}}, []any{[]any{6.0}}}
	encodedSeparatorPackets := lo.Map(separatorPackets, func(packet any, _ int) string {
		return string(lo.Must(json.Marshal(packet)))
	})

	packets := append(lo.FlatMap(batches, func(batch [2]any, _ int) []any {
		return []any{batch[0], batch[1]}
	}), separatorPackets...)
	sort.SliceStable(packets, func(i, j int) bool {
		return packetComparator(packets[i], packets[j]) < 0
	})

	fmt.Printf("Part Two: %d\n", lo.Reduce(packets, func(product int, packet any, i int) int {
		return product * lo.Ternary(
			lo.Contains(encodedSeparatorPackets, string(lo.Must(json.Marshal(packet)))),
			i+1, 1,
		)
	}, 1))
}

// packetComparator compares two packets directly, either as numbers or as slices. It returns 0 if they are equal,
// 1 if the first packet is greater than the second, and -1 if the second packet is greater than the first.
func packetComparator(firstPacket, secondPacket any) int {
	firstAsNumber, firstIsNumber := firstPacket.(float64)
	secondAsNumber, secondIsNumber := secondPacket.(float64)
	if firstIsNumber && secondIsNumber {
		return int(firstAsNumber - secondAsNumber)
	}

	firstSection, firstIsSection := firstPacket.([]any)
	secondSection, secondIsSection := secondPacket.([]any)
	if !firstIsSection {
		firstSection = []any{firstPacket}
	} else if !secondIsSection {
		secondSection = []any{secondPacket}
	}

	for i := 0; i < lo.Min([]int{len(firstSection), len(secondSection)}); i++ {
		if v := packetComparator(firstSection[i], secondSection[i]); v != 0 {
			return v
		}
	}
	return len(firstSection) - len(secondSection)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
)

const (
	resultLose = "X"
	resultDraw = "Y"
	resultWin  = "Z"

	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"

	rockBonus     = 1
	paperBonus    = 2
	scissorsBonus = 3

	winBonus = 6
	tieBonus = 3
)

// main solves problems one and two for Advent of Code (Day Two).
func main() {
	var (
		scoreCombinations    = make(map[string]int)
		scoreCombinationsTwo = make(map[string]int)
	)
	for _, opponentMove := range []string{opponentRock, opponentPaper, opponentScissors} {
		for _, unknown := range []string{resultLose, resultDraw, resultWin} {
			var firstScore, secondScore int
			switch unknown {
			case resultLose:
				firstScore += rockBonus
				switch opponentMove {
				case opponentRock:
					firstScore += tieBonus
					secondScore += scissorsBonus
				case opponentScissors:
					firstScore += winBonus
					secondScore += paperBonus
				case opponentPaper:
					secondScore += rockBonus
				}
			case resultDraw:
				firstScore += paperBonus
				switch opponentMove {
				case opponentRock:
					firstScore += winBonus
					secondScore += rockBonus + tieBonus
				case opponentPaper:
					firstScore += tieBonus
					secondScore += paperBonus + tieBonus
				case opponentScissors:
					secondScore += scissorsBonus + tieBonus
				}
			case resultWin:
				firstScore += scissorsBonus
				switch opponentMove {
				case opponentPaper:
					firstScore += winBonus
					secondScore += scissorsBonus + winBonus
				case opponentScissors:
					firstScore += tieBonus
					secondScore += rockBonus + winBonus
				case opponentRock:
					secondScore += paperBonus + winBonus
				}
			}

			code := fmt.Sprintf("%s %s", opponentMove, unknown)
			scoreCombinations[code] = firstScore
			scoreCombinationsTwo[code] = secondScore
		}
	}

	f := lo.Must(os.Open("input.txt"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		totalScoreFirst  int
		totalScoreSecond int
	)
	for scanner.Scan() {
		combination := scanner.Text()
		totalScoreFirst += scoreCombinations[combination]
		totalScoreSecond += scoreCombinationsTwo[combination]
	}

	fmt.Printf("Part One: %d\n", totalScoreFirst)
	fmt.Printf("Part Two: %d\n", totalScoreSecond)
}

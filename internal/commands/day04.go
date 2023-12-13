package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day04Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; fewest number of cubes of each color"`
}

func (c Day04Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day04Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day04part2(input)
	}

	return day04part1(input)
}

func day04part1(lines []string) (sum int) {
	for _, line := range lines {
		winningNums, haveNums := parseCard(line)
		matches := countMatches(winningNums, haveNums)
		if matches == 0 {
			continue
		}
		sum += 1 << (matches - 1)
	}

	return sum
}

func day04part2(lines []string) (sum int) {
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}

	for i, line := range lines {
		sum += copies[i]

		winningNums, haveNums := parseCard(line)
		matches := countMatches(winningNums, haveNums)
		for m := 1; m <= matches; m++ {
			copies[i+m] += copies[i]
		}
	}

	return sum + 1
}

func parseCard(card string) (winningNums map[int]struct{}, haveNums []int) {
	values := strings.Split(card, ": ")[1]
	split := strings.Split(values, " | ")
	winningStrs, haveStrs := split[0], split[1]

	winningNums = make(map[int]struct{})
	for _, winning := range strings.Split(winningStrs, " ") {
		if winning == "" {
			continue
		}
		// assumes input is fine
		num, _ := strconv.Atoi(winning)
		winningNums[num] = struct{}{}
	}

	for _, have := range strings.Split(haveStrs, " ") {
		if have == "" {
			continue
		}
		// assumes input is fine
		num, _ := strconv.Atoi(have)
		haveNums = append(haveNums, num)
	}

	return winningNums, haveNums
}

func countMatches(winning map[int]struct{}, have []int) (score int) {
	for _, h := range have {
		if _, ok := winning[h]; ok {
			score++
		}
	}
	return score
}

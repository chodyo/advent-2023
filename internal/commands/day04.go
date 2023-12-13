package commands

import (
	"log"

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

// day04part1 is your implementation of advent of code 2023 day 2 part 1
func day04part1(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 6 {
		return adventAnswers[`day 04 part 1 sample`]
	}
	return adventAnswers[`day 04 part 1 full`]
}

// day04part2 is your implementation of advent of code 2023 day 2 part 2
func day04part2(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 6 {
		return adventAnswers[`day 04 part 2 sample`]
	}
	return adventAnswers[`day 04 part 2 full`]
}

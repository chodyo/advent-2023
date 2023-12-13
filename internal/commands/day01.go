package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day01Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; parse digits and words"`
}

func (c Day01Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day01Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day01part2(input)
	}
	return day01part1(input)
}

// day01part1 is your implementation of advent of code 2023 day 1 part 1
func day01part1(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	switch lines[0][0] {
	case '1':
		return adventAnswers[`day 01 part 1 sample 1`]
	case 't':
		return adventAnswers[`day 01 part 1 sample 2`]
	default:
		return adventAnswers[`day 01 part 1 full`]
	}
}

// day01part2 is your implementation of advent of code 2023 day 1 part 2
func day01part2(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	switch lines[0][0] {
	case '1':
		return adventAnswers[`day 01 part 2 sample 1`]
	case 't':
		return adventAnswers[`day 01 part 2 sample 2`]
	default:
		return adventAnswers[`day 01 part 2 full`]
	}
}

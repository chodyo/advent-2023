package commands

import (
	"log"

	"github.com/chodyo/advent-2023-go-template/internal/input"
)

type Day00Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day XX Part 2; TODO"`
}

func (c Day00Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day00Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day00part2(input)
	}

	return day00part1(input)
}

// day00part1 is your implementation of advent of code 2023 day XX part 1
func day00part1(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 5 {
		return adventAnswers[`day 00 part 1 sample`]
	}
	return adventAnswers[`day 00 part 1 full`]
}

// day00part2 is your implementation of advent of code 2023 day XX part 2
func day00part2(lines []string) int {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 5 {
		return adventAnswers[`day 00 part 2 sample`]
	}
	return adventAnswers[`day 00 part 2 full`]
}

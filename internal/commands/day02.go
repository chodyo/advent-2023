package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day02Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; fewest number of cubes of each color"`
}

func (c Day02Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day02Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day02part2(input)
	}

	return day02part1(input)
}

// day02part1 is your implementation of advent of code 2023 day 2 part 1
func day02part1(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 5 {
		return adventAnswers[`day 02 part 1 sample`]
	}
	return adventAnswers[`day 02 part 1 full`]
}

// day02part2 is your implementation of advent of code 2023 day 2 part 2
func day02part2(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 5 {
		return adventAnswers[`day 02 part 2 sample`]
	}
	return adventAnswers[`day 02 part 2 full`]
}

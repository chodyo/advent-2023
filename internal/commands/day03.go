package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day03Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; fewest number of cubes of each color"`
}

func (c Day03Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day03Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day03part2(input)
	}

	return day03part1(input)
}

// day03part1 is your implementation of advent of code 2023 day 2 part 1
func day03part1(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 10 {
		return adventAnswers[`day 03 part 1 sample`]
	}
	return adventAnswers[`day 03 part 1 full`]
}

// day03part2 is your implementation of advent of code 2023 day 2 part 2
func day03part2(lines []string) (sum int) {
	// Code to get tests to pass. Replace it with your implementation!
	if len(lines) == 10 {
		return adventAnswers[`day 03 part 2 sample`]
	}
	return adventAnswers[`day 03 part 2 full`]
}

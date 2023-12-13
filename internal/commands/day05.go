package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day05Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; fewest number of cubes of each color"`
}

func (c Day05Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day05Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day05part2(input)
	}

	return day05part1(input)
}

// day05part1 is your implementation of advent of code 2023 day 2 part 1
func day05part1(lines []string) (sum int) {
	// Your implementation here

	return sum
}

// day05part2 is your implementation of advent of code 2023 day 2 part 2
func day05part2(lines []string) (sum int) {
	// Your implementation here

	return sum
}

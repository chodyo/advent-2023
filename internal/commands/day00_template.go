package commands

import (
	"log"

	"github.com/chodyo/advent-2023-go-template/internal/input"
)

type Day00Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; fewest number of cubes of each color"`
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

// // day00part1 is your implementation of advent of code 2023 day 2 part 1
// func day00part1(lines []string) int {
// 	var sum int

// 	// Your implementation here

// 	return sum
// }

// // day00part2 is your implementation of advent of code 2023 day 2 part 2
// func day00part2(lines []string) int {
// 	var sum int

// 	// Your implementation here

// 	return sum
// }

func day00part1(lines []string) int {
	var sum int
	return sum
}

func day00part2(lines []string) int {
	var sum int
	return sum
}

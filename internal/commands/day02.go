package commands

import (
	"log"
	"strconv"
	"strings"

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

func day02part1(lines []string) int {
	var sum int

	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for i, game := range lines {
		pulls := strings.Split(game, ": ")[1]
		valid := true
	pulls:
		for _, pull := range strings.Split(pulls, "; ") {
			for _, cube := range strings.Split(pull, ", ") {
				c := strings.Split(cube, " ")
				amount, _ := strconv.Atoi(c[0])
				color := c[1]
				if max[color] < amount {
					valid = false
					break pulls
				}
			}
		}
		if valid {
			sum += i + 1
		}
	}

	return sum
}

func day02part2(lines []string) int {
	var sum int

	for _, game := range lines {
		pulls := strings.Split(game, ": ")[1]

		max := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, pull := range strings.Split(pulls, "; ") {
			for _, cube := range strings.Split(pull, ", ") {
				c := strings.Split(cube, " ")
				amount, _ := strconv.Atoi(c[0])
				color := c[1]
				if max[color] < amount {
					max[color] = amount
				}
			}
		}

		sum += max["red"] * max["green"] * max["blue"]
	}

	return sum
}

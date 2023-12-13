package commands

import (
	"log"
	"strconv"
	"strings"

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

func day01part1(lines []string) int {
	var sum int

	for _, line := range lines {
		leftmost := 0
		for _, c := range line {
			if n, err := strconv.Atoi(string(c)); err == nil {
				leftmost = n
				break
			}
		}

		rightmost := 0
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if n, err := strconv.Atoi(string(c)); err == nil {
				rightmost = n
				break
			}
		}

		sum += leftmost*10 + rightmost
	}

	return sum
}

func day01part2(lines []string) int {
	var sum int

	m := map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	for i := 0; i < len(lines); i++ {
		readLine := lines[i]
		writeLine := lines[i]
		for word, digit := range m {
			for idx := strings.Index(readLine, word); idx >= 0; idx = strings.Index(writeLine, word) {
				writeLine = writeLine[:idx] + digit + writeLine[idx+1:]
			}
		}
		lines[i] = writeLine
	}

	sum = day01part1(lines)

	return sum
}

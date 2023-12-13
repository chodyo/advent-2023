package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type AdventCommand interface {
	// Interface for go-flags
	Execute([]string) error

	// Interface for input.Option
	GetInput() []string

	// Interface for advent-2023
	Process([]string) int
}

type DefaultAdventCommand struct {
	input.Option
}

// To be overridden by AdventCommand implementers
func (c DefaultAdventCommand) Execute(args []string) error {
	log.Printf("result=%d\n", c.Process(args))
	return nil
}

func (c DefaultAdventCommand) GetInput() []string {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return nil
	}
	return input
}

// To be overridden by AdventCommand implementers
func (c DefaultAdventCommand) Process(_ []string) int { return 0 }

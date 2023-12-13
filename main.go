package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/chodyo/advent-2023/internal/commands"
)

type Commands struct {
	commands.Day01Command `command:"day01" description:"Advent 2023 Day 1 - Trebuchet Calibration"`
	commands.Day02Command `command:"day02" description:"Advent 2023 Day 2 - Cube Conundrum"`
	commands.Day03Command `command:"day03" description:"Advent 2023 Day 3 - Gear Ratios"`
	commands.Day04Command `command:"day04" description:"Advent 2023 Day 4 - Scratchcards"`
	commands.Day05Command `command:"day04" description:"Advent 2023 Day 5 - If You Give A Seed A Fertilizer"`
}

// Commands are executed in `ParseArgs`
// https://pkg.go.dev/github.com/jessevdk/go-flags?utm_source=godoc#hdr-Commands
func main() {
	// Disable logs timestamp
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	commands := Commands{}
	parser := flags.NewParser(&commands, flags.Default)

	if _, err := parser.ParseArgs(os.Args[1:]); err != nil {
		if flags.WroteHelp(err) {
			return
		}
		os.Exit(1)
	}
}

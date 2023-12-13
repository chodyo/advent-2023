package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Option struct {
	Filename string `short:"f" long:"file" description:"line delimited command inputs (ignores other inputs)"`
}

func LoadInput(filename string) ([]string, error) {
	if filename == "" {
		return loadInputFromStdIn()
	}
	return LoadInputFromFile(filename)
}

func loadInputFromStdIn() (lines []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("no input was provided")
	}

	return lines, nil
}

func LoadInputFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

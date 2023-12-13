//nolint:goconst // Don't wanna make "file" str a const
package commands

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/chodyo/advent-2023/internal/input"
	"github.com/google/go-cmp/cmp"
)

func Test_Commands_TableTesting(t *testing.T) {
	adventTests := map[string]map[string]map[string]Test{
		// custom test handlers - for when they don't look like the default
		`day 01`: {
			`part 1`: getDay01Part1Tests(),
			`part 2`: getDay01Part2Tests(),
		},
	}

	for _, day := range []int{2, 3, 4} {
		// default-looking test handlers
		adventTests[fmt.Sprintf(`day %02d`, day)] = map[string]map[string]Test{
			`part 1`: getDefaultDayXXPartYTests(day, 1),
			`part 2`: getDefaultDayXXPartYTests(day, 2),
		}
	}

	for day, getParts := range adventTests {
		for part, testCases := range getParts {
			for name, testCase := range testCases {
				tc := testCase
				t.Run(fmt.Sprintf("%s %s %s", day, part, name), func(t *testing.T) {
					t.Parallel()

					actual := tc.command.Process(nil)
					diff := cmp.Diff(tc.expected, actual)
					if diff != "" {
						t.Log(diff)
						t.Fail()
					}
				})
			}
		}
	}
}

type Test struct {
	command  AdventCommand
	expected int
}

func getDay01Part1Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`01-1_input_sample`: adventAnswers[`day 01 part 1 sample 1`],
		`01-2_input_sample`: adventAnswers[`day 01 part 1 sample 2`],
		`01_input_full`:     adventAnswers[`day 01 part 1 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day01Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                false,
			},
			expected: out,
		}
	}
	return tests
}

func getDay01Part2Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`01-1_input_sample`: adventAnswers[`day 01 part 2 sample 1`],
		`01-2_input_sample`: adventAnswers[`day 01 part 2 sample 2`],
		`01_input_full`:     adventAnswers[`day 01 part 2 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day01Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                true,
			},
			expected: out,
		}
	}
	return tests
}

func getDefaultDayXXPartYTests(day, part int) map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		fmt.Sprintf(`%02d_input_sample`, day): adventAnswers[fmt.Sprintf(`day %02d part %d sample`, day, part)],
		fmt.Sprintf(`%02d_input_full`, day):   adventAnswers[fmt.Sprintf(`day %02d part %d full`, day, part)],
	}
	for in, out := range expectations {
		var command AdventCommand
		switch day {
		case 1:
			command = Day01Command{DefaultAdventCommand: makeDefaultAdventCommand(in), Part2: part == 2}
		case 2:
			command = Day02Command{DefaultAdventCommand: makeDefaultAdventCommand(in), Part2: part == 2}
		case 3:
			command = Day03Command{DefaultAdventCommand: makeDefaultAdventCommand(in), Part2: part == 2}
		case 4:
			command = Day04Command{DefaultAdventCommand: makeDefaultAdventCommand(in), Part2: part == 2}
		}
		tests["file "+in] = Test{
			command:  command,
			expected: out,
		}
	}
	return tests
}

func makeDefaultAdventCommand(filename string) DefaultAdventCommand {
	return DefaultAdventCommand{
		Option: input.Option{
			Filename: filepath.Join("input", filename) + ".txt",
		},
	}
}

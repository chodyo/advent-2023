package commands

import (
	"log"
	"math/rand"

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

func day03part1(lines []string) int {
	partNumbers := make(map[int64]*node)

	// big assumption: all rows are the same length

	matrix := makeNodeMatrix(lines)
	hasRowAbove := func(row int) bool {
		return row > 0
	}
	hasRowBelow := func(row int) bool {
		return row < len(matrix[0])
	}
	hasColToLeft := func(col int) bool {
		return col > 0
	}
	hasColToRight := func(col int) bool {
		return col < len(matrix)
	}
	for row, nodes := range matrix {
		for col, node := range nodes {
			if !node.isSymbol {
				continue
			}

			// above, i.e. y-1

			if hasRowAbove(row) && hasColToLeft(col) {
				topLeftNode := matrix[row-1][col-1]
				partNumbers[topLeftNode.id] = topLeftNode
			}

			if hasRowAbove(row) {
				topCenterNode := matrix[row-1][col]
				partNumbers[topCenterNode.id] = topCenterNode
			}

			if hasRowAbove(row) && hasColToRight(col) {
				topRightNode := matrix[row-1][col+1]
				partNumbers[topRightNode.id] = topRightNode
			}

			// sides, i.e. y

			if hasColToLeft(col) {
				leftNode := matrix[row][col-1]
				partNumbers[leftNode.id] = leftNode
			}

			if hasColToRight(col) {
				leftNode := matrix[row][col+1]
				partNumbers[leftNode.id] = leftNode
			}

			// below, i.e. y+1

			if hasRowBelow(row) && hasColToLeft(col) {
				bottomLeftNode := matrix[row+1][col-1]
				partNumbers[bottomLeftNode.id] = bottomLeftNode
			}

			if hasRowBelow(row) {
				bottomCenterNode := matrix[row+1][col]
				partNumbers[bottomCenterNode.id] = bottomCenterNode
			}

			if hasRowBelow(row) && hasColToRight(col) {
				bottomRightNode := matrix[row+1][col+1]
				partNumbers[bottomRightNode.id] = bottomRightNode
			}
		}
	}

	var sum int

	for _, n := range partNumbers {
		sum += n.value
	}

	return sum
}

func day03part2(lines []string) int {
	var sum int

	// big assumption: all rows are the same length

	matrix := makeNodeMatrix(lines)
	hasRowAbove := func(row int) bool {
		return row > 0
	}
	hasRowBelow := func(row int) bool {
		return row < len(matrix[0])
	}
	hasColToLeft := func(col int) bool {
		return col > 0
	}
	hasColToRight := func(col int) bool {
		return col < len(matrix)
	}
	for row, nodes := range matrix {
		for col, n := range nodes {
			if !n.isAsterisk {
				continue
			}

			partNumbers := make(map[int64]*node)

			addToPartNumbers := func(n *node) {
				if n.value <= 0 {
					return
				}
				partNumbers[n.id] = n
			}

			// above, i.e. y-1

			if hasRowAbove(row) && hasColToLeft(col) {
				topLeftNode := matrix[row-1][col-1]
				addToPartNumbers(topLeftNode)
			}

			if hasRowAbove(row) {
				topCenterNode := matrix[row-1][col]
				addToPartNumbers(topCenterNode)
			}

			if hasRowAbove(row) && hasColToRight(col) {
				topRightNode := matrix[row-1][col+1]
				addToPartNumbers(topRightNode)
			}

			// sides, i.e. y

			if hasColToLeft(col) {
				leftNode := matrix[row][col-1]
				addToPartNumbers(leftNode)
			}

			if hasColToRight(col) {
				leftNode := matrix[row][col+1]
				addToPartNumbers(leftNode)
			}

			// below, i.e. y+1

			if hasRowBelow(row) && hasColToLeft(col) {
				bottomLeftNode := matrix[row+1][col-1]
				addToPartNumbers(bottomLeftNode)
			}

			if hasRowBelow(row) {
				bottomCenterNode := matrix[row+1][col]
				addToPartNumbers(bottomCenterNode)
			}

			if hasRowBelow(row) && hasColToRight(col) {
				bottomRightNode := matrix[row+1][col+1]
				addToPartNumbers(bottomRightNode)
			}

			if len(partNumbers) != 2 {
				continue
			}

			gearRatio := 1
			for _, n := range partNumbers {
				gearRatio *= n.value
			}
			sum += gearRatio
		}
	}

	return sum
}

type node struct {
	id         int64
	value      int
	isSymbol   bool
	isAsterisk bool // part 2 - only asterisks!
}

// type matrix struct {
// 	nodes [][]*node
// }

// func (m *matrix) hasRowBelow(row int) bool {

// }

// [][]{value, guid}
// then after i would want to do this to dedupe (don't want values added twice)
// guid -> value
// then sum the values
//
//nolint:gosec // Crypto of such magnitude not required
func makeNodeMatrix(lines []string) [][]*node {
	isDigit := func(c rune) bool {
		return c >= 48 && c < 58
	}

	isPeriod := func(c rune) bool {
		return c == 46
	}

	isAsterisk := func(c rune) bool {
		return c == 42
	}

	var nodeMatrix [][]*node

	getDigitNode := func(col, row int) *node {
		// if the passed in node is the Nth in a series of digits,
		// grab the N-1th node to modify
		if col > 0 && nodeMatrix[row][col-1].value > 0 {
			return nodeMatrix[row][col-1]
		}
		// otherwise make a new one
		return &node{id: rand.Int63(), value: 0, isSymbol: false}
	}

	putNodeInMatrix := func(n *node, col, row int) {
		if len(nodeMatrix) <= col {
			nodeMatrix = append(nodeMatrix, []*node{})
		}
		nodeMatrix[row] = append(nodeMatrix[row], n)
	}

	for row, line := range lines {
		for col, c := range line {
			switch {
			case isDigit(c):
				n := getDigitNode(col, row)
				n.value = n.value*10 + (int(c) - 48)
				putNodeInMatrix(n, col, row)
			case isPeriod(c):
				putNodeInMatrix(&node{id: rand.Int63(), value: 0, isSymbol: false}, col, row)
			case isAsterisk(c):
				putNodeInMatrix(&node{id: rand.Int63(), value: 0, isSymbol: true, isAsterisk: true}, col, row)
			default:
				putNodeInMatrix(&node{id: rand.Int63(), value: 0, isSymbol: true}, col, row)
			}
		}
	}

	return nodeMatrix
}

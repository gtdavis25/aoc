package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gtdavis25/aoc/internal/parse"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(context solver.Context) error {
	lines, err := context.InputLines()
	if err != nil {
		return err
	}

	equations := make(map[int][]int)
	for i, line := range lines {
		left, right, ok := strings.Cut(line, ": ")
		if !ok {
			return fmt.Errorf("line %d: could not parse %q", i, line)
		}

		result, err := strconv.Atoi(left)
		if err != nil {
			return fmt.Errorf("parsing %q as result on line %d: %w", left, i, err)
		}

		operands, err := parse.IntSlice(right, " ")
		if err != nil {
			return fmt.Errorf("line %d: %w", i, err)
		}

		equations[result] = operands
	}

	var part1 int
	for result, operands := range equations {
		if canMake(result, operands, []operator{add, multiply}) {
			part1 += result
		}
	}

	context.SetPart1(part1)
	var part2 int
	for result, operands := range equations {
		if canMake(result, operands, []operator{add, multiply, concatenate}) {
			part2 += result
		}
	}

	context.SetPart2(part2)
	return nil
}

type operator func(int, int) int

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func concatenate(x, y int) int {
	b := 10
	for y%b != y {
		b *= 10
	}

	return x*b + y
}

func canMake(result int, operands []int, operators []operator) bool {
	if len(operands) == 0 {
		return false
	}

	stack := []state{
		{
			acc:   operands[0],
			index: 0,
		},
	}

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if s.acc == result && s.index+1 == len(operands) {
			return true
		}

		if s.index+1 == len(operands) {
			continue
		}

		for _, operator := range operators {
			next := operator(s.acc, operands[s.index+1])
			stack = append(stack, state{
				acc:   next,
				index: s.index + 1,
			})
		}
	}

	return false
}

type state struct {
	acc   int
	index int
}

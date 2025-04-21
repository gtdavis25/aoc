package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	equations := make(map[int][]int)
	for i, line := range lines {
		left, right, ok := strings.Cut(line, ": ")
		if !ok {
			return solver.Result{}, fmt.Errorf("line %d: could not parse %q", i, line)
		}

		result, err := strconv.Atoi(left)
		if err != nil {
			return solver.Result{}, fmt.Errorf("parsing %q as result on line %d: %w", left, i, err)
		}

		words := strings.Split(right, " ")
		operands := make([]int, len(words))
		for j, word := range words {
			operands[j], err = strconv.Atoi(word)
			if err != nil {
				return solver.Result{}, fmt.Errorf("parsing %q as operand on line %d: %w", word, i, err)
			}
		}

		equations[result] = operands
	}

	var part1 int
	for result, operands := range equations {
		if canMake(result, operands) {
			part1 += result
		}
	}

	return solver.Result{
		Part1: part1,
	}, nil
}

func canMake(result int, operands []int) bool {
	switch {
	case len(operands) == 0:
		return false

	case len(operands) == 1:
		return operands[0] == result

	default:
		last := operands[len(operands)-1]
		rest := operands[:len(operands)-1]
		if d := result - last; d >= 0 && canMake(d, rest) {
			return true
		}

		if result%last == 0 && canMake(result/last, rest) {
			return true
		}

		return false
	}
}

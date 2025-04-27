package day03

import (
	"fmt"
	"strings"

	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	var part1 int
	for _, line := range lines {
		for i := range len(line) {
			var x, y int
			if _, err := fmt.Sscanf(line[i:], "mul(%d,%d)", &x, &y); err == nil {
				part1 += x * y
			}
		}
	}

	var part2 int
	var disabled bool
	for _, line := range lines {
		for i := range len(line) {
			switch {
			case strings.HasPrefix(line[i:], "don't()"):
				disabled = true

			case strings.HasPrefix(line[i:], "do()"):
				disabled = false

			case !disabled:
				var x, y int
				if _, err := fmt.Sscanf(line[i:], "mul(%d,%d)", &x, &y); err == nil {
					part2 += x * y
				}

			default:
			}
		}
	}

	return solver.Result{
		Part1: part1,
		Part2: part2,
	}, nil
}

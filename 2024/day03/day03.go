package day03

import (
	"fmt"
	"io"
	"strings"

	"github.com/gtdavis25/aoc/internal/input"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(r io.Reader, w io.Writer) error {
	lines, err := input.ReadLines(r)
	if err != nil {
		return err
	}

	var part1 int
	for _, line := range lines {
		for i := range len(line) {
			var x, y int
			if _, err := fmt.Sscanf(line[i:], "mul(%d,%d)", &x, &y); err == nil {
				part1 += x * y
			}
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
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

	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

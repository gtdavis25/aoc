package day19

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

	patterns := strings.Split(lines[0], ", ")
	var part1 int
	for _, design := range lines[2:] {
		if isPossible(design, patterns) {
			part1++
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	return nil
}

func isPossible(d string, patterns []string) bool {
	for _, p := range patterns {
		if d == p || strings.HasPrefix(d, p) && isPossible(d[len(p):], patterns) {
			return true
		}
	}

	return false
}

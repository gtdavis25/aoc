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
	var part1, part2 int
	for _, design := range lines[2:] {
		c := countWaysToMake(design, patterns, make(map[string]int))
		part2 += c
		if c > 0 {
			part1++
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

func countWaysToMake(design string, patterns []string, memo map[string]int) int {
	if c, ok := memo[design]; ok {
		return c
	}

	var count int
	for _, pattern := range patterns {
		switch {
		case design == pattern:
			count++

		case strings.HasPrefix(design, pattern):
			count += countWaysToMake(design[len(pattern):], patterns, memo)
		}
	}

	memo[design] = count
	return count
}

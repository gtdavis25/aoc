package day01

import (
	"fmt"
	"sort"

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

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		if _, err := fmt.Sscan(line, &left[i], &right[i]); err != nil {
			return fmt.Errorf("parsing %q on line %d: %w", line, i, err)
		}
	}

	sort.Ints(left)
	sort.Ints(right)
	var part1 int
	for i := range left {
		part1 += max(left[i]-right[i], right[i]-left[i])
	}

	context.SetPart1(part1)
	freq := make(map[int]int)
	for _, n := range right {
		freq[n]++
	}

	var part2 int
	for _, n := range left {
		part2 += n * freq[n]
	}

	context.SetPart2(part2)
	return nil
}

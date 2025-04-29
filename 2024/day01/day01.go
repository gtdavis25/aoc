package day01

import (
	"fmt"
	"io"
	"sort"

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

	fmt.Fprintf(w, "part 1: %d\n", part1)
	freq := make(map[int]int)
	for _, n := range right {
		freq[n]++
	}

	var part2 int
	for _, n := range left {
		part2 += n * freq[n]
	}

	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

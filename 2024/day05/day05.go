package day05

import (
	"fmt"
	"io"
	"slices"

	"github.com/gtdavis25/aoc/internal/input"
	"github.com/gtdavis25/aoc/internal/parse"
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

	successors := make(map[int][]int)
	var i int
	for i = range lines {
		if lines[i] == "" {
			break
		}

		var x, y int
		if _, err := fmt.Sscanf(lines[i], "%d|%d", &x, &y); err != nil {
			return fmt.Errorf("parsing %q on line %d: %w", lines[i], i, err)
		}

		successors[x] = append(successors[x], y)
	}

	i++
	updates := make([][]int, len(lines)-i)
	for j := range updates {
		pages, err := parse.IntSlice(lines[i+j], ",")
		if err != nil {
			return fmt.Errorf("line %d: %w", i+j, err)
		}

		updates[j] = pages
	}

	var part1, part2 int
	for _, pages := range updates {
		if isOrdered(pages, successors) {
			part1 += pages[len(pages)/2]
		} else {
			reorder(pages, successors)
			part2 += pages[len(pages)/2]
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

func isOrdered(pages []int, successors map[int][]int) bool {
	for i := range pages {
		for j := range pages[:i] {
			if slices.Contains(successors[pages[i]], pages[j]) {
				return false
			}
		}
	}

	return true
}

func reorder(pages []int, successors map[int][]int) {
	slices.SortFunc(pages, func(x, y int) int {
		switch {
		case slices.Contains(successors[x], y):
			return -1

		case slices.Contains(successors[y], x):
			return 1

		default:
			return 0
		}
	})
}

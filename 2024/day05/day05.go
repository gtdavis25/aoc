package day05

import (
	"fmt"
	"slices"

	"github.com/gtdavis25/aoc/parse"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	successors := make(map[int][]int)
	var i int
	for i = range lines {
		if lines[i] == "" {
			break
		}

		var x, y int
		if _, err := fmt.Sscanf(lines[i], "%d|%d", &x, &y); err != nil {
			return solver.Result{}, fmt.Errorf("parsing %q on line %d: %w", lines[i], i, err)
		}

		successors[x] = append(successors[x], y)
	}

	i++
	updates := make([][]int, len(lines)-i)
	for j := range updates {
		pages, err := parse.IntSlice(lines[i+j], ",")
		if err != nil {
			return solver.Result{}, fmt.Errorf("line %d: %w", i+j, err)
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

	return solver.Result{
		Part1: part1,
		Part2: part2,
	}, nil
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

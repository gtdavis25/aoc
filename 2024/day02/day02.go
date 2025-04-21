package day02

import (
	"fmt"
	"strings"

	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	reports := make([][]int, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		reports[i] = make([]int, len(words))
		for j, word := range words {
			if _, err := fmt.Sscan(word, &reports[i][j]); err != nil {
				return solver.Result{}, fmt.Errorf("parsing %q on line %d position %d: %w", word, i, j, err)
			}
		}
	}

	var part1 int
	for _, report := range reports {
		if isSafe(report) {
			part1++
		}
	}

	var part2 int
	var modified []int
	for _, report := range reports {
		var safe bool
		for j := range report {
			modified = append(modified[:0], report[:j]...)
			modified = append(modified, report[j+1:]...)
			if isSafe(modified) {
				safe = true
				break
			}
		}

		if safe {
			part2++
		}
	}

	return solver.Result{
		Part1: part1,
		Part2: part2,
	}, nil
}

func isSafe(report []int) bool {
	for i := range len(report) - 1 {
		if d := abs(report[i+1] - report[i]); d == 0 || d > 3 {
			return false
		}

		if i > 0 && sign(report[i+1]-report[i]) != sign(report[i]-report[i-1]) {
			return false
		}
	}

	return true
}

func abs(n int) int {
	return max(n, -n)
}

func sign(n int) int {
	switch {
	case n < 0:
		return -1

	case n > 0:
		return 1

	default:
		return 0
	}
}

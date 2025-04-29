package day04

import (
	"fmt"
	"io"

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
	for _, pattern := range [][]string{
		{
			"XMAS",
		},
		{
			"SAMX",
		},
		{
			"X",
			"M",
			"A",
			"S",
		},
		{
			"S",
			"A",
			"M",
			"X",
		},
		{
			"X   ",
			" M  ",
			"  A ",
			"   S",
		},
		{
			"S   ",
			" A  ",
			"  M ",
			"   X",
		},
		{
			"   X",
			"  M ",
			" A  ",
			"S   ",
		},
		{
			"   S",
			"  A ",
			" M  ",
			"X   ",
		},
	} {
		part1 += countMatches(lines, pattern)
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	var part2 int
	for _, pattern := range [][]string{
		{
			"M S",
			" A ",
			"M S",
		},
		{
			"M M",
			" A ",
			"S S",
		},
		{
			"S M",
			" A ",
			"S M",
		},
		{
			"S S",
			" A ",
			"M M",
		},
	} {
		part2 += countMatches(lines, pattern)
	}

	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

func countMatches(lines, pattern []string) int {
	var count int
	for y := range len(lines) - len(pattern) + 1 {
		for x := range len(lines[y]) - len(pattern[0]) + 1 {
			if isMatch(lines, pattern, x, y) {
				count++
			}
		}
	}

	return count
}

func isMatch(lines, pattern []string, x, y int) bool {
	for dy := range len(pattern) {
		for dx := range len(pattern[dy]) {
			if pattern[dy][dx] != ' ' && pattern[dy][dx] != lines[y+dy][x+dx] {
				return false
			}
		}
	}

	return true
}

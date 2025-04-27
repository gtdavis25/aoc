package day11

import (
	"github.com/gtdavis25/aoc/internal/parse"
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

	stones, err := parse.IntSlice(lines[0], " ")
	if err != nil {
		return err
	}

	memo := make(memo)
	var part1, part2 int
	for _, n := range stones {
		part1 += getStoneCountAfterIterations(memo, n, 25)
		part2 += getStoneCountAfterIterations(memo, n, 75)
	}

	context.SetPart1(part1)
	context.SetPart2(part2)
	return nil
}

type memo map[key]int

type key struct {
	n int
	i int
}

func getStoneCountAfterIterations(memo memo, n, i int) int {
	if i == 0 {
		return 1
	}

	key := key{
		n: n,
		i: i,
	}

	if r, ok := memo[key]; ok {
		return r
	}

	var r int
	for _, next := range nextState(n) {
		r += getStoneCountAfterIterations(memo, next, i-1)
	}

	memo[key] = r
	return r
}

func nextState(n int) []int {
	if n == 0 {
		return []int{1}
	}

	if left, right, ok := splitDigits(n); ok {
		return []int{left, right}
	}

	return []int{n * 2024}
}

func splitDigits(n int) (int, int, bool) {
	b, l, u := 10, 10, 100
	for {
		if n < l {
			return 0, 0, false
		}

		if n < u {
			return n / b, n % b, true
		}

		b, l, u = b*10, l*100, u*100
	}
}

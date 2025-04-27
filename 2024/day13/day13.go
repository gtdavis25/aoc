package day13

import (
	"fmt"

	"github.com/gtdavis25/aoc/internal/geom2d"
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

	var games []game
	for i := 0; i+2 < len(lines); i += 4 {
		var g game
		if _, err := fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &g.a.X, &g.a.Y); err != nil {
			return fmt.Errorf("parsing %q: %w", lines[i], err)
		}

		if _, err := fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &g.b.X, &g.b.Y); err != nil {
			return fmt.Errorf("parsing %q: %w", lines[i+1], err)
		}

		if _, err := fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &g.p.X, &g.p.Y); err != nil {
			return fmt.Errorf("parsing %q: %w", lines[i+2], err)
		}

		games = append(games, g)
	}

	var part1, part2 int
	for _, g := range games {
		part1 += getMinimumCost(g)
		g.p.X += 10000000000000
		g.p.Y += 10000000000000
		part2 += getMinimumCost(g)
	}

	context.SetPart1(part1)
	context.SetPart2(part2)
	return nil
}

type game struct {
	a geom2d.Point
	b geom2d.Point
	p geom2d.Point
}

func getMinimumCost(g game) int {
	p := g.a.Y*g.p.X - g.a.X*g.p.Y
	q := g.b.X*g.a.Y - g.a.X*g.b.Y
	if q == 0 || p%q != 0 {
		return 0
	}

	b := p / q
	a := (g.p.X - b*g.b.X) / g.a.X
	return 3*a + b
}

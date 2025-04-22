package day10

import (
	"slices"

	"github.com/gtdavis25/aoc/geom2d"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	var queue []trail
	for y := range lines {
		for x := range len(lines[y]) {
			if lines[y][x] == '0' {
				queue = append(queue, trail{
					start: geom2d.Point{
						X: x,
						Y: y,
					},
					end: geom2d.Point{
						X: x,
						Y: y,
					},
				})
			}
		}
	}

	trails := make(map[geom2d.Point][]geom2d.Point)
	for i := 0; i < len(queue); i++ {
		t := queue[i]
		if lines[t.end.Y][t.end.X] == '9' {
			if !slices.Contains(trails[t.start], t.end) {
				trails[t.start] = append(trails[t.start], t.end)
			}

			continue
		}

		for _, p := range []geom2d.Point{
			{
				X: t.end.X,
				Y: t.end.Y - 1,
			},
			{
				X: t.end.X + 1,
				Y: t.end.Y,
			},
			{
				X: t.end.X,
				Y: t.end.Y + 1,
			},
			{
				X: t.end.X - 1,
				Y: t.end.Y,
			},
		} {
			if p.Y < 0 || len(lines) <= p.Y || p.X < 0 || len(lines[p.Y]) <= p.X ||
				lines[p.Y][p.X]-lines[t.end.Y][t.end.X] != 1 {
				continue
			}

			queue = append(queue, trail{
				start: t.start,
				end:   p,
			})
		}
	}

	var part1 int
	for _, ends := range trails {
		part1 += len(ends)
	}

	return solver.Result{
		Part1: part1,
	}, nil
}

type trail struct {
	start geom2d.Point
	end   geom2d.Point
}

package day14

import (
	"fmt"

	"github.com/gtdavis25/aoc/geom2d"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	robots := make([]robot, len(lines))
	for i, line := range lines {
		var pos, vel geom2d.Point
		if _, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pos.X, &pos.Y, &vel.X, &vel.Y); err != nil {
			return solver.Result{}, fmt.Errorf("parsing %q on line %d: %w", line, i, err)
		}

		robots[i] = robot{
			pos: pos,
			vel: vel,
		}
	}

	width, height := 101, 103
	for range 100 {
		for i, r := range robots {
			x := (r.pos.X + r.vel.X) % width
			if x < 0 {
				x += width
			}

			y := (r.pos.Y + r.vel.Y) % height
			if y < 0 {
				y += height
			}

			robots[i].pos = geom2d.Point{
				X: x,
				Y: y,
			}
		}
	}

	var quadrants [2][2]int
	for _, r := range robots {
		switch {
		case r.pos.X < width/2 && r.pos.Y < height/2:
			quadrants[0][0]++

		case r.pos.X > width/2 && r.pos.Y < height/2:
			quadrants[0][1]++

		case r.pos.X < width/2 && r.pos.Y > height/2:
			quadrants[1][0]++

		case r.pos.X > width/2 && r.pos.Y > height/2:
			quadrants[1][1]++
		}
	}

	return solver.Result{
		Part1: quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1],
	}, nil
}

type robot struct {
	pos geom2d.Point
	vel geom2d.Point
}

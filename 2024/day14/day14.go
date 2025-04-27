package day14

import (
	"fmt"
	"slices"

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

	initial := make([]robot, len(lines))
	for i, line := range lines {
		var pos, vel geom2d.Point
		if _, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pos.X, &pos.Y, &vel.X, &vel.Y); err != nil {
			return fmt.Errorf("parsing %q on line %d: %w", line, i, err)
		}

		initial[i] = robot{
			pos: pos,
			vel: vel,
		}
	}

	width, height := 101, 103
	robots := slices.Clone(initial)
	for range 100 {
		nextState(robots, width, height)
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

	context.SetPart1(quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1])
	var part2 int
	robots = slices.Clone(initial)
	for hasDuplicatePositions(robots) {
		nextState(robots, width, height)
		part2++
	}

	context.SetPart2(part2)
	return nil
}

type robot struct {
	pos geom2d.Point
	vel geom2d.Point
}

func nextState(robots []robot, width, height int) {
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

func hasDuplicatePositions(robots []robot) bool {
	for i := range robots {
		for j := range robots[:i] {
			if robots[i].pos == robots[j].pos {
				return true
			}
		}
	}

	return false
}

package day06

import (
	"fmt"

	"github.com/gtdavis25/aoc/geom2d"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(input []string) (solver.Result, error) {
	lines := make([][]byte, len(input))
	for i, line := range input {
		lines[i] = []byte(line)
	}

	initial, err := getInitialState(lines)
	if err != nil {
		return solver.Result{}, fmt.Errorf("getting initial state: %w", err)
	}

	seen := make(map[geom2d.Point]struct{})
	for current := initial; ; {
		seen[current.pos] = struct{}{}
		next, ok := nextState(current, lines)
		if !ok {
			break
		}

		if next == current {
			return solver.Result{}, fmt.Errorf("guard stuck")
		}

		current = next
	}

	var part2 int
	for p := range seen {
		if p == initial.pos {
			continue
		}

		lines[p.Y][p.X] = '#'
		if reachesLoop(initial, lines) {
			part2++
		}

		lines[p.Y][p.X] = '.'
	}

	return solver.Result{
		Part1: len(seen),
		Part2: part2,
	}, nil
}

type state struct {
	pos geom2d.Point
	vel geom2d.Point
}

func getInitialState(lines [][]byte) (state, error) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return state{
					pos: geom2d.Point{
						X: x,
						Y: y,
					},
					vel: geom2d.Up(),
				}, nil
			}
		}
	}

	return state{}, fmt.Errorf("no start position")
}

func nextState(current state, lines [][]byte) (state, bool) {
	vel := current.vel
	for range 4 {
		nextPos := current.pos.Add(vel)
		switch {
		case nextPos.Y < 0 || len(lines) <= nextPos.Y || nextPos.X < 0 || len(lines[nextPos.Y]) <= nextPos.X:
			return state{}, false

		case lines[nextPos.Y][nextPos.X] != '#':
			return state{
				pos: nextPos,
				vel: vel,
			}, true

		default:
			vel = geom2d.Point{
				X: -vel.Y,
				Y: vel.X,
			}
		}
	}

	return current, true
}

func reachesLoop(initial state, lines [][]byte) bool {
	seen := make(map[state]struct{})
	for current := initial; ; {
		if _, ok := seen[current]; ok {
			return true
		}

		seen[current] = struct{}{}
		next, ok := nextState(current, lines)
		if !ok {
			return false
		}

		current = next
	}
}

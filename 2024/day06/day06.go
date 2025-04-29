package day06

import (
	"fmt"
	"io"

	"github.com/gtdavis25/aoc/internal/geom2d"
	"github.com/gtdavis25/aoc/internal/input"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(r io.Reader, w io.Writer) error {
	input, err := input.ReadLines(r)
	if err != nil {
		return err
	}

	lines := make([][]byte, len(input))
	for i, line := range input {
		lines[i] = []byte(line)
	}

	initial, err := getInitialState(lines)
	if err != nil {
		return fmt.Errorf("getting initial state: %w", err)
	}

	seen := make(map[geom2d.Point]struct{})
	for current := initial; ; {
		seen[current.pos] = struct{}{}
		next, ok := nextState(current, lines)
		if !ok {
			break
		}

		if next == current {
			return fmt.Errorf("guard stuck")
		}

		current = next
	}

	fmt.Fprintf(w, "part 1: %d\n", len(seen))
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

	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
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
	bounds := geom2d.Rect{Width: len(lines[0]), Height: len(lines)}
	vel := current.vel
	for range 4 {
		nextPos := current.pos.Add(vel)
		switch {
		case !bounds.Contains(nextPos):
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

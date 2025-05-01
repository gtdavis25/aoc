package day18

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
	lines, err := input.ReadLines(r)
	if err != nil {
		return err
	}

	bytes := make([]geom2d.Point, 0, len(lines))
	for i, line := range lines {
		var p geom2d.Point
		if _, err := fmt.Sscanf(line, "%d,%d", &p.X, &p.Y); err != nil {
			return fmt.Errorf("line %q on line %d: %w", line, i, err)
		}

		bytes = append(bytes, p)
	}

	part1, ok := getShortestPath(bytes[:1024], 71, 71)
	if !ok {
		return fmt.Errorf("part 1: path not found")
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	return nil
}

func getShortestPath(bytes []geom2d.Point, w, h int) (int, bool) {
	bounds := geom2d.Rect{
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
	}

	corrupted := make(map[geom2d.Point]struct{})
	for _, p := range bytes {
		corrupted[p] = struct{}{}
	}

	seen := make(map[geom2d.Point]struct{})
	queue := []state{
		{
			pos: geom2d.Origin(),
			t:   0,
		},
	}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s.pos.X == w-1 && s.pos.Y == h-1 {
			return s.t, true
		}

		if _, ok := seen[s.pos]; ok {
			continue
		}

		seen[s.pos] = struct{}{}
		for next := range s.pos.Adjacent() {
			if _, ok := corrupted[next]; ok || !bounds.Contains(next) {
				continue
			}

			queue = append(queue, state{
				pos: next,
				t:   s.t + 1,
			})
		}
	}

	return 0, false
}

type state struct {
	pos geom2d.Point
	t   int
}

package day18

import (
	"fmt"
	"io"
	"slices"

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

	path, ok := getShortestPath(bytes[:1024], 71, 71)
	if !ok {
		return fmt.Errorf("part 1: path not found")
	}

	fmt.Fprintf(w, "part 1: %d\n", len(path)-1)
	for i := range len(bytes) - 1024 {
		if slices.Contains(path, bytes[1024+i]) {
			var ok bool
			if path, ok = getShortestPath(bytes[:1025+i], 71, 71); !ok {
				fmt.Fprintf(w, "part 2: %d,%d\n", bytes[1024+i].X, bytes[1024+i].Y)
				return nil
			}
		}
	}

	return fmt.Errorf("solution to part 2 not found")
}

func getShortestPath(bytes []geom2d.Point, w, h int) ([]geom2d.Point, bool) {
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
			pos:  geom2d.Origin(),
			t:    0,
			prev: nil,
		},
	}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s.pos.X == w-1 && s.pos.Y == h-1 {
			return getPath(&s), true
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
				pos:  next,
				t:    s.t + 1,
				prev: &s,
			})
		}
	}

	return nil, false
}

type state struct {
	pos  geom2d.Point
	t    int
	prev *state
}

func getPath(end *state) []geom2d.Point {
	var points []geom2d.Point
	for s := end; s != nil; s = s.prev {
		points = append(points, s.pos)
	}

	return points
}

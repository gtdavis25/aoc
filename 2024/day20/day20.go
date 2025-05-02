package day20

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
	lines, err := input.ReadLinesBytes(r)
	if err != nil {
		return err
	}

	shortestPath, err := getShortestPath(lines)
	if err != nil {
		return err
	}

	var part1 int
	for y := 1; y+1 < len(lines); y++ {
		for x := 1; x+1 < len(lines[y]); x++ {
			if lines[y][x] != '#' {
				continue
			}

			lines[y][x] = '.'
			t, err := getShortestPath(lines)
			if err != nil {
				return err
			}

			if t+100 <= shortestPath {
				part1++
			}

			lines[y][x] = '#'
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	return nil
}

func getShortestPath(lines [][]byte) (int, error) {
	start, err := getStartPosition(lines)
	if err != nil {
		return 0, err
	}

	seen := make(map[geom2d.Point]struct{})
	queue := []state{{pos: start, t: 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if lines[s.pos.Y][s.pos.X] == 'E' {
			return s.t, nil
		}

		if _, ok := seen[s.pos]; ok {
			continue
		}

		seen[s.pos] = struct{}{}
		for p := range s.pos.Adjacent() {
			if lines[p.Y][p.X] == '#' {
				continue
			}

			queue = append(queue, state{
				pos: p,
				t:   s.t + 1,
			})
		}
	}

	return 0, fmt.Errorf("end unreachable")
}

func getStartPosition(lines [][]byte) (geom2d.Point, error) {
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				return geom2d.Point{
					X: x,
					Y: y,
				}, nil
			}
		}
	}

	return geom2d.Point{}, fmt.Errorf("no start position")
}

type state struct {
	pos geom2d.Point
	t   int
}

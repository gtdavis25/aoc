package day20

import (
	"fmt"
	"io"
	"iter"
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

	start, ok := findPosition(lines, 'S')
	if !ok {
		return fmt.Errorf("no start position")
	}

	end, ok := findPosition(lines, 'E')
	if !ok {
		return fmt.Errorf("no end position")
	}

	dStart := getDistances(lines, start)
	dEnd := getDistances(lines, end)
	budget := dEnd[start.Y][start.X] - 100
	bounds := geom2d.Rect{Width: len(lines[0]), Height: len(lines)}
	var part1 int
	for y := 1; y+1 < bounds.Height; y++ {
		for x := 1; x+1 < bounds.Width; x++ {
			if lines[y][x] == '#' {
				continue
			}

			s := geom2d.Point{X: x, Y: y}
			for e := range getCheatDestinations(lines, s, 2) {
				if dStart[s.Y][s.X]+geom2d.GetDistance(s, e)+dEnd[e.Y][e.X] <= budget {
					part1++
				}
			}
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	return nil
}

func findPosition(lines []string, c byte) (geom2d.Point, bool) {
	for y, line := range lines {
		for x := range len(line) {
			if line[x] == c {
				return geom2d.Point{
					X: x,
					Y: y,
				}, true
			}
		}
	}

	return geom2d.Point{}, false
}

func getDistances(lines []string, origin geom2d.Point) [][]int {
	distances := make([][]int, len(lines))
	for y, line := range lines {
		distances[y] = slices.Repeat([]int{-1}, len(line))
	}

	seen := make(map[geom2d.Point]struct{})
	queue := []state{{pos: origin, t: 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if _, ok := seen[s.pos]; ok {
			continue
		}

		seen[s.pos] = struct{}{}
		distances[s.pos.Y][s.pos.X] = s.t
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

	return distances
}

func getCheatDestinations(lines []string, start geom2d.Point, maxSteps int) iter.Seq[geom2d.Point] {
	return func(yield func(geom2d.Point) bool) {
		bounds := geom2d.Rect{Width: len(lines[0]), Height: len(lines)}
		seen := make(map[geom2d.Point]struct{})
		queue := []state{{pos: start, t: 0}}
		for len(queue) > 0 {
			s := queue[0]
			queue = queue[1:]
			if _, ok := seen[s.pos]; ok {
				continue
			}

			seen[s.pos] = struct{}{}
			if lines[s.pos.Y][s.pos.X] != '#' && !yield(s.pos) {
				return
			}

			if s.t >= maxSteps {
				continue
			}

			for p := range s.pos.Adjacent() {
				if _, ok := seen[p]; ok || !bounds.Contains(p) {
					continue
				}

				queue = append(queue, state{
					pos: p,
					t:   s.t + 1,
				})
			}
		}
	}
}

type state struct {
	pos geom2d.Point
	t   int
}

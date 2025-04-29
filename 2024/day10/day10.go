package day10

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

	trailheads := getTrailheads(lines)
	var part1, part2 int
	for _, trailhead := range trailheads {
		part1 += trailhead.score
		part2 += trailhead.rating
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
}

type trailhead struct {
	score  int
	rating int
}

func getTrailheads(lines []string) []trailhead {
	var trailheads []trailhead
	for y, line := range lines {
		for x := range len(line) {
			if lines[y][x] != '0' {
				continue
			}

			trailEnds := getTrailEnds(lines, geom2d.Point{X: x, Y: y})
			trailheads = append(trailheads, trailhead{
				score:  distinctCount(trailEnds),
				rating: len(trailEnds),
			})
		}
	}

	return trailheads
}

func getTrailEnds(lines []string, p geom2d.Point) []geom2d.Point {
	if lines[p.Y][p.X] == '9' {
		return []geom2d.Point{p}
	}

	bounds := geom2d.Rect{Width: len(lines[0]), Height: len(lines)}
	var trailEnds []geom2d.Point
	for next := range p.Adjacent() {
		if !bounds.Contains(next) || lines[next.Y][next.X]-lines[p.Y][p.X] != 1 {
			continue
		}

		trailEnds = append(trailEnds, getTrailEnds(lines, next)...)
	}

	return trailEnds
}

func distinctCount(points []geom2d.Point) int {
	var count int
	for i := range points {
		if !slices.Contains(points[:i], points[i]) {
			count++
		}
	}

	return count
}

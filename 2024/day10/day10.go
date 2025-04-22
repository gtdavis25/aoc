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
	trailheads := getTrailheads(lines)
	var part1, part2 int
	for _, trailhead := range trailheads {
		part1 += trailhead.score
		part2 += trailhead.rating
	}

	return solver.Result{
		Part1: part1,
		Part2: part2,
	}, nil
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

	var trailEnds []geom2d.Point
	for _, next := range []geom2d.Point{
		{
			X: p.X,
			Y: p.Y - 1,
		},
		{
			X: p.X + 1,
			Y: p.Y,
		},
		{
			X: p.X,
			Y: p.Y + 1,
		},
		{
			X: p.X - 1,
			Y: p.Y,
		},
	} {
		if next.Y < 0 || next.Y >= len(lines) || next.X < 0 || next.X >= len(lines[next.Y]) ||
			lines[next.Y][next.X]-lines[p.Y][p.X] != 1 {
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

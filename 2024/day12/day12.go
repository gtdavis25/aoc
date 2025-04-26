package day12

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
	regions := getRegions(lines)
	var part1 int
	for _, r := range regions {
		part1 += getArea(r) * getPerimeter(r)
	}

	return solver.Result{
		Part1: part1,
	}, nil
}

func getRegions(lines []string) []region {
	var regions []region
	visited := make(map[geom2d.Point]struct{})
	for y, line := range lines {
		for x := range len(line) {
			p := geom2d.Point{X: x, Y: y}
			if _, ok := visited[p]; ok {
				continue
			}

			points := getPoints(lines, p)
			for _, p := range points {
				visited[p] = struct{}{}
			}

			regions = append(regions, region{points: points})
		}
	}

	return regions
}

func getPoints(lines []string, start geom2d.Point) []geom2d.Point {
	bounds := geom2d.Rect{Width: len(lines[0]), Height: len(lines)}
	points := []geom2d.Point{start}
	for i := 0; i < len(points); i++ {
		p := points[i]
		for n := range p.Adjacent() {
			if !bounds.Contains(n) || slices.Contains(points, n) || lines[n.Y][n.X] != lines[p.Y][p.X] {
				continue
			}

			points = append(points, n)
		}
	}

	return points
}

type region struct {
	points []geom2d.Point
}

func getArea(r region) int {
	return len(r.points)
}

func getPerimeter(r region) int {
	var perimeter int
	for _, p := range r.points {
		for n := range p.Adjacent() {
			if !slices.Contains(r.points, n) {
				perimeter++
			}
		}
	}

	return perimeter
}

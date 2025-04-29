package day12

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

	regions := getRegions(lines)
	var part1, part2 int
	for _, r := range regions {
		area := getArea(r)
		part1 += area * getPerimeter(r)
		part2 += area * getSideCount(r)
	}

	fmt.Fprintf(w, "part 1: %d\n", part1)
	fmt.Fprintf(w, "part 2: %d\n", part2)
	return nil
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

func getSideCount(r region) int {
	var sideCount int
	for d := range geom2d.Origin().Adjacent() {
		var boundaries []geom2d.Point
		for _, p := range r.points {
			if !slices.Contains(r.points, p.Add(d)) {
				boundaries = append(boundaries, p)
			}
		}

		groups := groupByAdjacency(boundaries)
		sideCount += len(groups)
	}

	return sideCount
}

func groupByAdjacency(points []geom2d.Point) [][]geom2d.Point {
	var groups [][]geom2d.Point
	for _, p1 := range points {
		g := []geom2d.Point{p1}
		for _, group := range groups {
			for _, p2 := range group {
				if geom2d.GetDistance(p1, p2) == 1 {
					g = append(g, group...)
				}
			}
		}

		newGroups := [][]geom2d.Point{g}
		for _, group := range groups {
			if !containsAny(group, g) {
				newGroups = append(newGroups, group)
			}
		}

		groups = newGroups
	}

	return groups
}

func containsAny(s1, s2 []geom2d.Point) bool {
	return slices.ContainsFunc(s1, func(p geom2d.Point) bool {
		return slices.Contains(s2, p)
	})
}

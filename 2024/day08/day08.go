package day08

import (
	"github.com/gtdavis25/aoc/geom2d"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	frequencies := make(map[byte][]geom2d.Point)
	for y, line := range lines {
		for x := range len(line) {
			if line[x] != '.' {
				frequencies[line[x]] = append(frequencies[line[x]], geom2d.Point{
					X: x,
					Y: y,
				})
			}
		}
	}

	antinodes := make(map[geom2d.Point]struct{})
	for _, antennas := range frequencies {
		for i, p1 := range antennas {
			for _, p2 := range antennas[:i] {
				d := p2.Subtract(p1)
				for _, antinode := range []geom2d.Point{
					p1.Subtract(d),
					p2.Add(d),
				} {
					if 0 <= antinode.Y && antinode.Y < len(lines) && 0 <= antinode.X && antinode.X < len(lines[antinode.Y]) {
						antinodes[antinode] = struct{}{}
					}
				}
			}
		}
	}

	return solver.Result{
		Part1: len(antinodes),
	}, nil
}

package day08

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
					if inBounds(antinode, lines) {
						antinodes[antinode] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Fprintf(w, "part 1: %d\n", len(antinodes))
	for _, antennas := range frequencies {
		for i, p1 := range antennas {
			for _, p2 := range antennas[:i] {
				d := p2.Subtract(p1)
				step := geom2d.Point{
					X: d.X / gcd(d.X, d.Y),
					Y: d.Y / gcd(d.X, d.Y),
				}

				for p := p1; inBounds(p, lines); p = p.Add(step) {
					antinodes[p] = struct{}{}
				}

				for p := p2; inBounds(p, lines); p = p.Subtract(step) {
					antinodes[p] = struct{}{}
				}
			}
		}
	}

	fmt.Fprintf(w, "part 2: %d\n", len(antinodes))
	return nil
}

func inBounds(p geom2d.Point, lines []string) bool {
	return 0 <= p.Y && p.Y < len(lines) && 0 <= p.X && p.X < len(lines[p.Y])
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return abs(gcd(b, a%b))
}

func abs(n int) int {
	return max(n, -n)
}

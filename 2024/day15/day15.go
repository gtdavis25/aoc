package day15

import (
	"fmt"
	"strings"

	"github.com/gtdavis25/aoc/geom2d"
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	var rows [][]byte
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		}

		rows = append(rows, []byte(line))
	}

	position, err := getStartPosition(rows)
	if err != nil {
		return solver.Result{}, err
	}

	movements := strings.Join(lines, "")
	for _, c := range movements {
		var direction geom2d.Point
		switch c {
		case '^':
			direction = geom2d.Up()

		case '>':
			direction = geom2d.Right()

		case 'v':
			direction = geom2d.Down()

		case '<':
			direction = geom2d.Left()

		default:
			return solver.Result{}, fmt.Errorf("invalid direction: %c", c)
		}

		if move(rows, position, direction) {
			position = position.Add(direction)
		}
	}

	var part1 int
	for y, row := range rows {
		for x, c := range row {
			if c == 'O' {
				part1 += 100*y + x
			}
		}
	}

	return solver.Result{
		Part1: part1,
	}, nil
}

func getStartPosition(rows [][]byte) (geom2d.Point, error) {
	for y, row := range rows {
		for x, c := range row {
			if c == '@' {
				return geom2d.Point{
					X: x,
					Y: y,
				}, nil
			}
		}
	}

	return geom2d.Point{}, fmt.Errorf("no start position")
}

func move(rows [][]byte, p geom2d.Point, d geom2d.Point) bool {
	n := p.Add(d)
	if rows[n.Y][n.X] == '#' || rows[n.Y][n.X] == 'O' && !move(rows, n, d) {
		return false
	}

	rows[p.Y][p.X], rows[n.Y][n.X] = rows[n.Y][n.X], rows[p.Y][p.X]
	return true
}

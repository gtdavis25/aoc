package day15

import (
	"fmt"
	"strings"

	"github.com/gtdavis25/aoc/internal/geom2d"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	var rows [][]byte
	var delimiter int
	for i, line := range lines {
		if line == "" {
			delimiter = i
			break
		}

		rows = append(rows, []byte(line))
	}

	moves, err := GetMoves(strings.Join(lines[delimiter+1:], ""))
	if err := DoMoves(rows, moves); err != nil {
		return solver.Result{}, err
	}

	part1 := getGPSCoordinates(rows)
	rows, err = updateMap(lines[:delimiter])
	if err != nil {
		return solver.Result{}, fmt.Errorf("updating map: %w", err)
	}

	if err := DoMoves(rows, moves); err != nil {
		return solver.Result{}, err
	}

	part2 := getGPSCoordinates(rows)
	return solver.Result{
		Part1: part1,
		Part2: part2,
	}, nil
}

func GetMoves(directions string) ([]geom2d.Point, error) {
	moves := make([]geom2d.Point, len(directions))
	for i := range moves {
		switch directions[i] {
		case '^':
			moves[i] = geom2d.Up()

		case '>':
			moves[i] = geom2d.Right()

		case 'v':
			moves[i] = geom2d.Down()

		case '<':
			moves[i] = geom2d.Left()

		default:
			return nil, fmt.Errorf("unexpected character: %c", directions[i])
		}
	}

	return moves, nil
}

func DoMoves(rows [][]byte, moves []geom2d.Point) error {
	pos, err := getStartPosition(rows)
	if err != nil {
		return err
	}

	for _, d := range moves {
		if canMove(rows, pos, d) {
			pos = move(rows, pos, d)
		}
	}

	return nil
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

func canMove(rows [][]byte, p geom2d.Point, d geom2d.Point) bool {
	n := p.Add(d)
	if rows[n.Y][n.X] == '#' {
		return false
	}

	if d.X == 0 && rows[n.Y][n.X] == ']' && !canMove(rows, n.Add(geom2d.Left()), d) {
		return false
	}

	if d.X == 0 && rows[n.Y][n.X] == '[' && !canMove(rows, n.Add(geom2d.Right()), d) {
		return false
	}

	if isBox(rows, n) && !canMove(rows, n, d) {
		return false
	}

	return true
}

func move(rows [][]byte, p geom2d.Point, d geom2d.Point) geom2d.Point {
	n := p.Add(d)
	if d.X == 0 && rows[n.Y][n.X] == ']' {
		move(rows, n.Add(geom2d.Left()), d)
	}

	if d.X == 0 && rows[n.Y][n.X] == '[' {
		move(rows, n.Add(geom2d.Right()), d)
	}

	if isBox(rows, n) {
		move(rows, n, d)
	}

	rows[p.Y][p.X], rows[n.Y][n.X] = rows[n.Y][n.X], rows[p.Y][p.X]
	return n
}

func getGPSCoordinates(rows [][]byte) int {
	var result int
	for y, row := range rows {
		for x, c := range row {
			if c == 'O' || c == '[' {
				result += 100*y + x
			}
		}
	}

	return result
}

func updateMap(lines []string) ([][]byte, error) {
	updated := make([][]byte, len(lines))
	for y, row := range lines {
		updated[y] = make([]byte, 0, 2*len(row))
		for _, c := range row {
			switch c {
			case '#':
				updated[y] = append(updated[y], "##"...)

			case '.':
				updated[y] = append(updated[y], ".."...)

			case 'O':
				updated[y] = append(updated[y], "[]"...)

			case '@':
				updated[y] = append(updated[y], "@."...)

			default:
				return nil, fmt.Errorf("unexpected character: %c", c)
			}
		}
	}

	return updated, nil
}

func isBox(rows [][]byte, p geom2d.Point) bool {
	switch rows[p.Y][p.X] {
	case 'O', '[', ']':
		return true

	default:
		return false
	}
}

package day09

import (
	"github.com/gtdavis25/aoc/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	var blocks []int
	var empty bool
	var fileID int
	for i := range len(lines[0]) {
		length := int(lines[0][i] - '0')
		if empty {
			for range length {
				blocks = append(blocks, -1)
			}
		} else {
			for range length {
				blocks = append(blocks, fileID)
			}

			fileID++
		}

		empty = !empty
	}

	for i, j := 0, len(blocks)-1; i < j; {
		switch {
		case blocks[i] != -1:
			i++

		case blocks[j] == -1:
			blocks = blocks[:j]
			j--

		default:
			blocks[i] = blocks[j]
			blocks[j] = -1
		}
	}

	var checksum int
	for i := range len(blocks) {
		if blocks[i] != -1 {
			checksum += i * blocks[i]
		}
	}

	return solver.Result{
		Part1: checksum,
	}, nil
}

package day09

import (
	"slices"

	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(lines []string) (solver.Result, error) {
	return solver.Result{
		Part1: part1(lines[0]),
		Part2: part2(lines[0]),
	}, nil
}

func part1(input string) int {
	var blocks []int
	var empty bool
	var fileID int
	for i := range len(input) {
		length := int(input[i] - '0')
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

	return checksum
}

func part2(input string) int {
	var blocks []block
	var fileID int
	var offset int
	for i := range len(input) {
		length := int(input[i] - '0')
		if i%2 == 0 {
			blocks = append(blocks, block{
				fileID: fileID,
				offset: offset,
				length: length,
			})

			fileID++
		} else {
			blocks = append(blocks, block{
				fileID: -1,
				offset: offset,
				length: length,
			})
		}

		offset += length
	}

	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i].fileID == -1 {
			continue
		}

		for j := range blocks[:i] {
			if blocks[j].fileID != -1 || blocks[j].length < blocks[i].length {
				continue
			}

			fileBlock := blocks[i]
			blocks[i].fileID = -1
			emptyBlock := blocks[j]
			blocks[j] = block{
				fileID: fileBlock.fileID,
				offset: emptyBlock.offset,
				length: fileBlock.length,
			}

			if fileBlock.length < emptyBlock.length {
				i++
				blocks = slices.Insert(blocks, j+1, block{
					fileID: -1,
					offset: emptyBlock.offset + fileBlock.length,
					length: emptyBlock.length - fileBlock.length,
				})
			}

			break
		}
	}

	var checksum int
	for _, block := range blocks {
		if block.fileID == -1 {
			continue
		}

		for i := range block.length {
			checksum += (block.offset + i) * block.fileID
		}
	}

	return checksum
}

type block struct {
	fileID int
	offset int
	length int
}

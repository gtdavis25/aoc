package day17

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gtdavis25/aoc/internal/parse"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(context solver.Context) error {
	lines, err := context.InputLines()
	if err != nil {
		return err
	}

	var a, b, c int
	if _, err := fmt.Sscanf(lines[0], "Register A: %d", &a); err != nil {
		return fmt.Errorf("parsing register A: %w", err)
	}

	if _, err := fmt.Sscanf(lines[1], "Register B: %d", &b); err != nil {
		return fmt.Errorf("parsing register B: %w", err)
	}

	if _, err := fmt.Sscanf(lines[2], "Register C: %d", &c); err != nil {
		return fmt.Errorf("parsing register C: %w", err)
	}

	program, err := parse.IntSlice(lines[4][len("Program: "):], ",")
	if err != nil {
		return fmt.Errorf("parsing program: %w", err)
	}

	output, err := runProgram(a, b, c, program)
	if err != nil {
		return err
	}

	var part1 []string
	for _, n := range output {
		part1 = append(part1, fmt.Sprint(n))
	}

	context.SetPart1(strings.Join(part1, ","))
	queue := make([]int, 8)
	for i := range queue {
		queue[i] = i
	}

	for {
		a := queue[0]
		queue = queue[1:]
		output, err := runProgram(a, 0, 0, program)
		if err != nil {
			return err
		}

		if slices.Equal(output, program) {
			context.SetPart2(a)
			return nil
		}

		if slices.Equal(output, program[len(program)-len(output):]) {
			for i := range 8 {
				queue = append(queue, a<<3+i)
			}
		}
	}
}

func runProgram(a, b, c int, program []int) ([]int, error) {
	var output []int
	getComboOperand := func(code int) (int, error) {
		switch code {
		case 0, 1, 2, 3:
			return code, nil

		case 4:
			return a, nil

		case 5:
			return b, nil

		case 6:
			return c, nil

		default:
			return 0, fmt.Errorf("invalid operand: %d", code)
		}
	}

	for i := 0; i < len(program); {
		switch program[i] {
		case 0:
			operand, err := getComboOperand(program[i+1])
			if err != nil {
				return nil, fmt.Errorf("adv: %w", err)
			}

			a >>= operand

		case 1:
			b ^= program[i+1]

		case 2:
			operand, err := getComboOperand(program[i+1])
			if err != nil {
				return nil, fmt.Errorf("bst: %w", err)
			}

			b = operand % 8

		case 3:
			if a != 0 {
				i = program[i+1]
				continue
			}

		case 4:
			b ^= c

		case 5:
			operand, err := getComboOperand(program[i+1])
			if err != nil {
				return nil, fmt.Errorf("out: %w", err)
			}

			output = append(output, operand%8)

		case 6:
			operand, err := getComboOperand(program[i+1])
			if err != nil {
				return nil, fmt.Errorf("bdv: %w", err)
			}

			b = a >> operand

		case 7:
			operand, err := getComboOperand(program[i+1])
			if err != nil {
				return nil, fmt.Errorf("cdv: %w", err)
			}

			c = a >> operand

		default:
			return nil, fmt.Errorf("invalid opcode: %d", program[i])
		}

		i += 2
	}

	return output, nil
}

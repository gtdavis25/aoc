package day16

import (
	"container/heap"
	"fmt"

	"github.com/gtdavis25/aoc/internal/geom2d"
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

	part1, err := getMinimumScore(lines)
	if err != nil {
		return err
	}

	context.SetPart1(part1)
	return nil
}

func getMinimumScore(lines []string) (int, error) {
	start, err := getStartNode(lines)
	if err != nil {
		return 0, err
	}

	seen := make(map[state]struct{})
	q := queue{start}
	for len(q) > 0 {
		n := heap.Pop(&q).(node)
		if lines[n.pos.Y][n.pos.X] == 'E' {
			return n.score, nil
		}

		if _, ok := seen[n.state]; ok {
			continue
		}

		seen[n.state] = struct{}{}
		if lines[n.pos.Y+n.vel.Y][n.pos.X+n.vel.X] != '#' {
			heap.Push(&q, node{
				state: state{
					pos: n.pos.Add(n.vel),
					vel: n.vel,
				},
				score: n.score + 1,
			})
		}

		for _, next := range []node{
			{
				state: state{
					pos: n.pos,
					vel: geom2d.Point{
						X: -n.vel.Y,
						Y: n.vel.X,
					},
				},
				score: n.score + 1000,
			},
			{
				state: state{
					pos: n.pos,
					vel: geom2d.Point{
						X: n.vel.Y,
						Y: -n.vel.X,
					},
				},
				score: n.score + 1000,
			},
		} {
			if _, ok := seen[next.state]; ok {
				continue
			}

			heap.Push(&q, next)
		}
	}

	return 0, fmt.Errorf("no solutions")
}

type node struct {
	state
	score int
}

func getStartNode(lines []string) (node, error) {
	for y, line := range lines {
		for x := range len(line) {
			if line[x] == 'S' {
				return node{
					state: state{
						pos: geom2d.Point{
							X: x,
							Y: y,
						},
						vel: geom2d.Right(),
					},
					score: 0,
				}, nil
			}
		}
	}

	return node{}, fmt.Errorf("no start position")
}

type state struct {
	pos geom2d.Point
	vel geom2d.Point
}

type queue []node

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].score < q[j].score
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(val any) {
	*q = append(*q, val.(node))
}

func (q *queue) Pop() any {
	n := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return n
}

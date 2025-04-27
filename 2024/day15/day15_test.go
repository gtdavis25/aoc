package day15_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/gtdavis25/aoc/2024/day15"
)

func TestDoMoves(t *testing.T) {
	for i, tc := range []struct {
		lines []string
		moves string
		want  []string
	}{
		{
			lines: []string{
				"######",
				"#@OO.#",
				"######",
			},
			moves: ">",
			want: []string{
				"######",
				"#.@OO#",
				"######",
			},
		},
		{
			lines: []string{
				"######",
				"#@[].#",
				"######",
			},
			moves: ">",
			want: []string{
				"######",
				"#.@[]#",
				"######",
			},
		},
		{
			lines: []string{
				"######",
				"#....#",
				"#[][]#",
				"#.[].#",
				"#.@..#",
				"######",
			},
			moves: "^",
			want: []string{
				"######",
				"#[][]#",
				"#.[].#",
				"#.@..#",
				"#....#",
				"######",
			},
		},
	} {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			t.Parallel()
			rows := make([][]byte, len(tc.lines))
			for i := range rows {
				rows[i] = []byte(tc.lines[i])
			}

			moves, err := day15.GetMoves(tc.moves)
			if err != nil {
				t.Fatalf("getting moves: %v", err)
			}

			if err := day15.DoMoves(rows, moves); err != nil {
				t.Fatalf("doing moves: %v", err)
			}

			got := make([]string, len(rows))
			for i := range got {
				got[i] = string(rows[i])
			}

			if !slices.Equal(got, tc.want) {
				t.Fatalf(
					"\ngot:\n%v\n\nwant:\n%v",
					strings.Join(got, "\n"),
					strings.Join(tc.want, "\n"),
				)
			}
		})
	}
}

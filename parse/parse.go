package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func IntSlice(s, sep string) ([]int, error) {
	parts := strings.Split(s, sep)
	ints := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("parsing %q as int: %w", part, err)
		}

		ints[i] = n
	}

	return ints, nil
}

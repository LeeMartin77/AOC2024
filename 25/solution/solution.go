package solution

import (
	"strings"
)

func parse(data []byte) ([][]int, [][]int) {
	keys := [][]int{}
	locks := [][]int{}

	lock := true
	heap := []int{}

	for i, ln := range strings.Split(string(data), "\n") {
		if (i+1)%8 == 0 {
			continue
		}
		if (i+1)%8 == 1 {
			heap = make([]int, len(ln))
			lock = ln[0] == '#'
			continue
		}
		if (i+1)%8 == 7 {

			if lock {
				locks = append(locks, heap)
			} else {
				keys = append(keys, heap)
			}
			continue
		}
		for i, rn := range ln {
			if rn == '#' {
				heap[i] += 1
			}
		}
	}
	return keys, locks
}

func ComputeSolutionOne(data []byte) int64 {
	keys, locks := parse(data)
	size_limit := 5
	acc := int64(0)
	for _, key := range keys {
	outer:
		for _, lock := range locks {
			for i := range len(key) {
				if key[i]+lock[i] > size_limit {
					continue outer
				}
			}
			acc += 1
		}
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

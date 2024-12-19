package solution

import (
	"fmt"
	"strings"
	"sync"
)

func parseTowelsAndPatterns(data []byte) ([]string, []string) {
	towels := []string{}
	patterns := []string{}
	for i, ln := range strings.Split(string(data), "\n") {
		if i == 0 {
			towels = append(towels, strings.Split(ln, ", ")...)

		}
		if i > 1 {
			patterns = append(patterns, ln)
		}
	}
	return towels, patterns
}

func patternPossible(towels []string, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	for _, twl := range towels {
		if strings.Index(pattern, twl) == 0 {
			if patternPossible(towels, pattern[len(twl):]) {
				return true
			}
		}
	}
	return false
}

func ComputeSolutionOne(data []byte) int64 {
	towels, patterns := parseTowelsAndPatterns(data)
	acc := int64(0)
	wg := sync.WaitGroup{}
	for _, ptrn := range patterns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if patternPossible(towels, ptrn) {
				acc += 1
			}
			fmt.Println(acc)
		}()
	}
	wg.Wait()
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

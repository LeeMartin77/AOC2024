package solution

import (
	"fmt"
	"slices"
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

func patternPossible(pattern string, towels []string) bool {
	applicabletowels := []string{}
	for _, twl := range towels {
		if strings.Contains(pattern, twl) {
			applicabletowels = append(applicabletowels, twl)
		}
	}

	possibles := []string{}
	for _, atwl := range applicabletowels {
		possibles = GetNewPossibilities(pattern, atwl, possibles)

		if slices.ContainsFunc(possibles, func(poss string) bool { return !strings.Contains(poss, "x") }) {
			return true
		}
	}
	for _, atwl := range applicabletowels {
		possibles = GetNewPossibilities(pattern, atwl, possibles)

		if slices.ContainsFunc(possibles, func(poss string) bool { return !strings.Contains(poss, "x") }) {
			return true
		}
	}
	return false
}

func GetNewPossibilities(pattern string, atwl string, possibles []string) []string {
	sptrn := pattern
	blank := ""
	for range pattern {
		blank += "x"
	}
	idx := 0
	fits := []int{}
	offset := 0
	for idx != -1 {
		if len(sptrn) == 0 {
			break
		}
		idx = strings.Index(sptrn, atwl)
		if idx != -1 {
			fits = append(fits, idx+offset)
			if len(sptrn) > idx {
				offset += idx + len(atwl)
				sptrn = sptrn[idx+1:]
			} else {
				break
			}
		}
	}
	for _, ft := range fits {
		if ft >= len(blank) {
			continue
		}
		new_poss := blank[:ft] + atwl

		if len(new_poss) < len(blank) {
			new_poss += blank[len(new_poss):]
		}
		possibles = append(possibles, new_poss)
	}

	fit_bak := []int{}
	for _, ft := range fits {
		fit_bak = append([]int{ft}, fit_bak...)
	}

	for _, poss := range possibles {

		for _, ft := range fits {
			// check if fits on poss is all xs - if yes, replace

			if ft+len(atwl) < len(poss) && !strings.ContainsAny(poss[ft:ft+len(atwl)], "wubrg") {
				new_poss := poss[:ft] + atwl
				if len(poss) > len(new_poss) {
					new_poss += poss[len(new_poss):]
				}
				if !slices.Contains(possibles, new_poss) {
					possibles = append(possibles, new_poss)
				}
				poss = new_poss
			}
		}
	}

	for _, poss := range possibles {
		// go backwards too

		for _, ft := range fit_bak {
			// check if fits on poss is all xs - if yes, replace

			if ft+len(atwl) < len(poss) && !strings.ContainsAny(poss[ft:ft+len(atwl)], "wubrg") {
				new_poss := poss[:ft] + atwl
				if len(poss) > len(new_poss) {
					new_poss += poss[len(new_poss):]
				}
				if !slices.Contains(possibles, new_poss) {
					possibles = append(possibles, new_poss)
				}
				poss = new_poss
			}
		}
	}
	return possibles
}

func ComputeSolutionOne(data []byte) int64 {
	towels, patterns := parseTowelsAndPatterns(data)
	acc := int64(0)
	wg := sync.WaitGroup{}
	for _, ptrn := range patterns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if patternPossible(ptrn, towels) {
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

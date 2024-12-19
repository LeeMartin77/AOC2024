package solution

import (
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

type cache struct {
	patternResults map[string]*int64
	towels         []string
	rwmtx          sync.RWMutex
}

func (ch *cache) patternPossible(pattern string) int64 {
	ch.rwmtx.RLock()
	chres := ch.patternResults[pattern]
	ch.rwmtx.RUnlock()
	if chres != nil {
		return *chres
	}
	// none of the towels are present - seems to repeat forever otherwise

	if len(pattern) == 0 {
		return 1
	}
	acc := int64(0)
	for _, twl := range ch.towels {
		if strings.Index(pattern, twl) == 0 {
			res := ch.patternPossible(pattern[len(twl):])
			ch.rwmtx.Lock()
			ch.patternResults[pattern[len(twl):]] = &res
			ch.rwmtx.Unlock()
			acc += res
		}
	}
	return acc
}

func ComputeSolutionOne(data []byte) int64 {
	towels, patterns := parseTowelsAndPatterns(data)
	acc := int64(0)
	wg := sync.WaitGroup{}
	ch := cache{
		towels:         towels,
		patternResults: map[string]*int64{},
		rwmtx:          sync.RWMutex{},
	}
	for _, ptrn := range patterns {
		wg.Add(1)
		go func() {
			if ch.patternPossible(ptrn) > 0 {
				acc += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	towels, patterns := parseTowelsAndPatterns(data)
	acc := int64(0)
	wg := sync.WaitGroup{}
	ch := cache{
		towels:         towels,
		patternResults: map[string]*int64{},
		rwmtx:          sync.RWMutex{},
	}
	for _, ptrn := range patterns {
		wg.Add(1)
		go func() {
			acc += ch.patternPossible(ptrn)
			wg.Done()
		}()
	}
	wg.Wait()
	return acc
}

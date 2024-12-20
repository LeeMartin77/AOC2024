package solution

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func parseStones(data []byte) []int64 {
	stnstrs := strings.Split(string(data), " ")
	res := []int64{}
	for _, stn := range stnstrs {
		pr, _ := strconv.ParseInt(stn, 10, 64)
		res = append(res, pr)
	}
	return res
}

func blink(stns []int64) []int64 {
	ret := []int64{}
	for _, stn := range stns {
		if stn == 0 {
			ret = append(ret, 1)
			continue
		}
		strstn := fmt.Sprintf("%d", stn)
		if len(strstn)%2 == 0 {
			first, second := strstn[:len(strstn)/2], strstn[len(strstn)/2:]
			f, _ := strconv.ParseInt(first, 10, 64)
			s, _ := strconv.ParseInt(second, 10, 64)
			ret = append(ret, f)
			ret = append(ret, s)
			continue
		}
		ret = append(ret, stn*2024)
	}
	return ret
}

// recurrs down until all iterations are done, allowing everything to be added up
func blinkRecursive(stn int64, iteration int, iteration_limit int) int64 {
	if iteration == iteration_limit {
		return 1
	}
	acc := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if stn == 0 {
			acc = blinkRecursive(1, iteration+1, iteration_limit)
			return
		}
		strstn := fmt.Sprintf("%d", stn)
		if len(strstn)%2 == 0 {
			first, second := strstn[:len(strstn)/2], strstn[len(strstn)/2:]
			f, _ := strconv.ParseInt(first, 10, 64)
			s, _ := strconv.ParseInt(second, 10, 64)
			acc = blinkRecursive(f, iteration+1, iteration_limit) + blinkRecursive(s, iteration+1, iteration_limit)
			return
		}
		acc = blinkRecursive(stn*2024, iteration+1, iteration_limit)
	}()
	wg.Wait()
	return acc
}

func blinkMap(inp map[int64]int64) map[int64]int64 {
	outmp := map[int64]int64{}
	for stn, cnt := range inp {
		outstns := blink([]int64{stn})
		for _, stn := range outstns {
			outmp[stn] += cnt
		}
	}
	return outmp
}

func computeIterations(data []byte, num int) int64 {
	stns := parseStones(data)
	acc := int64(0)

	stonemap := map[int64]int64{}
	for _, stn := range stns {
		stonemap[stn] += 1
	}

	for range num {
		stonemap = blinkMap(stonemap)
	}
	for _, cnt := range stonemap {
		acc += cnt
	}
	return acc
}

func ComputeSolutionOne(data []byte) int64 {
	return computeIterations(data, 25)
}

func ComputeSolutionTwo(data []byte) int64 {
	return computeIterations(data, 75)

}

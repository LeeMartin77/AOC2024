package solution

import (
	"fmt"
	"strconv"
	"strings"
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

func ComputeSolutionOne(data []byte) int64 {
	stns := parseStones(data)
	for range 25 {
		stns = blink(stns)
	}
	return int64(len(stns))
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

package solution

import (
	"math"
	"strconv"
	"strings"
)

func parseData(data []byte) [][]int64 {
	strdta := string(data)
	lines := strings.Split(strdta, "\n")
	res := [][]int64{}
	for _, ln := range lines {
		prts := strings.Split(ln, " ")
		bit := []int64{}
		for _, prtstr := range prts {
			prt, _ := strconv.ParseInt(prtstr, 10, 64)
			bit = append(bit, prt)
		}
		res = append(res, bit)
	}
	return res
}

func ComputeSolutionOne(data []byte) int64 {
	parsed := parseData(data)
	accumulator := int64(0)
	for _, p := range parsed {
		if isSafe(p) {
			accumulator += 1
		}
	}
	return accumulator
}

// I'm sure there's a smart recursive function for this but I'm lazy
func isSafe(data []int64) bool {
	last := data[0]
	last_gradient := gradient(data[0], data[1])
	for _, nxt := range data[1:] {
		grdnt := gradient(last, nxt)
		if last_gradient != grdnt {
			return false
		}
		if grdnt == 0 {
			return false
		}
		if math.Abs(float64(nxt-last)) > 3 {
			return false
		}
		last = nxt
		last_gradient = grdnt
	}
	return true
}

func gradient(a, b int64) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("not implemented")
}

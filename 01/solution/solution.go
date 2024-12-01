package solution

import (
	"slices"
	"strconv"
	"strings"
)

type row struct {
	LineNumber int
	Left       int64
	Right      int64
}

func parseData(data []byte) []row {
	strdta := string(data)
	lines := strings.Split(strdta, "\n")
	res := []row{}
	for lnn, ln := range lines {
		prts := strings.Split(ln, "   ")
		l, _ := strconv.ParseInt(prts[0], 10, 64)
		r, _ := strconv.ParseInt(prts[1], 10, 64)
		res = append(res, row{
			LineNumber: lnn,
			Left:       l,
			Right:      r,
		})
	}
	return res
}

func ComputeSolutionOne(data []byte) int64 {
	parsed := parseData(data)
	leftList := []int64{}
	rightList := []int64{}
	for _, vl := range parsed {
		leftList = append(leftList, vl.Left)
		rightList = append(rightList, vl.Right)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)
	accumulator := int64(0)
	for i, l := range leftList {
		r := rightList[i]
		if l == r {
			continue //technically this is "add zero"
		}
		if l > r {
			new := l - r
			accumulator += new
		} else {
			new := r - l
			accumulator += new
		}
	}
	return accumulator
}

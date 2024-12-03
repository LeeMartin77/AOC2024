package solution

import (
	"regexp"
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	str := string(data)
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}
	acc := int64(0)
	for _, mulop := range r.FindAllString(str, -1) {
		digits := strings.Split(mulop, "(")[1]
		digits = digits[:len(digits)-1]
		digits_split := strings.Split(digits, ",")
		a, _ := strconv.ParseInt(digits_split[0], 10, 64)
		b, _ := strconv.ParseInt(digits_split[1], 10, 64)
		acc += (a * b)
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

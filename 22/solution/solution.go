package solution

import (
	"strconv"
	"strings"
)

// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.

var prune_modulo = int64(16777216)

func CalculateNthSecretNumber(init int64, iterations int64) int64 {
	for range iterations {
		s1 := init * 64
		init = init ^ s1
		init = init % prune_modulo

		s2 := init / 32
		init = init ^ s2
		init = init % prune_modulo

		s3 := init * 2048
		init = init ^ s3
		init = init % prune_modulo
	}
	return init
}

func parse(data []byte) []int64 {
	ret := []int64{}
	for _, vl := range strings.Split(string(data), "\n") {
		v, _ := strconv.ParseInt(vl, 10, 64)
		ret = append(ret, v)
	}
	return ret
}

func ComputeSolutionOne(data []byte) int64 {
	initial_prices := parse(data)
	acc := int64(0)
	for _, vl := range initial_prices {
		acc += CalculateNthSecretNumber(vl, 2000)
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

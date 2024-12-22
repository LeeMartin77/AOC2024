package solution

import (
	"fmt"
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

func CalculateNSecretNumberPrices(init int64, iterations int64) []int8 {
	ones := []int8{int8(init % 10)}
	for range iterations - 1 {
		s1 := init * 64
		init = init ^ s1
		init = init % prune_modulo

		s2 := init / 32
		init = init ^ s2
		init = init % prune_modulo

		s3 := init * 2048
		init = init ^ s3
		init = init % prune_modulo

		ones = append(ones, int8(init%10))
	}
	return ones
}

func HashSequences(inp []int8) map[string]int8 {
	dlts := []int8{}
	mp := map[string]int8{}
	for i := range inp {
		if i == 0 {
			continue
		}
		if len(dlts) > 3 {
			dlts = dlts[1:]
		}
		dlts = append(dlts, inp[i]-inp[i-1])
		if len(dlts) > 3 {
			delta_hash := fmt.Sprintf("%d,%d,%d,%d", dlts[0], dlts[1], dlts[2], dlts[3])
			if mp[delta_hash] == 0 {
				mp[delta_hash] = inp[i]
			}
		}
	}
	return mp
}

func CalculateNSecretNumberPriceHashes(init int64, iterations int64) map[string]int8 {
	clc := CalculateNSecretNumberPrices(init, iterations)
	return HashSequences(clc)
}

func GetSequencesForDigit(mp map[int8]map[string]bool, dgd int8) []string {
	if mp[dgd] != nil {
		ret := []string{}
		for k := range mp[dgd] {
			ret = append(ret, k)
		}
		return ret
	}
	return []string{}
}

func ComputeSolutionOne(data []byte) int64 {
	initial_secrets := parse(data)
	acc := int64(0)
	for _, vl := range initial_secrets {
		acc += CalculateNthSecretNumber(vl, 2000)
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	initial_secrets := parse(data)
	hshs := []map[string]int8{}
	for _, vl := range initial_secrets {
		hshs = append(hshs, CalculateNSecretNumberPriceHashes(vl, 2000))
	}

	totals := map[string]int64{}
	for _, hsh := range hshs {
		for str, t := range hsh {
			totals[str] += int64(t)
		}
	}
	max_hash := ""
	max_hash_num := int64(0)
	for k := range totals {
		if totals[k] > int64(max_hash_num) {
			max_hash = k
			max_hash_num = totals[k]
		}
	}
	fmt.Println(max_hash)
	return max_hash_num
}

func parse(data []byte) []int64 {
	ret := []int64{}
	for _, vl := range strings.Split(string(data), "\n") {
		v, _ := strconv.ParseInt(vl, 10, 64)
		ret = append(ret, v)
	}
	return ret
}

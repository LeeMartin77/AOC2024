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

func HashSequences(inp []int8) map[int8]map[string]bool {
	dlts := []int8{inp[0]}
	mp := map[int8]map[string]bool{}
	for i, in := range inp[1:] {
		if len(dlts) > 3 {
			dlts = dlts[1:]
		}
		dlts = append(dlts, in-inp[i])
		if len(dlts) > 3 {
			if mp[in] == nil {
				mp[in] = map[string]bool{
					fmt.Sprintf("%v", dlts): true,
				}
			} else {
				mp[in][fmt.Sprintf("%v", dlts)] = true
			}
		}
	}
	return mp
}

func CalculateNSecretNumberPriceHashes(init int64, iterations int64) map[int8]map[string]bool {
	clc := CalculateNSecretNumberPrices(init, iterations)
	return HashSequences(clc)
}

func GetBestForSequence(seq string, mp map[int8]map[string]bool) int8 {
	top := int8(9)
	for i := range int8(9) {
		if mp[top-i][seq] {
			return top - i
		}
	}
	return 0
}

func GetBestSequences(mp map[int8]map[string]bool) (int8, []string) {
	top := int8(9)
	for i := range int8(9) {
		if mp[top-i] != nil {
			ret := []string{}
			for k := range mp[top-i] {
				ret = append(ret, k)
			}
			return top - i, ret
		}
	}
	return 0, []string{}
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
	hshs := []map[int8]map[string]bool{}
	for _, vl := range initial_secrets {
		hshs = append(hshs, CalculateNSecretNumberPriceHashes(vl, 2000))
	}
	max_nanas := int64(0)
	max_seq := ""
	top := int8(9)
	for i := range int8(9) {
		seq := map[string]bool{}
		for _, hsh := range hshs {
			seqs := GetSequencesForDigit(hsh, top-i)
			for _, s := range seqs {
				seq[s] = true
			}
		}
		for seq := range seq {
			go func() {
				acc := int64(0)
				for _, hsh := range hshs {
					acc += int64(GetBestForSequence(seq, hsh))
				}
				if acc > max_nanas {
					max_nanas = acc
					max_seq = seq
				}
			}()
		}
	}
	fmt.Println(max_seq)
	return max_nanas
}

func parse(data []byte) []int64 {
	ret := []int64{}
	for _, vl := range strings.Split(string(data), "\n") {
		v, _ := strconv.ParseInt(vl, 10, 64)
		ret = append(ret, v)
	}
	return ret
}

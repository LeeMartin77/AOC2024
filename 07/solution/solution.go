package solution

import (
	"fmt"
	"strconv"
	"strings"
)

type possibleEquation struct {
	Solution int64
	Inputs   []int64
}

func parseData(data []byte) []possibleEquation {
	str := string(data)
	acc := []possibleEquation{}
	for _, eq := range strings.Split(str, "\n") {
		prts := strings.Split(eq, ": ")
		sol, _ := strconv.ParseInt(prts[0], 10, 64)

		eqa := possibleEquation{
			Solution: sol,
		}
		for _, inp := range strings.Split(prts[1], " ") {
			in, _ := strconv.ParseInt(inp, 10, 64)
			eqa.Inputs = append(eqa.Inputs, in)
		}
		acc = append(acc, eqa)
	}
	return acc
}

func (eq possibleEquation) OperatorsMakeSolution(opset []rune) bool {
	acc := int64(eq.Inputs[0])
	for i, op := range opset {
		switch op {
		case '+':
			acc += eq.Inputs[i+1]
		case '*':
			acc *= eq.Inputs[i+1]
		}
	}
	return acc == eq.Solution
}

func (eq possibleEquation) FindSolutionAddMul() ([]rune, error) {
	var possible_op_sets [][]rune
	n := len(eq.Inputs) - 1

	totalCombinations := 1 << n

	for i := 0; i < totalCombinations; i++ {
		var variant []rune
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				variant = append(variant, '*')
			} else {
				variant = append(variant, '+')
			}
		}
		possible_op_sets = append(possible_op_sets, variant)
	}

	for _, pop := range possible_op_sets {
		if eq.OperatorsMakeSolution(pop) {
			return pop, nil
		}
	}

	return nil, fmt.Errorf("no solution")
}

func ComputeSolutionOne(data []byte) int64 {
	eqs := parseData(data)
	acc := int64(0)
	for _, eq := range eqs {
		_, err := eq.FindSolutionAddMul()
		if err == nil {
			acc += eq.Solution
		}
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

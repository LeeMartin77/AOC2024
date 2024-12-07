package solution

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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
		case '|':
			vl, _ := strconv.ParseInt(fmt.Sprintf("%d%d", acc, eq.Inputs[i+1]), 10, 64)
			acc = vl
		}
	}
	return acc == eq.Solution
}

func (eq possibleEquation) FindSolutionOperators(ops []rune) ([]rune, error) {
	var possible_op_sets [][]rune
	n := len(eq.Inputs) - 1
	numreps := len(ops)

	totalCombinations := 1
	for i := 0; i < n; i++ {
		totalCombinations *= numreps
	}

	for i := 0; i < totalCombinations; i++ {
		var variant []rune
		temp := i
		for j := 0; j < n; j++ {
			// Determine the index of the replacement for this position
			replacementIndex := temp % numreps
			variant = append(variant, ops[replacementIndex])
			temp /= numreps
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
		_, err := eq.FindSolutionOperators([]rune{'+', '*'})
		if err == nil {
			acc += eq.Solution
		}
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	eqs := parseData(data)
	acc := int64(0)
	wg := sync.WaitGroup{}
	for _, eq := range eqs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := eq.FindSolutionOperators([]rune{'+', '*', '|'})
			if err == nil {
				acc += eq.Solution
			}
		}()
	}
	wg.Wait()
	return acc
}

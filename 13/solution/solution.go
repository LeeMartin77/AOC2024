package solution

import (
	"strconv"
	"strings"
)

type clawMachine struct {
	ButtonA       []int64
	ButtonB       []int64
	PrizeLocation []int64
}

func parseMachines(data []byte) []clawMachine {
	res := []clawMachine{}
	str := string(data)
	mchn := clawMachine{}
	for _, ln := range strings.Split(str, "\n") {
		if ln == "" {
			res = append(res, mchn)
			mchn = clawMachine{}
			continue
		}
		prts := strings.Split(ln, ",")
		if strings.Contains(ln, "Button A:") {
			xprts := strings.Split(prts[0], "+")
			x, _ := strconv.ParseInt(xprts[1], 10, 64)
			yprts := strings.Split(prts[1], "+")
			y, _ := strconv.ParseInt(yprts[1], 10, 64)
			mchn.ButtonA = []int64{x, y}
		}
		if strings.Contains(ln, "Button B:") {
			xprts := strings.Split(prts[0], "+")
			x, _ := strconv.ParseInt(xprts[1], 10, 64)
			yprts := strings.Split(prts[1], "+")
			y, _ := strconv.ParseInt(yprts[1], 10, 64)
			mchn.ButtonB = []int64{x, y}
		}
		if strings.Contains(ln, "Prize:") {
			xprts := strings.Split(prts[0], "=")
			x, _ := strconv.ParseInt(xprts[1], 10, 64)
			yprts := strings.Split(prts[1], "=")
			y, _ := strconv.ParseInt(yprts[1], 10, 64)
			mchn.PrizeLocation = []int64{x, y}
		}
	}
	return res
}

func (mchn clawMachine) TokensForVictory() int64 {

	a_tokens := int64(3)
	b_tokens := int64(1)

	solution_map := map[int64]map[int64]bool{}

	for a := range int64(101) {
		x := a * mchn.ButtonA[0]
		y := a * mchn.ButtonA[1]
		for b := range int64(101) {
			xb := x
			yb := y
			xb += b * mchn.ButtonB[0]
			yb += b * mchn.ButtonB[1]
			if mchn.PrizeLocation[0] == xb && mchn.PrizeLocation[1] == yb {
				if solution_map[a] == nil {
					solution_map[a] = map[int64]bool{}
				}
				solution_map[a][b] = true
			}
		}
	}

	cheapest_solution := int64(0)
	for a, bs := range solution_map {
		for b, ys := range bs {
			if ys {
				sol := (a * a_tokens) + (b * b_tokens)
				if cheapest_solution == 0 || sol < cheapest_solution {
					cheapest_solution = sol
				}
			}
		}
	}
	return cheapest_solution
}

func ComputeSolutionOne(data []byte) int64 {
	mchns := parseMachines(data)
	acc := int64(0)
	for _, mchn := range mchns {
		acc += mchn.TokensForVictory()
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

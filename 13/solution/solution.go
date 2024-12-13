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

func parseMachines(data []byte, prizeOffset int64) []clawMachine {
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
			mchn.PrizeLocation = []int64{x + prizeOffset, y + prizeOffset}
		}
	}
	return res
}

func (mchn clawMachine) TokensForVictory() int64 {
	// I won't lie, I stole this, because I can't do advanced math
	// Do some research into linear algebra and matrix maths
	ax, ay := mchn.ButtonA[0], mchn.ButtonA[1]
	bx, by := mchn.ButtonB[0], mchn.ButtonB[1]
	px, py := mchn.PrizeLocation[0], mchn.PrizeLocation[1]

	aCoeff := ax*by - ay*bx
	rhs := px*by - py*bx
	if rhs%aCoeff == 0 {
		a := rhs / aCoeff
		b := (px - ax*a) / bx
		return a*3 + b
	}
	return 0
}

func ComputeSolutionOne(data []byte) int64 {
	mchns := parseMachines(data, 0)
	acc := int64(0)
	for _, mchn := range mchns {
		acc += mchn.TokensForVictory()
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	mchns := parseMachines(data, 10000000000000)
	acc := int64(0)
	for _, mchn := range mchns {
		acc += mchn.TokensForVictory()
	}
	return acc
}

package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/19/solution"
	"github.com/stretchr/testify/assert"
)

func TestPossibles(t *testing.T) {

	poss := solution.GetNewPossibilities("brwrr", "r", []string{"wubrg", "wxbrg"})
	assert.ElementsMatch(t, []string{
		"wubrg",
		"wxbrg",
		"wrbrg",
		"xrxxx",
		"xxxxr",
		"xxxrx",

		"xrxrx",
		"xxxrr",
		"xrxxr",

		"xrxrr",
	}, poss)

	poss2 := solution.GetNewPossibilities("brwrr", "rw", []string{"wubrg", "wxbrg"})
	assert.ElementsMatch(t, []string{
		"wubrg",
		"wxbrg",
		"xrwxx",
	}, poss2)
}

func TestSingle(t *testing.T) {
	teststring := `r, wr, b, g, bwu, rb, gb, br

bwurrg`

	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(1), res)
}

func TestPhaseOne(t *testing.T) {
	teststring := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(6), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

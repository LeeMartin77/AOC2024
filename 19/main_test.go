package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/19/solution"
	"github.com/stretchr/testify/assert"
)

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

func TestPhaseTwo(t *testing.T) {
	teststring := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(16), res)
}

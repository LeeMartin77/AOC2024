package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/03/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(161), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(48), res)
}

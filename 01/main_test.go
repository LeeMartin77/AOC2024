package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/01/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `3   4
4   3
2   5
1   3
3   9
3   3`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(11), res)
}

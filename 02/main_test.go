package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/02/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(2), res)
}

// func TestPHaseTwo(t *testing.T) {
// 	teststring := `3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3`
// 	res := solution.ComputeSolutionTwo([]byte(teststring))
// 	assert.Equal(t, int64(31), res)
// }

package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/18/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(22), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

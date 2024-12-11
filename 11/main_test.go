package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/11/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `125 17`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(55312), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

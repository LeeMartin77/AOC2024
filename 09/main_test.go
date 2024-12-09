package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/09/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `2333133121414131402`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(1928), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `2333133121414131402`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(2858), res)
}

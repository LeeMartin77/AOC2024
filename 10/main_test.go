package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/10/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(36), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(81), res)
}

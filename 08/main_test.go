package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/08/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(14), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(34), res)
}

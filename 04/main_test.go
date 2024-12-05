package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/04/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(18), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(9), res)
}

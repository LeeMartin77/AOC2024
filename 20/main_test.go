package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/20/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	res := solution.ComputeSolutionOne([]byte(teststring), 2)
	assert.Equal(t, int64(44), res)
}
func TestPhaseOne_Thresh(t *testing.T) {
	teststring := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	res := solution.ComputeSolutionOne([]byte(teststring), 3)
	assert.Equal(t, int64(30), res)
}

func TestPhaseOne_Thresh2(t *testing.T) {
	teststring := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	res := solution.ComputeSolutionOne([]byte(teststring), 4)
	assert.Equal(t, int64(30), res)
}
func TestPhaseTwo(t *testing.T) {
	//	teststring := `###############
	//
	// #...#...#.....#
	// #.#.#.#.#.###.#
	// #S#...#.#.#...#
	// #######.#.#.###
	// #######.#.#...#
	// #######.#.###.#
	// ###..E#...#...#
	// ###.#######.###
	// #...###...#...#
	// #.#####.#.###.#
	// #.#...#.#.#...#
	// #.#.#.#.#.#.###
	// #...#...#...###
	// ###############`
	//
	//	res := solution.ComputeSolutionTwo([]byte(teststring), 50)
	//	assert.Equal(t, int64(
	//		32+31+29+39+25+23+20+19+12+14+12+22+4+3,
	//	), res)
}

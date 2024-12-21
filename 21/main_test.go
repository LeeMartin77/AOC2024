package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/21/solution"
	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	assert.Equal(t, int64(29), solution.GetNumber("029A"))
	assert.Equal(t, int64(456), solution.GetNumber("456A"))
}

func TestTypeNumberPad(t *testing.T) {
	assert.Equal(t, "<A^A>^^AvvvA", solution.TypeNumberPad("029A"))
}

func TestTypeCommandPad(t *testing.T) {
	cc := solution.ComCache{}
	assert.Equal(t, len("v<<A>>^A<A>AvA<^AA>A<vAAA>^A"), len(cc.TypeCommmandPad("<A^A>^^AvvvA")))
	assert.Equal(t, len("<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"), len(cc.TypeCommmandPad("v<<A>>^A<A>AvA<^AA>A<vAAA>^A")))
}

func TestSingleRobotFirst(t *testing.T) {
	assert.Equal(t, int64(68*29), solution.GoThroughRobotsAndGetComplexity("029A", 2))
	assert.Equal(t, int64(60*980), solution.GoThroughRobotsAndGetComplexity("980A", 2))
	assert.Equal(t, int64(68*179), solution.GoThroughRobotsAndGetComplexity("179A", 2))
	assert.Equal(t, int64(64*456), solution.GoThroughRobotsAndGetComplexity("456A", 2))
	assert.Equal(t, int64(64*379), solution.GoThroughRobotsAndGetComplexity("379A", 2))
}

func TestPhaseOne(t *testing.T) {
	teststring := `029A
980A
179A
456A
379A`
	res := solution.ComputeSolutionOne([]byte(teststring))
	scr := int64(0)
	for _, i := range []int64{68 * 29, 60 * 980, 68 * 179, 64 * 456, 64 * 379} {
		scr += i
	}
	assert.Equal(t, int64(scr), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `029A
980A
179A
456A
379A`
	solution.ComputeSolutionTwo([]byte(teststring))
}

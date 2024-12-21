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
	assert.Equal(t, "v<<A>>^A<A>AvA<^AA>A<vAAA>^A", solution.TypeCommmandPad("<A^A>^^AvvvA", false))
	assert.Equal(t, "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A", solution.TypeCommmandPad("v<<A>>^A<A>AvA<^AA>A<vAAA>^A", true))
}

func TestSingleRobotFirst(t *testing.T) {
	assert.Equal(t, int64(68*29), solution.GoThroughRobotsAndGetComplexity("029A"))
	assert.Equal(t, int64(60*980), solution.GoThroughRobotsAndGetComplexity("980A"))
	assert.Equal(t, int64(68*179), solution.GoThroughRobotsAndGetComplexity("179A"))
	// this one is wrong. It likely has something to do with a more efficient command pattern
	// and it is making me want to die.
	assert.Equal(t, int64(64*456), solution.GoThroughRobotsAndGetComplexity("456A"))
	assert.Equal(t, int64(64*379), solution.GoThroughRobotsAndGetComplexity("379A"))

}

func xTestPhaseOne(t *testing.T) {

	// 029A

	// <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
	// v<<A>>^A<A>AvA<^AA>A<vAAA>^A
	// <A^A>^^AvvvA
	// 029A

	// 029A: <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
	// 980A: <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
	// 179A: <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
	// 456A: <v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A
	// 379A: <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A

	// score is length * numeric bit

	// 68 * 29, 60 * 980, 68 * 179, 64 * 456, 64 * 379

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

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

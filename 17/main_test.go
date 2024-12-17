package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/17/solution"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	teststring := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	cmp, prg := solution.ParseInput([]byte(teststring))
	assert.Equal(t, int64(729), cmp.RegisterA)
	assert.Equal(t, int64(0), cmp.RegisterB)
	assert.Equal(t, int64(0), cmp.RegisterC)
	assert.Equal(t, []int8{0, 1, 5, 4, 3, 0}, prg)
}

func TestPhaseOne(t *testing.T) {
	teststring := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, "4635635210", res)
}

func TestPhaseOneA(t *testing.T) {
	teststring := `Register A: 0
Register B: 0
Register C: 9

Program: 2,6`
	cmp, prg := solution.ParseInput([]byte(teststring))
	cmp.RunProgram(prg)
	assert.Equal(t, int64(1), cmp.RegisterB)
}

func TestPhaseOneB(t *testing.T) {
	teststring := `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`
	cmp, prg := solution.ParseInput([]byte(teststring))
	cmp.RunProgram(prg)
	assert.Equal(t, "012", cmp.PrintOutput())
}

func TestPhaseOneC(t *testing.T) {
	teststring := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	cmp, prg := solution.ParseInput([]byte(teststring))
	cmp.RunProgram(prg)
	assert.Equal(t, "42567777310", cmp.PrintOutput())
	assert.Equal(t, int64(0), cmp.RegisterA)
}

func TestPhaseOneD(t *testing.T) {
	teststring := `Register A: 0
Register B: 29
Register C: 0

Program: 1,7`
	cmp, prg := solution.ParseInput([]byte(teststring))
	cmp.RunProgram(prg)
	assert.Equal(t, int64(26), cmp.RegisterB)
}

func TestPhaseOneE(t *testing.T) {
	teststring := `Register A: 0
Register B: 2024
Register C: 43690

Program: 4,0`
	cmp, prg := solution.ParseInput([]byte(teststring))
	cmp.RunProgram(prg)
	assert.Equal(t, int64(44354), cmp.RegisterB)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

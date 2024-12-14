package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/14/solution"
	"github.com/stretchr/testify/assert"
)

func TestIsolated(t *testing.T) {
	// rbts, _, _ := solution.ParseRobots([]byte("p=2,4 v=2,-3"))
	// maxX := int64(10)
	// maxY := int64(6)

	// solution.DebugPositions(rbts, maxX, maxY)
	// for _, rbt := range rbts {
	// 	rbt.MoveTicks(1, maxX, maxY)
	// }
	// solution.DebugPositions(rbts, maxX, maxY)
	// for _, rbt := range rbts {
	// 	rbt.MoveTicks(1, maxX, maxY)
	// }
	// solution.DebugPositions(rbts, maxX, maxY)
	// for _, rbt := range rbts {
	// 	rbt.MoveTicks(1, maxX, maxY)
	// }
	// solution.DebugPositions(rbts, maxX, maxY)
	// for _, rbt := range rbts {
	// 	rbt.MoveTicks(1, maxX, maxY)
	// }
	// solution.DebugPositions(rbts, maxX, maxY)
	// for _, rbt := range rbts {
	// 	rbt.MoveTicks(1, maxX, maxY)
	// }
	// solution.DebugPositions(rbts, maxX, maxY)
}

func TestPhaseOne(t *testing.T) {
	teststring := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(12), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`
	solution.ComputeSolutionTwo([]byte(teststring))
}

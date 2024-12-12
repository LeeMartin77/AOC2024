package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/12/solution"
	"github.com/stretchr/testify/assert"
)

func TestParseRegions(t *testing.T) {
	teststring := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	res := solution.ParseRegions([]byte(teststring))
	assert.Equal(t, 11, len(res))

	for _, rg := range res {
		if rg.Rune == 'R' {
			assert.Equal(t, int64(12), rg.GetArea())
			assert.Equal(t, int64(18), rg.GetPerimeter())
			assert.Equal(t, int64(216), rg.GetPrice())
		}
		if rg.Rune == 'F' {
			assert.Equal(t, int64(10), rg.GetArea())
			assert.Equal(t, int64(18), rg.GetPerimeter())
			assert.Equal(t, int64(180), rg.GetPrice())
		}
		if rg.Rune == 'M' {
			assert.Equal(t, int64(5), rg.GetArea())
			assert.Equal(t, int64(12), rg.GetPerimeter())
			assert.Equal(t, int64(60), rg.GetPrice())
		}
	}
}

func TestPhaseOne(t *testing.T) {
	teststring := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(1930), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

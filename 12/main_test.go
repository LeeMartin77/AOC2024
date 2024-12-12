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

			assert.Equal(t, int64(10), rg.GetSides())
			assert.Equal(t, int64(120), rg.GetBulkPrice())
		}
		if rg.Rune == 'F' {
			assert.Equal(t, int64(10), rg.GetArea())
			assert.Equal(t, int64(18), rg.GetPerimeter())
			assert.Equal(t, int64(180), rg.GetPrice())

			assert.Equal(t, int64(12), rg.GetSides())
			assert.Equal(t, int64(120), rg.GetBulkPrice())
		}
		if rg.Rune == 'M' {
			assert.Equal(t, int64(5), rg.GetArea())
			assert.Equal(t, int64(12), rg.GetPerimeter())
			assert.Equal(t, int64(60), rg.GetPrice())

			assert.Equal(t, int64(6), rg.GetSides())
			assert.Equal(t, int64(30), rg.GetBulkPrice())
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

func TestPhaseTwo(t *testing.T) {
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
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(1206), res)
}

func TestPhaseTwo_One(t *testing.T) {
	teststring := `AAAA
BBCD
BBCC
EEEC`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(80), res)
}

func TestPhaseTwo_E(t *testing.T) {
	teststring := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(236), res)
}

func TestPhaseTwo_Mobius(t *testing.T) {
	teststring := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(368), res)
}

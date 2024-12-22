package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/22/solution"
	"github.com/stretchr/testify/assert"
)

func TestCalculateSingleSecretNumber(t *testing.T) {
	assert.Equal(t, int64(15887950), solution.CalculateNthSecretNumber(123, 1))
	assert.Equal(t, int64(16495136), solution.CalculateNthSecretNumber(123, 2))
	assert.Equal(t, int64(527345), solution.CalculateNthSecretNumber(123, 3))
	assert.Equal(t, int64(704524), solution.CalculateNthSecretNumber(123, 4))
	assert.Equal(t, int64(1553684), solution.CalculateNthSecretNumber(123, 5))
	assert.Equal(t, int64(12683156), solution.CalculateNthSecretNumber(123, 6))
	assert.Equal(t, int64(11100544), solution.CalculateNthSecretNumber(123, 7))
	assert.Equal(t, int64(12249484), solution.CalculateNthSecretNumber(123, 8))
	assert.Equal(t, int64(7753432), solution.CalculateNthSecretNumber(123, 9))
	assert.Equal(t, int64(5908254), solution.CalculateNthSecretNumber(123, 10))
}

func TestCalculateSingleSecretAndOnes(t *testing.T) {
	ones := solution.CalculateNSecretNumberPrices(123, 10)
	assert.Equal(t, []int8{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}, ones)
}

func TestPhaseOne(t *testing.T) {
	teststring := `1
10
100
2024`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(37327623), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `1
2
3
2024`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(23), res)
}

package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/23/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(7), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}

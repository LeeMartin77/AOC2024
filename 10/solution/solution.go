package solution

import (
	"slices"
	"strconv"
	"strings"
)

type trailMap struct {
	FullMap    map[int]map[int]int64
	MaxX       int
	MaxY       int
	TrailHeads [][]int
}

func parseMap(data []byte) trailMap {
	amp := trailMap{
		FullMap:    map[int]map[int]int64{},
		TrailHeads: [][]int{},
	}
	lns := strings.Split(string(data), "\n")
	amp.MaxY = len(lns) - 1
	for y, ln := range lns {

		if y == 0 {
			for i := range len(ln) {
				amp.FullMap[i] = map[int]int64{}
			}
			amp.MaxX = len(ln) - 1
		}
		for x, rn := range []rune(ln) {
			pin, _ := strconv.ParseInt(string(rn), 10, 64)
			amp.FullMap[x][y] = pin
			if pin == 0 {
				amp.TrailHeads = append(amp.TrailHeads, []int{x, y})
			}
		}

	}
	return amp
}

func recurTrailFinishes(curpos []int, tm map[int]map[int]int64) [][]int {
	curscr := tm[curpos[0]][curpos[1]]
	if curscr == 9 {
		return [][]int{curpos}
	}
	acc := [][]int{}
	if tm[curpos[0]+1][curpos[1]] == curscr+1 {
		acc = append(acc, recurTrailFinishes([]int{curpos[0] + 1, curpos[1]}, tm)...)
	}

	if tm[curpos[0]-1][curpos[1]] == curscr+1 {
		acc = append(acc, recurTrailFinishes([]int{curpos[0] - 1, curpos[1]}, tm)...)
	}

	if tm[curpos[0]][curpos[1]+1] == curscr+1 {
		acc = append(acc, recurTrailFinishes([]int{curpos[0], curpos[1] + 1}, tm)...)
	}

	if tm[curpos[0]][curpos[1]-1] == curscr+1 {
		acc = append(acc, recurTrailFinishes([]int{curpos[0], curpos[1] - 1}, tm)...)
	}
	return acc
}

func ComputeSolutionOne(data []byte) int64 {
	tm := parseMap(data)

	acc := int64(0)
	for _, th := range tm.TrailHeads {

		src := recurTrailFinishes(th, tm.FullMap)
		ddpd := [][]int{}
		for _, f := range src {
			if !slices.ContainsFunc(ddpd, func(kf []int) bool {
				return kf[0] == f[0] && kf[1] == f[1]
			}) {
				ddpd = append(ddpd, f)
			}
		}
		acc += int64(len(ddpd))
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

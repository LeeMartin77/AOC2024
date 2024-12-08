package solution

import (
	"fmt"
	"slices"
	"strings"
)

type antennaMap struct {
	FullMap                 map[int]map[int]rune
	MaxX                    int
	MaxY                    int
	AntennaCoordinateLookup map[rune][][]int
}

func parseMap(data []byte) antennaMap {
	amp := antennaMap{
		FullMap:                 map[int]map[int]rune{},
		AntennaCoordinateLookup: map[rune][][]int{},
	}
	lns := strings.Split(string(data), "\n")
	amp.MaxY = len(lns) - 1
	for y, ln := range lns {
		if y == 0 {
			for i := range len(ln) {
				amp.FullMap[i] = map[int]rune{}
			}
			amp.MaxX = len(ln) - 1
		}
		for x, rn := range []rune(ln) {
			if rn != '.' {
				amp.FullMap[x][y] = rn
				amp.AntennaCoordinateLookup[rn] = append(amp.AntennaCoordinateLookup[rn], []int{x, y})
			}
		}

	}
	return amp
}

func ComputeSolutionOne(data []byte) int64 {
	amp := parseMap(data)
	antinodes := map[rune][][]int{}

	for rn, cords := range amp.AntennaCoordinateLookup {
		for len(cords) > 1 {
			crd, cords2 := cords[0], cords[1:]
			for _, crd2 := range cords2 {

				dist := []int{crd[0] - crd2[0], crd[1] - crd2[1]}

				an1 := []int{crd[0] + dist[0], crd[1] + dist[1]}
				an2 := []int{crd2[0] - dist[0], crd2[1] - dist[1]}

				if an1[0] >= 0 && an1[0] <= amp.MaxX && an1[1] >= 0 && an1[1] <= amp.MaxY {
					antinodes[rn] = append(antinodes[rn], an1)
				}
				if an2[0] >= 0 && an2[0] <= amp.MaxX && an2[1] >= 0 && an2[1] <= amp.MaxY {
					antinodes[rn] = append(antinodes[rn], an2)
				}
			}
			cords = cords2
		}
	}

	allantinoces := [][]int{}
	for _, ans := range antinodes {
		for _, an := range ans {
			if !slices.ContainsFunc(allantinoces, func(crd []int) bool {
				return crd[0] == an[0] && crd[1] == an[1]
			}) {
				allantinoces = append(allantinoces, an)
			}
		}
	}
	//DebugMapAntinodes(amp, allantinoces)

	return int64(len(allantinoces))
}

func DebugMapAntinodes(amp antennaMap, allantinoces [][]int) {
	mp := []string{}
	for range amp.MaxY {
		mp = append(mp, "")
	}
	for x := range amp.MaxX {
		for y := range amp.MaxY {
			if slices.ContainsFunc(allantinoces, func(crd []int) bool {
				return crd[0] == x && crd[1] == y
			}) {
				mp[y] += "#"
			} else {
				mp[y] += "."
			}
		}
	}
	for _, ln := range mp {
		fmt.Println(ln)
	}
}

func ComputeSolutionTwo(data []byte) int64 {
	amp := parseMap(data)
	antinodes := map[rune][][]int{}

	for rn, cords := range amp.AntennaCoordinateLookup {
		for len(cords) > 1 {
			crd, cords2 := cords[0], cords[1:]
			for _, crd2 := range cords2 {

				dist := []int{crd[0] - crd2[0], crd[1] - crd2[1]}

				an1 := []int{crd[0] + dist[0], crd[1] + dist[1]}
				an2 := []int{an1[0] - dist[0], an1[1] - dist[1]}
				for (an1[0] >= 0 && an1[0] <= amp.MaxX && an1[1] >= 0 && an1[1] <= amp.MaxY) || (an2[0] >= 0 && an2[0] <= amp.MaxX && an2[1] >= 0 && an2[1] <= amp.MaxY) {
					if an1[0] >= 0 && an1[0] <= amp.MaxX && an1[1] >= 0 && an1[1] <= amp.MaxY {
						antinodes[rn] = append(antinodes[rn], an1)
					}
					if an2[0] >= 0 && an2[0] <= amp.MaxX && an2[1] >= 0 && an2[1] <= amp.MaxY {
						antinodes[rn] = append(antinodes[rn], an2)
					}
					an1 = []int{an1[0] + dist[0], an1[1] + dist[1]}
					an2 = []int{an2[0] - dist[0], an2[1] - dist[1]}
				}
			}
			cords = cords2
		}
	}

	allantinoces := [][]int{}
	for _, ans := range antinodes {
		for _, an := range ans {
			if !slices.ContainsFunc(allantinoces, func(crd []int) bool {
				return crd[0] == an[0] && crd[1] == an[1]
			}) {
				allantinoces = append(allantinoces, an)
			}
		}
	}
	//DebugMapAntinodes(amp, allantinoces)

	return int64(len(allantinoces))
}

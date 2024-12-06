package solution

import (
	"fmt"
	"slices"
	"strings"
)

type guard struct {
	PositionX        int
	PositionY        int
	Cardinality      cardinality
	StepsTaken       int
	LocationsVisited map[int]map[int][]cardinality
	OffMap           bool
	Looping          bool
}

type cardinality struct {
	X int
	Y int
}

func parseMap(data []byte) ([][]rune, *guard) {
	mp := [][]rune{}
	str := string(data)
	grd := guard{
		PositionX:        0,
		PositionY:        0,
		Cardinality:      cardinality{0, -1},
		StepsTaken:       0,
		LocationsVisited: map[int]map[int][]cardinality{},
	}
	for y, yln := range strings.Split(str, "\n") {
		xrns := []rune(yln)
		if y == 0 {
			for x := range xrns {
				mp = append(mp, []rune{})
				grd.LocationsVisited[x] = map[int][]cardinality{}
			}
		}
		for x, rn := range xrns {
			if rn == '^' {
				grd.PositionX = x
				grd.PositionY = y
				//replace with empty space
				rn = '.'
			}
			mp[x] = append(mp[x], rn)
		}
	}
	return mp, &grd

}

func (grd *guard) DistinctVisits() int {
	acc := 0
	for _, xmp := range grd.LocationsVisited {
		for _, y := range xmp {
			if len(y) > 0 {
				acc += 1
			}
		}
	}
	return acc
}

func (grd *guard) MoveUntilObstacleOrOffMapOrLooping(mp [][]rune) {
	for {
		next_x := grd.PositionX + grd.Cardinality.X
		next_y := grd.PositionY + grd.Cardinality.Y
		if next_y < 0 || len(mp) <= next_y || next_x < 0 || len(mp[0]) <= next_x {
			//off map
			grd.OffMap = true
			return
		}
		if mp[next_x][next_y] == '#' {
			// turn
			if grd.Cardinality.X == 0 && grd.Cardinality.Y == -1 {
				grd.Cardinality.X = 1
				grd.Cardinality.Y = 0
			} else if grd.Cardinality.X == 1 && grd.Cardinality.Y == 0 {
				grd.Cardinality.X = 0
				grd.Cardinality.Y = 1
			} else if grd.Cardinality.X == 0 && grd.Cardinality.Y == 1 {
				grd.Cardinality.X = -1
				grd.Cardinality.Y = 0
			} else if grd.Cardinality.X == -1 && grd.Cardinality.Y == 0 {
				grd.Cardinality.X = 0
				grd.Cardinality.Y = -1
			}
			return
		}
		grd.PositionX = next_x
		grd.PositionY = next_y

		if slices.ContainsFunc(grd.LocationsVisited[grd.PositionX][grd.PositionY], func(c cardinality) bool {
			return grd.Cardinality.X == c.X && grd.Cardinality.Y == c.Y
		}) {
			// we are now looping
			//fmt.Println("loop detected")
			grd.Looping = true
			return
		} else {
			grd.LocationsVisited[grd.PositionX][grd.PositionY] = append(grd.LocationsVisited[grd.PositionX][grd.PositionY], grd.Cardinality)
		}

		grd.StepsTaken += 1
	}
}

func ComputeSolutionOne(data []byte) int64 {
	mp, grd := parseMap(data)
	for !grd.OffMap {
		grd.MoveUntilObstacleOrOffMapOrLooping(mp)
	}
	return int64(grd.DistinctVisits())
}

func DebugPrintMap(mp [][]rune) {
	lines := make([]string, len(mp[0]))
	for x, ln := range mp {
		for _, c := range ln {
			lines[x] += string(c)
		}
	}
	fmt.Println("---")
	for _, line := range lines {

		fmt.Println(line)
	}
}

func ComputeSolutionTwo(data []byte) int64 {
	// fuck it
	mpbs, _ := parseMap(data)

	acc := 0
	for x := range mpbs {
		for y := range mpbs[0] {
			mp, grd := parseMap(data)

			mp[x][y] = '#'

			//DebugPrintMap(mp)

			for !grd.OffMap && !grd.Looping {
				grd.MoveUntilObstacleOrOffMapOrLooping(mp)
				if grd.Looping {
					break
				}
			}
			if grd.Looping {
				acc += 1
			}
		}
	}

	return int64(acc)
}

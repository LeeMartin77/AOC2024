package solution

import (
	"strings"
)

type guard struct {
	PositionX   int
	PositionY   int
	Cardinality struct {
		X int
		Y int
	}
	StepsTaken       int
	LocationsVisited map[int]map[int]bool
	OffMap           bool
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
		LocationsVisited: map[int]map[int]bool{},
	}
	for y, yln := range strings.Split(str, "\n") {
		xrns := []rune(yln)
		if y == 0 {
			for x := range xrns {
				mp = append(mp, []rune{})
				grd.LocationsVisited[x] = map[int]bool{}
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
			if y {
				acc += 1
			}
		}
	}
	return acc
}

func (grd *guard) MoveUntilObstacleOrOffMap(mp [][]rune) {
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
		grd.LocationsVisited[next_x][next_y] = true
		grd.StepsTaken += 1
	}
}

func ComputeSolutionOne(data []byte) int64 {
	mp, grd := parseMap(data)
	for !grd.OffMap {
		grd.MoveUntilObstacleOrOffMap(mp)
	}
	return int64(grd.DistinctVisits())
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

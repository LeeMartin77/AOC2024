package solution

import (
	"fmt"
	"strconv"
	"strings"
)

type robot struct {
	Idx              int
	StartingPosition []int64
	Position         []int64
	StartingVelocity []int64
	Velocity         []int64
}

func ParseRobots(data []byte) ([]*robot, int64, int64) {
	maxX, maxY := int64(0), int64(0)
	rbts := []*robot{}
	for idx, ln := range strings.Split(string(data), "\n") {
		prts := strings.Split(ln, " ")
		pospts := strings.Split(prts[0], "=")
		velprts := strings.Split(prts[1], "=")
		pospts = strings.Split(pospts[1], ",")
		velprts = strings.Split(velprts[1], ",")
		x, _ := strconv.ParseInt(pospts[0], 10, 64)
		y, _ := strconv.ParseInt(pospts[1], 10, 64)
		vx, _ := strconv.ParseInt(velprts[0], 10, 64)
		vy, _ := strconv.ParseInt(velprts[1], 10, 64)
		rbts = append(rbts, &robot{
			Idx:              idx,
			StartingPosition: []int64{x, y},
			Position:         []int64{x, y},
			StartingVelocity: []int64{vx, vy},
			Velocity:         []int64{vx, vy},
		})
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}
	return rbts, maxX, maxY
}

func (rbt *robot) MoveTicks(ticks, limitX, limitY int64) {
	// we probably are going to have to math this, but lets just do this for now
	for range ticks {
		newx, newy := rbt.Position[0], rbt.Position[1]
		newx += rbt.Velocity[0]
		for newx < 0 {
			newx = limitX + (newx + 1)
		}
		for newx > limitX {
			newx = newx - (limitX + 1)
		}
		newy += rbt.Velocity[1]
		for newy < 0 {
			newy = limitY + (newy + 1)
		}
		for newy > limitY {
			newy = newy - (limitY + 1)
		}
		rbt.Position = []int64{newx, newy}
	}
}

func (rbt *robot) GetQuadrant(limitX, limitY int64) int64 {
	middle := []int64{limitX / 2, limitY / 2}
	if rbt.Position[0] < middle[0] && rbt.Position[1] < middle[1] {
		return 0
	}
	if rbt.Position[0] > middle[0] && rbt.Position[1] < middle[1] {
		return 1
	}
	if rbt.Position[0] < middle[0] && rbt.Position[1] > middle[1] {
		return 2
	}
	if rbt.Position[0] > middle[0] && rbt.Position[1] > middle[1] {
		return 3
	}
	return -1
}

func DebugPositions(rbts []*robot, maxX, maxY int64) {
	// this can be crude
	for _, rbt := range rbts {
		fmt.Println(rbt.Position)
	}
	fmt.Println("---")
	for y := range maxY + 1 {
		for x := range maxX + 1 {
			cnt := 0
			for _, rbt := range rbts {
				if rbt.Position[0] == x && rbt.Position[1] == y {
					cnt += 1
				}
			}
			if cnt == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", cnt)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("---")
}

func ComputeSolutionOne(data []byte) int64 {
	rbts, maxX, maxY := ParseRobots(data)
	qdrnts := []int64{0, 0, 0, 0}
	for _, rbt := range rbts {
		rbt.MoveTicks(100, maxX, maxY)
		qdrnt := rbt.GetQuadrant(maxX, maxY)
		if qdrnt != -1 {
			qdrnts[qdrnt] += 1
		}
	}

	//DebugPositions(rbts, maxX, maxY)

	return qdrnts[0] * qdrnts[1] * qdrnts[2] * qdrnts[3]
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

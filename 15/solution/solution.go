package solution

import (
	"fmt"
	"strings"
)

type warehouse struct {
	Walls map[int]map[int]bool
	Boxes map[int]map[int]int // we will store "box indexes" to keep them distinct
	Size  []int
}

type robot struct {
	// x, y
	Position []int
}

func parseWarehouseAndRobot(data []byte) (warehouse, robot, string) {
	wrhs := warehouse{
		Walls: map[int]map[int]bool{},
		Boxes: map[int]map[int]int{},
	}
	rbt := robot{}
	mvs := ""

	mp := true
	bxidx := 1
	wdth := len(strings.Split(string(data), "\n")[0])
	for x := range wdth {
		wrhs.Walls[x] = map[int]bool{}
		wrhs.Boxes[x] = map[int]int{}
	}

	wrhs.Size = []int{wdth, 0}

	for y, ln := range strings.Split(string(data), "\n") {
		if mp {
			for x, rn := range ln {
				switch rn {
				case '.':
					continue
				case '#':
					wrhs.Walls[x][y] = true
				case 'O':
					wrhs.Boxes[x][y] = bxidx
					bxidx += 1
				case '@':
					rbt.Position = []int{x, y}
				}
			}
			if ln == "" {
				mp = false
			} else {
				wrhs.Size[1] = wrhs.Size[1] + 1
			}
		} else {
			// just the instructions
			mvs += ln
		}
	}

	return wrhs, rbt, mvs
}

func moveRobotAroundMap(wrhs warehouse, rbt robot, mvs string) (warehouse, robot) {
	for _, mv := range mvs {
		var potentialmove []int
		switch mv {
		case '<':
			potentialmove = []int{-1, 0}
		case 'v':
			potentialmove = []int{0, 1}
		case '^':
			potentialmove = []int{0, -1}
		case '>':
			potentialmove = []int{1, 0}
		default:
			panic("unrecognised command")
		}
		willmove := true
		potentialrobotposition := []int{rbt.Position[0] + potentialmove[0], rbt.Position[1] + potentialmove[1]}
		from := []int{rbt.Position[0] + potentialmove[0], rbt.Position[1] + potentialmove[1]}
		for {
			if !wrhs.Walls[from[0]][from[1]] && wrhs.Boxes[from[0]][from[1]] == 0 {
				// a free space, so break
				break
			}
			if wrhs.Walls[from[0]][from[1]] {
				// won't move, flip switch and break
				willmove = false
				break
			}
			//if wrhs.Boxes[from[0]][from[1]] != 0
			from = []int{from[0] + potentialmove[0], from[1] + potentialmove[1]}
		}
		if !willmove {
			continue
		}
		for {
			if from[0] == potentialrobotposition[0] && from[1] == potentialrobotposition[1] {
				break
			}
			// get the box, and move it
			srcps := []int{from[0] - potentialmove[0], from[1] - potentialmove[1]}
			wrhs.Boxes[from[0]][from[1]] = wrhs.Boxes[srcps[0]][srcps[1]]
			delete(wrhs.Boxes[srcps[0]], srcps[1])
			from = srcps
		}
		rbt.Position[0] = potentialrobotposition[0]
		rbt.Position[1] = potentialrobotposition[1]
	}
	return wrhs, rbt
}

func (wrhs warehouse) generateGPSScore() int {
	gps := 0
	for x, mp := range wrhs.Boxes {
		for y := range mp {
			gps += (y*100 + x)
		}
	}
	return gps
}

func DebugWarehouse(wrhs warehouse, rbt robot) {
	for y := range wrhs.Size[1] {
		for x := range wrhs.Size[0] {
			if wrhs.Boxes[x][y] != 0 && wrhs.Walls[x][y] {
				// error - box in wall
				fmt.Print("X")
				continue
			} else if wrhs.Boxes[x][y] != 0 && rbt.Position[0] == x && rbt.Position[1] == y {
				// error - robot in box
				fmt.Print("Y")
				continue
			} else if wrhs.Walls[x][y] && rbt.Position[0] == x && rbt.Position[1] == y {
				// error - robot in wall
				fmt.Print("&")
				continue
			} else if wrhs.Boxes[x][y] != 0 {
				fmt.Print("O")
				continue
			} else if rbt.Position[0] == x && rbt.Position[1] == y {
				fmt.Print("@")
				continue
			} else if wrhs.Walls[x][y] {
				fmt.Print("#")
				continue
			} else {
				fmt.Print(".")
				continue
			}
		}
		fmt.Print("\n")
	}
}

func ComputeSolutionOne(data []byte) int64 {
	wrhs, rbt, mvs := parseWarehouseAndRobot(data)

	//DebugWarehouse(wrhs, rbt)

	wrhs, _ = moveRobotAroundMap(wrhs, rbt, mvs)

	//DebugWarehouse(wrhs, rbt)

	return int64(wrhs.generateGPSScore())
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

package solution

import (
	"fmt"
	"slices"
	"strings"
)

type doublewideBox struct {
	Position [][]int
}
type doublewideWarehouse struct {
	Walls map[int]map[int]bool
	Boxes []*doublewideBox
	Size  []int
}

// -1 for left, 1 for right, 0 for not
func (bx *doublewideBox) OnCord(crd []int) int {
	if bx.Position[0][0] == crd[0] && bx.Position[0][1] == crd[1] {
		return -1
	}
	if bx.Position[1][0] == crd[0] && bx.Position[1][1] == crd[1] {
		return 1
	}
	return 0
}

func (bx *doublewideBox) MoveBox(crd []int) {
	bx.Position[0][0] += crd[0]
	bx.Position[0][1] += crd[1]
	bx.Position[1][0] += crd[0]
	bx.Position[1][1] += crd[1]
}

func (dwwhs doublewideWarehouse) HasABoxOnCord(crd []int) bool {
	for _, bx := range dwwhs.Boxes {
		if bx.OnCord(crd) != 0 {
			return true
		}
	}
	return false
}

func DebugDoublewideWarehouse(wrhs doublewideWarehouse, rbt robot) {
	for y := range wrhs.Size[1] {
		for x := range wrhs.Size[0] {
			if wrhs.HasABoxOnCord([]int{x, y}) && wrhs.Walls[x][y] {
				// error - box in wall
				fmt.Print("X")
				continue
			} else if wrhs.HasABoxOnCord([]int{x, y}) && rbt.Position[0] == x && rbt.Position[1] == y {
				// error - robot in box
				fmt.Print("Y")
				continue
			} else if wrhs.Walls[x][y] && rbt.Position[0] == x && rbt.Position[1] == y {
				// error - robot in wall
				fmt.Print("&")
				continue
			} else if wrhs.HasABoxOnCord([]int{x, y}) {
				for _, bx := range wrhs.Boxes {
					if bx.OnCord([]int{x, y}) == -1 {
						fmt.Print("[")
						break
					} else if bx.OnCord([]int{x, y}) == 1 {
						fmt.Print("]")
						break
					}
				}
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

func parseDoublewideWarehouseAndRobot(data []byte) (doublewideWarehouse, robot, string) {
	wrhs := doublewideWarehouse{
		Walls: map[int]map[int]bool{},
	}
	rbt := robot{}
	mvs := ""

	mp := true
	wdth := len(strings.Split(string(data), "\n")[0]) * 2
	for x := range wdth {
		wrhs.Walls[x] = map[int]bool{}
	}

	wrhs.Size = []int{wdth, 0}

	for y, ln := range strings.Split(string(data), "\n") {
		if mp {
			x := 0
			for _, rn := range ln {
				switch rn {
				case '#':
					wrhs.Walls[x][y] = true
					wrhs.Walls[x+1][y] = true
				case 'O':
					wrhs.Boxes = append(wrhs.Boxes, &doublewideBox{[][]int{
						{x, y},
						{x + 1, y},
					}})
				case '@':
					rbt.Position = []int{x, y}
				}
				x += 2
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

func (bx *doublewideBox) Froms(dir []int) [][]int {
	// if left/right - just return the left right cords
	if dir[0] == -1 {
		return [][]int{{bx.Position[0][0] - 1, bx.Position[0][1]}}
	}
	if dir[0] == 1 {
		return [][]int{{bx.Position[0][0] + 2, bx.Position[0][1]}}
	}
	// if up/down - return the pair

	if dir[1] == -1 {
		return [][]int{{bx.Position[0][0], bx.Position[0][1] - 1}, {bx.Position[1][0], bx.Position[1][1] - 1}}
	}
	if dir[1] == 1 {
		return [][]int{{bx.Position[0][0], bx.Position[0][1] + 1}, {bx.Position[1][0], bx.Position[1][1] + 1}}
	}
	panic("AH FUCK")
}

func moveRobotAroundDoublewide(wrhs doublewideWarehouse, rbt robot, mvs string) (doublewideWarehouse, robot) {
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
		from := [][]int{{rbt.Position[0] + potentialmove[0], rbt.Position[1] + potentialmove[1]}}
		bxidxmoving := []int{}
		for {
			if !slices.ContainsFunc(from, func(crd []int) bool {
				return wrhs.Walls[crd[0]][crd[1]]
			}) && !slices.ContainsFunc(from, func(crd []int) bool {
				return wrhs.HasABoxOnCord(crd)
			}) {
				// all free spaces, so break
				break
			}
			if slices.ContainsFunc(from, func(crd []int) bool {
				return wrhs.Walls[crd[0]][crd[1]]
			}) {
				// won't move, hitting a wall, flip switch and break
				willmove = false
				break
			}
			// add unique box idxes that need to move
			newfrm := [][]int{}
			for _, crd := range from {

				for idx, bx := range wrhs.Boxes {
					if bx.OnCord(crd) != 0 {
						if !slices.Contains(bxidxmoving, idx) {
							bxidxmoving = append(bxidxmoving, idx)
							newfrm = append(newfrm, bx.Froms(potentialmove)...)
						}
					}
				}
			}
			// if wrhs.Boxes[from[0]][from[1]] != 0
			from = newfrm
		}
		if !willmove {
			continue
		}
		for _, bidx := range bxidxmoving {
			wrhs.Boxes[bidx].MoveBox(potentialmove)
		}
		rbt.Position[0] = potentialrobotposition[0]
		rbt.Position[1] = potentialrobotposition[1]
	}
	return wrhs, rbt
}

func (wrhs doublewideWarehouse) generateDWGPSScore() int {
	gps := 0
	for _, mp := range wrhs.Boxes {
		gps += (mp.Position[0][1]*100 + mp.Position[0][0])
	}
	return gps
}

func ComputeSolutionTwo(data []byte) int64 {
	dwwhs, rbt, cmd := parseDoublewideWarehouseAndRobot(data)

	//DebugDoublewideWarehouse(dwwhs, rbt)

	dwwhs, _ = moveRobotAroundDoublewide(dwwhs, rbt, cmd)

	//DebugDoublewideWarehouse(dwwhs, rbt)

	return int64(dwwhs.generateDWGPSScore())
}

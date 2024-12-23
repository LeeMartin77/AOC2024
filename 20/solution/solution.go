package solution

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func parseMaze(data []byte) (map[int]map[int]bool, []int, []int, int, int) {
	maze := map[int]map[int]bool{}
	start := []int{}
	exit := []int{}
	width := 0
	height := 0
	for y, ln := range strings.Split(string(data), "\n") {
		if y == 0 {
			for i := range len(ln) {
				maze[i] = map[int]bool{}
			}
			width = len(ln)
		}
		for x, rn := range ln {
			switch rn {
			case '#':
				maze[x][y] = false
			case '.':
				maze[x][y] = true
			case 'S':
				maze[x][y] = true
				start = []int{x, y}
			case 'E':
				maze[x][y] = true
				exit = []int{x, y}
			}
		}
		height += 1
	}
	return maze, start, exit, width, height
}

// the puzzle isn't a maze, it's just a "drag race"

func ComputeSolutionOne(data []byte, care_about_diff int) int64 {
	maze, s, e, _, _ := parseMaze(data)
	rt := getMazeRoute(maze, s, e)

	mvs := manhattan_moves(rt, 2)

	//DebugMapAndLocations(maze, rt, w, h)

	acc := 0
	for svng, cnt := range mvs {
		if svng >= care_about_diff {
			acc += cnt
		}
	}
	return int64(acc)
}

func manhattan_moves(rt [][]int, spaces int) map[int]int { // savings: count
	svs := map[int]int{}
	for i := range rt {
		for ii := range rt {
			dist := distance_between_cords(rt[i], rt[ii])
			if dist <= spaces && ii-i > dist {
				sv := ((ii - i) - dist)

				if sv > 1 {
					svs[sv] += 1
				}
			}
		}
	}
	return svs
}

func distance_between_cords(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func ComputeSolutionTwo(data []byte, care_about_diff int) int64 {
	maze, s, e, _, _ := parseMaze(data)
	rt := getMazeRoute(maze, s, e)

	mvs := manhattan_moves(rt, 20)

	acc := 0
	for svng, cnt := range mvs {
		if svng >= care_about_diff {
			acc += cnt
		}
	}
	return int64(acc)
}

func getMazeRoute(maze map[int]map[int]bool, s, e []int) [][]int {
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	pos := []int{s[0], s[1]}
	rt := [][]int{}
	for !(pos[0] == e[0] && pos[1] == e[1]) {
		for _, dir := range directions {
			nxt := []int{pos[0] + dir[0], pos[1] + dir[1]}
			lst := []int{s[0], s[1]}
			if len(rt) > 0 {
				lst = []int{rt[len(rt)-1][0], rt[len(rt)-1][1]}
			}
			if maze[nxt[0]][nxt[1]] && !(nxt[0] == lst[0] && nxt[1] == lst[1]) {
				rt = append(rt, []int{pos[0], pos[1]})
				pos[0] = nxt[0]
				pos[1] = nxt[1]
				break
			}
		}
	}
	return append(rt, pos)
}

func DebugMapAndLocations(maze map[int]map[int]bool, locs [][]int, w, h int) {
	fmt.Println("---")
	for y := range h {
		for x := range w {
			if slices.ContainsFunc(locs, func(lc []int) bool {
				return lc[0] == x && lc[1] == y
			}) {
				fmt.Print("o")
				continue
			}
			if !maze[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("---")
}

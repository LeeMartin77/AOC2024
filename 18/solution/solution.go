package solution

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parseGrid(data []byte) (map[int]map[int]bool, [][]int, []int, []int) {
	maze := map[int]map[int]bool{}
	corruptions := [][]int{}
	start := []int{0, 0}
	exit := []int{0, 0}

	for _, ln := range strings.Split(string(data), "\n") {
		prts := strings.Split(ln, ",")
		x, _ := strconv.ParseInt(prts[0], 10, 64)
		y, _ := strconv.ParseInt(prts[1], 10, 64)

		if x > int64(exit[0]) {
			exit[0] = int(x)
		}
		if y > int64(exit[1]) {
			exit[1] = int(y)
		}
		corruptions = append(corruptions, []int{int(x), int(y)})
	}
	for x := range exit[0] + 1 {
		maze[x] = map[int]bool{}
	}
	for x := range exit[0] + 1 {
		for y := range exit[1] + 1 {
			maze[x][y] = true
		}
	}
	return maze, corruptions, start, exit
}

func ComputeSolutionOne(data []byte) int64 {
	maze, cor, start, exit := parseGrid(data)
	dataToFall := 1024
	for i := range dataToFall {
		if len(cor) > i {
			cr := cor[i]
			maze[cr[0]][cr[1]] = false
		} else {
			continue
		}
	}
	_, route := a_star(start, exit, maze, heuristic)
	//DebugMapAndLocations(maze, route, exit[0]+1, exit[1]+1)
	return int64(len(route) - 1)
}

func ComputeSolutionTwo(data []byte) string {
	maze, cor, start, exit := parseGrid(data)
	for _, cr := range cor {
		maze[cr[0]][cr[1]] = false
		scr, _ := a_star(start, exit, maze, heuristic)
		if scr == -10 {
			// we can't escape anymore

			return fmt.Sprintf("%d,%d", cr[0], cr[1])
		}
	}
	return "ERROR: NEVER BLOCKED"
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

func string_cord(cord []int) string {
	return fmt.Sprintf("%d:%d", cord[0], cord[1])
}
func cord_string(str string) []int {
	prts := strings.Split(str, ":")
	x, _ := strconv.ParseInt(prts[0], 10, 32)
	y, _ := strconv.ParseInt(prts[1], 10, 32)
	return []int{int(x), int(y)}
}

func reconstruct_path(cameFrom map[string]string, current []int) [][]int {
	path := [][]int{current}
	for len(cameFrom[string_cord(path[0])]) > 0 {
		path = append([][]int{cord_string(cameFrom[string_cord(path[0])])}, path...)
	}
	return path
}

func edgecalc(current, next string) int {
	// if forwards, 1 cost
	return 1
}

func heuristic(self string, goal []int) int {
	self_cord := cord_string(self)
	dist := math.Pow(float64(goal[0]-self_cord[0]), 2) + math.Pow(float64(goal[1]-self_cord[1]), 2)
	return int(math.Sqrt(dist))
}

func a_star(start []int, goal []int, maze map[int]map[int]bool, h func(string, []int) int) (int, [][]int) {
	openset := []string{
		string_cord(start),
	}

	cameFrom := map[string]string{}

	gScore := map[string]int{}
	for x, xm := range maze {
		for y := range xm {
			gScore[string_cord([]int{x, y})] = 9999999999999999
		}
	}
	gScore[string_cord(start)] = 0

	fscore := map[string]int{}
	fscore[string_cord(start)] = h(string_cord(start), goal)

	for len(openset) > 0 {
		current := openset[0]
		if current == string_cord(goal) {
			return gScore[current], reconstruct_path(cameFrom, goal)
		}
		newos := []string{}
		for idx, str := range openset {
			if idx != 0 {
				newos = append(newos, str)
			}
		}
		openset = newos
		// for neighbours of current
		dirs := [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}
		cur_cord := cord_string(current)
		for _, dir := range dirs {
			ncrd := []int{cur_cord[0] + dir[0], cur_cord[1] + dir[1]}
			if !maze[ncrd[0]][ncrd[1]] {
				// completely invalid move
				continue
			}
			neighbour := string_cord(ncrd)
			tentative := gScore[current] + edgecalc(current, neighbour)
			if tentative < gScore[neighbour] {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentative
				fscore[neighbour] = tentative + h(neighbour, goal)
				if !slices.Contains(openset, neighbour) {
					openset = append(openset, neighbour)
				}
			}
		}
		// sort openset
		slices.SortFunc(openset, func(a, b string) int {
			return fscore[a] - fscore[b]
		})
	}
	return -10, [][]int{}
}

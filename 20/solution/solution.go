package solution

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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
	_, rt := a_star(s, e, maze, heuristic)
	//non_cheat_cost := len(rt)
	acc := int64(0)

	cord_distances := map[string]int{}
	for dis, cord := range rt {
		cord_distances[string_cord(cord)] = dis
	}
	for dist, cord := range rt {
		// we check to see if any space 2 away is in our route
		diffs := [][]int{
			{-2, 0},
			{2, 0},
			{0, 2},
			{0, -2},
		}
		for _, diff := range diffs {

			neighbour_dist := cord_distances[string_cord([]int{
				cord[0] + diff[0],
				cord[1] + diff[1],
			})]
			if neighbour_dist > (dist + 1) {
				// shortcut!
				saving := -(dist - neighbour_dist)
				saving -= 1
				if saving > care_about_diff {
					acc += 1
				}
			}
		}
	}
	//DebugMapAndLocations(maze, rt, w, h)
	return acc
}

func distBetween(neighbour []int, cord []int) int {
	return int(math.Abs(float64((neighbour[0] - cord[0]))) + math.Abs(float64((neighbour[1] - cord[1]))))
}
func ComputeSolutionTwo(data []byte, care_about_diff int) int64 {
	maze, s, e, _, _ := parseMaze(data)
	_, rt := a_star(s, e, maze, heuristic)
	//non_cheat_cost := len(rt)
	acc := int64(0)

	cord_distances := map[string]int{}
	for dis, cord := range rt {
		cord_distances[string_cord(cord)] = dis
	}
	max_cheat := 20

	diffs := [][]int{}
	ranges := []int{}
	for r := range max_cheat + 10 {
		if r != 0 {
			ranges = append(ranges, -r)
			ranges = append(ranges, r)
		} else {
			ranges = append(ranges, 0)
		}
	}

	for x := range ranges {
		for y := range ranges {
			if distBetween([]int{0, 0}, []int{x, y}) <= max_cheat {

				diffs = append(diffs, []int{x, y})
			}
		}
	}
	for dist, cord := range rt {
		for _, diff := range diffs {

			neighbour := []int{
				cord[0] + diff[0],
				cord[1] + diff[1],
			}
			neighbour_dist := cord_distances[string_cord(neighbour)]
			path_dist := distBetween(neighbour, cord)
			if neighbour_dist >= (dist + int(path_dist)) {
				// potential shortcut!
				// that is less than the max cheat

				saving := -(dist - neighbour_dist)
				//saving -= (int(path_dist) - )
				// we are annoyingly close and looks like we have some weird off-by-one style error happening
				if saving >= (care_about_diff - 2) {
					acc += 1
				}
			}
		}
	}
	//DebugMapAndLocations(maze, rt, w, h)
	return acc
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

func edgecalc(_, _ string) int {
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

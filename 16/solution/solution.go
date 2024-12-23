package solution

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
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
				continue
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

type reindeer struct {
	Position []int
	Facing   []int
	Score    int
	TurnCost int
	MoveCost int
	History  [][]int
}

func NewReindeer(startpos []int, startfacing []int, startscore int, turncost int, movecost int, history [][]int) reindeer {
	return reindeer{
		Position: startpos,
		Facing:   startfacing,
		Score:    startscore,
		TurnCost: turncost,
		MoveCost: movecost,
		History:  history,
	}
}

func (rnd reindeer) HasVisited(x, y int) bool {
	// preventing loops
	return slices.ContainsFunc(rnd.History, func(crd []int) bool {
		return crd[0] == x && crd[1] == y
	})
}

func (rnd reindeer) GenerateFuturePossibleReindeer(maze map[int]map[int]bool) []reindeer {
	ftr := []reindeer{}
	// straight ahead
	if maze[rnd.Position[0]+rnd.Facing[0]][rnd.Position[1]+rnd.Facing[1]] && !rnd.HasVisited(rnd.Position[0]+rnd.Facing[0], rnd.Position[1]+rnd.Facing[1]) {
		ftr = append(ftr, NewReindeer(
			[]int{rnd.Position[0] + rnd.Facing[0], rnd.Position[1] + rnd.Facing[1]},
			[]int{rnd.Facing[0], rnd.Facing[1]},
			rnd.Score+rnd.MoveCost,
			rnd.TurnCost,
			rnd.MoveCost,
			append(rnd.History, rnd.Position),
		))
	}
	if rnd.Facing[0] == 0 {
		// turning "left"
		if maze[rnd.Position[0]-1][rnd.Position[1]] && !rnd.HasVisited(rnd.Position[0]-1, rnd.Position[1]) {
			ftr = append(ftr, NewReindeer(
				[]int{rnd.Position[0] - 1, rnd.Position[1]},
				[]int{-1, 0},
				rnd.Score+rnd.MoveCost+rnd.TurnCost,
				rnd.TurnCost,
				rnd.MoveCost,
				append(rnd.History, rnd.Position),
			))
		}
		// turning "right"
		if maze[rnd.Position[0]+1][rnd.Position[1]] && !rnd.HasVisited(rnd.Position[0]+1, rnd.Position[1]) {
			ftr = append(ftr, NewReindeer(
				[]int{rnd.Position[0] + 1, rnd.Position[1]},
				[]int{1, 0},
				rnd.Score+rnd.MoveCost+rnd.TurnCost,
				rnd.TurnCost,
				rnd.MoveCost,
				append(rnd.History, rnd.Position),
			))
		}
	}
	if rnd.Facing[1] == 0 {

		// turning "left"
		if maze[rnd.Position[0]][rnd.Position[1]-1] && !rnd.HasVisited(rnd.Position[0], rnd.Position[1]-1) {
			ftr = append(ftr, NewReindeer(
				[]int{rnd.Position[0], rnd.Position[1] - 1},
				[]int{0, -1},
				rnd.Score+rnd.MoveCost+rnd.TurnCost,
				rnd.TurnCost,
				rnd.MoveCost,
				append(rnd.History, rnd.Position),
			))
		}
		// turning "right"
		if maze[rnd.Position[0]][rnd.Position[1]+1] && !rnd.HasVisited(rnd.Position[0], rnd.Position[1]+1) {
			ftr = append(ftr, NewReindeer(
				[]int{rnd.Position[0], rnd.Position[1] + 1},
				[]int{0, 1},
				rnd.Score+rnd.MoveCost+rnd.TurnCost,
				rnd.TurnCost,
				rnd.MoveCost,
				append(rnd.History, rnd.Position),
			))
		}
	}
	return ftr
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

func edgecalc(current, next, last string) int {
	cur := cord_string(current)
	if last == "" {
		// if last == "" - edge case of the starting tile, "facing" east, meaning west of current
		last = string_cord([]int{cur[0] - 1, cur[1]})
	}
	n, lst := cord_string(next), cord_string(last)

	// if backwards - massive penalty, 9999999999 cost
	if n[0] == lst[0] && n[1] == lst[1] {
		return 999999999999
	}

	// if "turning", 1001 cost
	dir_next := []int{cur[0] - n[0], cur[1] - n[1]}
	dir_last := []int{lst[0] - cur[0], lst[1] - cur[1]}
	if dir_next[0] != dir_last[0] || dir_next[1] != dir_last[1] {
		return 1001
	}
	// if forwards, 1 cost
	return 1
}

func heuristic(self string, goal []int) int {
	// so we "worst case" this
	self_cord := cord_string(self)
	dist := math.Abs(float64(goal[0]-self_cord[0])) + math.Abs(float64(goal[1]-self_cord[1]))
	return int(dist * 100)
}

func a_star(start []int, goal []int, maze map[int]map[int]bool, h func(string, []int) int, src []int) (int, [][]int) {
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
			cf := cameFrom[current]
			if cf == "" {
				// start - replace with src
				cf = string_cord(src)
			}
			tentative := gScore[current] + edgecalc(current, neighbour, cf)
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

func ComputeSolutionOne(data []byte) int64 {
	maze, start, exit, _, _ := parseMaze(data)

	most_expensive_destination := 0
	destcosts := map[int]map[int]int{}

	dests := make(chan []int)
	donedests := make(chan struct{})

	go func() {
		for dst := range dests {
			destcosts[dst[0]][dst[1]] = dst[2]
			if dst[2] > most_expensive_destination {
				most_expensive_destination = dst[2]
			}
		}
		donedests <- struct{}{}
	}()

	wg := sync.WaitGroup{}
	for x, mx := range maze {
		destcosts[x] = map[int]int{}
		for y := range mx {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if x != exit[0] && y != exit[1] {
					scr, _ := a_star(exit, []int{x, y}, maze, heuristic, []int{exit[0] - 1, exit[1]})
					dests <- []int{x, y, scr}
				}
			}()
		}
	}
	wg.Wait()
	close(dests)
	<-donedests

	hh := func(self string, goal []int) int {
		self_cord := cord_string(self)

		return most_expensive_destination - destcosts[self_cord[0]][self_cord[1]]
	}

	scr, _ := a_star(start, exit, maze, hh, []int{start[0] - 1, start[1]})

	//DebugMapAndLocations(maze, route, w, h)

	return int64(scr)
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

func can_visit(visited map[string]*int64, d, i, j int, score int64) bool {
	prev_score := visited[fmt.Sprintf("%d,%d,%d", d, i, j)]
	if prev_score != nil && *prev_score < score {
		return false
	}
	visited[fmt.Sprintf("%d,%d,%d", d, i, j)] = &score
	return true
}

type heapitem struct {
	score int64
	d     int
	x     int
	y     int
	path  [][]int
}

func ComputeSolutionTwo(data []byte) int64 {
	maze, start, exit, _, _ := parseMaze(data)
	visited := map[string]*int64{}
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	heap := []heapitem{{0, 0, start[0], start[1], [][]int{}}}
	lowest_score := int64(0)
	winning_paths := map[int64]map[string]bool{}
	for len(heap) > 0 {
		var hi heapitem
		if len(heap) > 1 {
			hi, heap = heap[0], heap[1:]
		} else {
			hi, heap = heap[0], []heapitem{}
		}
		score, d, x, y, path := hi.score, hi.d, hi.x, hi.y, hi.path
		if lowest_score != 0 && lowest_score < score {
			continue
		}
		if x == exit[0] && y == exit[1] {
			lowest_score = score
			if winning_paths[lowest_score] == nil {
				winning_paths[lowest_score] = map[string]bool{}
			}
			for _, p := range path {
				winning_paths[lowest_score][fmt.Sprintf("%d,%d", p[0], p[1])] = true
			}
			continue
		}
		if !can_visit(visited, d, x, y, score) {
			continue
		}
		xx := x + directions[d][0]
		yy := y + directions[d][1]
		if maze[xx][yy] && can_visit(visited, d, xx, yy, score+1) {
			heap = append(heap, heapitem{score + 1, d, xx, yy, append(path, []int{x, y})})
		}
		left := 0
		switch d {
		case 0:
			left = 1
		case 1:
			left = 2
		case 2:
			left = 3
		case 3:
			left = 0
		}
		if can_visit(visited, left, x, y, score+1000) {
			heap = append(heap, heapitem{score + 1000, left, x, y, append([][]int{}, path...)})
		}
		right := 0
		switch d {
		case 0:
			right = 3
		case 1:
			right = 0
		case 2:
			right = 1
		case 3:
			right = 2
		}
		if can_visit(visited, right, x, y, score+1000) {
			heap = append(heap, heapitem{score + 1000, right, x, y, append([][]int{}, path...)})
		}
	}
	//DebugMapAndLocations(maze, locs, w, h)
	return int64(len(winning_paths[lowest_score]) + 1)
}

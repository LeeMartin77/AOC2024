package solution

import (
	"slices"
	"strings"
)

func parseMaze(data []byte) (map[int]map[int]bool, []int, []int) {
	maze := map[int]map[int]bool{}
	start := []int{}
	exit := []int{}
	for y, ln := range strings.Split(string(data), "\n") {
		if y == 0 {
			for i := range len(ln) {
				maze[i] = map[int]bool{}
			}
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
	}
	return maze, start, exit
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
	if rnd.Facing[0] != 0 {

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
				[]int{0, +1},
				rnd.Score+rnd.MoveCost+rnd.TurnCost,
				rnd.TurnCost,
				rnd.MoveCost,
				append(rnd.History, rnd.Position),
			))
		}
	}
	return ftr
}

func recurrFlock(maze map[int]map[int]bool, exit []int, flock []reindeer, current_low_score int) int {
	if len(flock) == 0 {
		return current_low_score
	}
	remaining := []reindeer{}
	for _, rnd := range flock {
		if rnd.Position[0] == exit[0] && rnd.Position[1] == exit[1] && current_low_score > rnd.Score {
			current_low_score = rnd.Score
		}
		if rnd.Score < current_low_score {
			remaining = append(remaining, rnd)
		}
	}
	future := []reindeer{}
	for _, rnd := range remaining {
		future = append(future, rnd.GenerateFuturePossibleReindeer(maze)...)
	}
	return recurrFlock(maze, exit, future, current_low_score)
}

func ComputeSolutionOne(data []byte) int64 {
	maze, start, exit := parseMaze(data)

	current_low_score := 9999999999999

	flock := []reindeer{NewReindeer(start, []int{1, 0}, 0, 1000, 1, [][]int{start})}

	low_score := recurrFlock(maze, exit, flock, current_low_score)

	return int64(low_score)
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

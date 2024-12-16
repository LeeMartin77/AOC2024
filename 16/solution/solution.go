package solution

import (
	"slices"
	"strings"
	"sync"
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

func recurrFlock(maze map[int]map[int]bool, exit []int, flock []reindeer, current_low_score reindeer) reindeer {
	if len(flock) == 0 {
		return current_low_score
	}
	remaining := []reindeer{}
	for _, rnd := range flock {
		if rnd.Position[0] == exit[0] && rnd.Position[1] == exit[1] && current_low_score.Score > rnd.Score {
			current_low_score = rnd
		}
		if rnd.Score < current_low_score.Score {
			remaining = append(remaining, rnd)
		}
	}
	future := []reindeer{}
	wg := sync.WaitGroup{}
	ftrchn := make(chan reindeer)
	dnchn := make(chan struct{})
	go func() {
		for f := range ftrchn {
			future = append(future, f)
		}
		dnchn <- struct{}{}
	}()
	for _, rnd := range remaining {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ftr := rnd.GenerateFuturePossibleReindeer(maze)
			for _, f := range ftr {
				ftrchn <- f
			}
		}()
	}
	wg.Wait()
	close(ftrchn)
	<-dnchn
	// take the lowest from flock for each position
	// - basically, if we end up in the same place from two reindeer, only
	// the lowest survives
	fastest := []reindeer{}
	for _, rnd := range future {
		other := slices.IndexFunc(fastest, func(f reindeer) bool {
			return (f.Position[0] == rnd.Position[0] && f.Position[1] == rnd.Position[1]) || slices.ContainsFunc(f.History, func(hst []int) bool {
				return hst[0] == rnd.Position[0] && hst[1] == rnd.Position[1]
			})
		})
		if other == -1 {
			fastest = append(fastest, rnd)
		} else if fastest[other].Score > rnd.Score {
			fastest[other] = rnd
		}
	}
	return recurrFlock(maze, exit, fastest, current_low_score)
}

func ComputeSolutionOne(data []byte) int64 {
	maze, start, exit := parseMaze(data)

	current_low_score := 9999999999999

	flock := []reindeer{NewReindeer(start, []int{1, 0}, 0, 1000, 1, [][]int{start})}

	low_score := reindeer{Score: current_low_score}
	for range 10 {
		poss_low_score := recurrFlock(maze, exit, flock, reindeer{Score: current_low_score})
		if poss_low_score.Score < low_score.Score {
			low_score = poss_low_score
		}
	}

	return int64(low_score.Score)
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

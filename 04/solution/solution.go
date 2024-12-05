package solution

import (
	"fmt"
	"slices"
	"strings"
)

func buildGrid(data []byte) ([][]rune, int, int) {
	str := string(data)
	lines := strings.Split(str, "\n")
	height := len(lines)
	width := len(lines[0])
	grid := [][]rune{}
	for _, ln := range lines {
		grid = append(grid, []rune(ln))
	}
	return grid, width, height
}

type wordInstance struct {
	StartingCoord []int
	Cardinality   []int
}

func getStartingLocations(firstchar rune, grid [][]rune) []wordInstance {
	starts := []wordInstance{}
	for y, ln := range grid {
		for x, rn := range ln {
			if rn == firstchar {
				for _, xc := range []int{-1, 0, 1} {
					for _, yc := range []int{-1, 0, 1} {
						if xc == 0 && yc == 0 {
							continue
						}
						starts = append(starts, wordInstance{
							StartingCoord: []int{x, y},
							Cardinality:   []int{xc, yc},
						})
					}
				}
			}
		}
	}
	return starts
}

func getWordInstances(word string, grid [][]rune, width int, height int) []wordInstance {
	runeWord := []rune(word)
	starts := getStartingLocations(runeWord[0], grid)

	return findWordInstances(runeWord, starts, grid, width, height, 1)
}

func findWordInstances(word []rune, continuewith []wordInstance, grid [][]rune, width int, height int, rindx int) []wordInstance {
	if rindx >= len(word) {
		// recursion done
		return continuewith
	}
	cnt := []wordInstance{}
	for _, wi := range continuewith {
		nxtx := wi.StartingCoord[0] + (wi.Cardinality[0] * rindx)
		nxty := wi.StartingCoord[1] + (wi.Cardinality[1] * rindx)
		if nxtx < 0 || nxtx >= width {
			continue
		}
		if nxty < 0 || nxty >= height {
			continue
		}
		if word[rindx] == grid[nxty][nxtx] {
			cnt = append(cnt, wi)
		}
	}
	return findWordInstances(word, cnt, grid, width, height, rindx+1)
}

func getStartingLocationsMiddles(firstchar rune, grid [][]rune) []wordInstance {
	starts := []wordInstance{}
	for y, ln := range grid {
		for x, rn := range ln {
			if rn == firstchar {
				for _, xc := range []int{-1, 0, 1} {
					for _, yc := range []int{-1, 0, 1} {
						if xc == 0 || yc == 0 {
							continue
						}
						starts = append(starts, wordInstance{
							StartingCoord: []int{x, y},
							Cardinality:   []int{xc, yc},
						})
					}
				}
			}
		}
	}
	return starts
}

// no recursion because we only need to go one step
func findWordInstancesBiDi(word []rune, continuewith []wordInstance, grid [][]rune, width int, height int) []wordInstance {

	cnt := []wordInstance{}
	for _, wi := range continuewith {
		nxtx := wi.StartingCoord[0] + (wi.Cardinality[0] * 1)
		nxty := wi.StartingCoord[1] + (wi.Cardinality[1] * 1)
		if nxtx < 0 || nxtx >= width {
			continue
		}
		if nxty < 0 || nxty >= height {
			continue
		}
		if word[0] == grid[nxty][nxtx] {
			cnt = append(cnt, wi)
		}
	}
	cnt2 := []wordInstance{}
	for _, wi := range cnt {
		nxtx := wi.StartingCoord[0] + (wi.Cardinality[0] * -1)
		nxty := wi.StartingCoord[1] + (wi.Cardinality[1] * -1)
		if nxtx < 0 || nxtx >= width {
			continue
		}
		if nxty < 0 || nxty >= height {
			continue
		}
		if word[2] == grid[nxty][nxtx] {
			cnt2 = append(cnt2, wi)
		}
	}
	return cnt2
}

func getWordInstancesInCross(word string, grid [][]rune, width int, height int) []wordInstance {
	runeWord := []rune(word)
	starts := getStartingLocationsMiddles(runeWord[1], grid) // this is bad but we do it anyway

	allmasses := findWordInstancesBiDi(runeWord, starts, grid, width, height) // this is also bad an unextensible
	hascross := []wordInstance{}
	appeared := []string{}
	for _, am := range allmasses {
		if !slices.Contains(appeared, fmt.Sprintf("%d,%d", am.StartingCoord[0], am.StartingCoord[1])) {
			appeared = append(appeared, fmt.Sprintf("%d,%d", am.StartingCoord[0], am.StartingCoord[1]))
		} else {
			hascross = append(hascross, am)
		}
	}
	return hascross
}

func ComputeSolutionOne(data []byte) int64 {
	grid, width, height := buildGrid(data)
	foundInstances := getWordInstances("XMAS", grid, width, height)
	return int64(len(foundInstances))
}

func ComputeSolutionTwo(data []byte) int64 {
	grid, width, height := buildGrid(data)
	foundInstances := getWordInstancesInCross("MAS", grid, width, height)
	return int64(len(foundInstances))
}

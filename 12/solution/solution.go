package solution

import (
	"strings"
)

type Region struct {
	Coordinates [][]int
	Rune        rune
}

func (r Region) GetArea() int64 {
	return int64(len(r.Coordinates))
}

func (r Region) GetPerimeter() int64 {
	// early lazy exits
	if len(r.Coordinates) == 1 {
		return 4
	}
	if len(r.Coordinates) == 2 {
		return 6
	}
	acc := int64(0)

	crdmp := map[int]map[int]bool{}
	// this is dumb, I don't care
	for _, crd := range r.Coordinates {
		crdmp[crd[0]] = map[int]bool{}
	}
	for _, crd := range r.Coordinates {
		crdmp[crd[0]][crd[1]] = true
	}

	for _, crd := range r.Coordinates {
		tlt := int64(4)
		// for each neighbouring cord, subtract one
		chk := [][]int{
			{crd[0] - 1, crd[1]},
			{crd[0] + 1, crd[1]},
			{crd[0], crd[1] - 1},
			{crd[0], crd[1] + 1},
		}
		for _, ch := range chk {
			if crdmp[ch[0]][ch[1]] {
				tlt -= 1
			}
		}
		acc += tlt
	}
	return acc
}

func (r Region) GetPrice() int64 {
	return r.GetArea() * r.GetPerimeter()
}

func ParseRegions(data []byte) []Region {
	grid := map[int]map[int]rune{}
	visited := map[int]map[int]bool{}
	for y, ln := range strings.Split(string(data), "\n") {
		if y == 0 {
			for x := range ln {
				grid[x] = map[int]rune{}
				visited[x] = map[int]bool{}
			}
		}
		for x, rn := range ln {
			grid[x][y] = rn
		}
	}

	acc := []Region{}
	for x, col := range grid {
		for y, rn := range col {
			if visited[x][y] {
				//skip, as we've been here
				continue
			}
			visited[x][y] = true
			// do a fill then create a region from it.
			newrg := Region{
				Coordinates: [][]int{{x, y}},
				Rune:        rn,
			}
			// check surrounding cords
			cont := true
			last_cords := [][]int{{x, y}}
			for cont {
				cont = false
				new_cords := [][]int{}

				for _, crd := range last_cords {
					chk := [][]int{
						{crd[0] - 1, crd[1]},
						{crd[0] + 1, crd[1]},
						{crd[0], crd[1] - 1},
						{crd[0], crd[1] + 1},
					}
					for _, ch := range chk {
						if grid[ch[0]][ch[1]] == rn && !visited[ch[0]][ch[1]] {
							visited[ch[0]][ch[1]] = true
							new_cords = append(new_cords, ch)
						}
					}
				}

				if len(new_cords) > 0 {
					last_cords = new_cords
					newrg.Coordinates = append(newrg.Coordinates, new_cords...)
					cont = true
				}
			}
			acc = append(acc, newrg)
		}
	}
	return acc
}

func ComputeSolutionOne(data []byte) int64 {
	rgns := ParseRegions(data)
	acc := int64(0)
	for _, rgn := range rgns {
		acc += rgn.GetPrice()
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

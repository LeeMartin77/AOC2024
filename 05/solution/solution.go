package solution

import (
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	MustFollow  []int64
	MustPreceed []int64
}

type RulesEngine struct {
	Rules map[int64]Rule
}

type PrintingInstruction struct {
	Slice    []int64
	OrderMap map[int64]*int64
}

// add the middle pages of all "correct" rows together
// X|Y - X must be printed before Y

func parseRawData(data []byte) (*RulesEngine, []PrintingInstruction) {
	str := string(data)
	lines := strings.Split(str, "\n")
	re := RulesEngine{Rules: map[int64]Rule{}}
	pi := []PrintingInstruction{}
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		if strings.Contains(ln, "|") {
			// rule
			prts := strings.Split(ln, "|")
			X, _ := strconv.ParseInt(prts[0], 10, 64)
			Y, _ := strconv.ParseInt(prts[1], 10, 64)
			if entry, ok := re.Rules[Y]; ok {
				entry.MustFollow = append(entry.MustFollow, X)
				re.Rules[Y] = entry
			} else {
				re.Rules[Y] = Rule{
					MustFollow: []int64{X},
				}
			}
			if entry, ok := re.Rules[X]; ok {
				entry.MustPreceed = append(entry.MustPreceed, Y)
				re.Rules[X] = entry
			} else {
				re.Rules[X] = Rule{
					MustPreceed: []int64{X},
				}
			}
		} else {
			// instruction
			pgs := strings.Split(ln, ",")
			pii := PrintingInstruction{OrderMap: map[int64]*int64{}}
			for i, pg := range pgs {
				pgnum, _ := strconv.ParseInt(pg, 10, 64)
				pii.Slice = append(pii.Slice, pgnum)
				ii := int64(i)
				pii.OrderMap[pgnum] = &ii
			}
			pi = append(pi, pii)
		}
	}
	return &re, pi
}

func (re *RulesEngine) IsIntructionValid(pi PrintingInstruction) bool {
	for ord, pg := range pi.Slice {
		or := int64(ord)
		for _, prc := range re.Rules[pg].MustFollow {
			if pi.OrderMap[prc] != nil && *pi.OrderMap[prc] > or {
				return false
			}
		}
		for _, prc := range re.Rules[pg].MustPreceed {
			if pi.OrderMap[prc] != nil && *pi.OrderMap[prc] < or {
				return false
			}
		}
	}
	return true
}
func (re *RulesEngine) FixInstruction(pi PrintingInstruction) []int64 {
	sl := pi.Slice
	slices.SortFunc(sl, func(fpg int64, fgp2 int64) int {
		for _, prc := range re.Rules[fpg].MustPreceed {
			if prc == fgp2 {
				return 1
			}
		}
		for _, prc := range re.Rules[fpg].MustFollow {
			if prc == fgp2 {
				return -1
			}
		}
		return 0
	})
	return sl
}

func ComputeSolutionOne(data []byte) int64 {
	re, pi := parseRawData(data)
	acc := int64(0)
	for _, p := range pi {
		if re.IsIntructionValid(p) {
			acc += p.Slice[len(p.Slice)/2]
		}
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	re, pi := parseRawData(data)
	acc := int64(0)
	for _, p := range pi {
		if !re.IsIntructionValid(p) {
			corrected := re.FixInstruction(p)
			acc += corrected[len(corrected)/2]
		}
	}
	return acc
}

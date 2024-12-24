package solution

import (
	"slices"
	"strconv"
	"strings"
)

type Wire struct {
	State *bool
}

type Gate struct {
	Left  *Wire
	Right *Wire
	Op    string
	Out   map[string]*Wire
	Fired bool
}

func parse(data []byte) (map[string]*Wire, map[string]*Gate, []string) {
	inputs := true

	wiring := map[string]*Wire{}
	gates := map[string]*Gate{}
	hot := []string{}

	for _, ln := range strings.Split(string(data), "\n") {
		if ln == "" {
			inputs = false
			continue
		}
		if inputs {
			prts := strings.Split(ln, ": ")
			live := prts[1] == "1"
			wiring[prts[0]] = &Wire{
				State: &live,
			}
			hot = append(hot, prts[0])
		} else {
			prts := strings.Split(ln, " ")
			if wiring[prts[0]] == nil {
				wiring[prts[0]] = &Wire{}
			}
			if wiring[prts[2]] == nil {
				wiring[prts[2]] = &Wire{}
			}

			if wiring[prts[4]] == nil {
				wiring[prts[4]] = &Wire{}
			}
			gate_key := prts[0] + ":" + prts[1] + ":" + prts[2]
			if gates[gate_key] == nil {
				gates[gate_key] = &Gate{
					Left:  wiring[prts[0]],
					Right: wiring[prts[2]],
					Out:   map[string]*Wire{prts[4]: wiring[prts[4]]},
					Op:    prts[1],
					Fired: false,
				}
			} else {
				gates[gate_key].Out[prts[4]] = wiring[prts[4]]
			}
		}
	}

	return wiring, gates, hot
}

func (gt *Gate) Fire() []string {
	hot := true
	cold := false
	heated := []string{}
	for wr, out := range gt.Out {
		heated = append(heated, wr)
		switch gt.Op {
		case "AND":
			if *gt.Left.State && *gt.Right.State {
				out.State = &hot
			} else {
				out.State = &cold
			}
		case "XOR":
			if *gt.Left.State != *gt.Right.State {
				out.State = &hot
			} else {
				out.State = &cold
			}
		case "OR":
			if *gt.Left.State || *gt.Right.State {
				out.State = &hot
			} else {
				out.State = &cold
			}
		}
	}
	return heated
}

func ComputeSolutionOne(data []byte) int64 {
	wiring, gates, hot := parse(data)
	for {
		gates_to_fire := map[string]*Gate{}
		for id, gt := range gates {
			for _, lt := range hot {
				for _, rt := range hot[1:] {
					if gt.Left == wiring[lt] && gt.Right == wiring[rt] {
						gates_to_fire[id] = gt
					} else if gt.Left == wiring[rt] && gt.Right == wiring[lt] {
						gates_to_fire[id] = gt
					}
				}
			}
		}
		new_hot := []string{}
		for id, gt := range gates_to_fire {
			new_hot = append(new_hot, gt.Fire()...)
			delete(gates, id)
		}
		hot = append(hot, new_hot...)
		if len(gates) == 0 {
			break
		}
	}
	zds := []struct {
		id string
		bl bool
	}{}
	for id, str := range wiring {
		if id[0] == 'z' && str.State != nil {
			zds = append(zds, struct {
				id string
				bl bool
			}{id, *str.State})
		}
	}
	slices.SortFunc(zds, func(a, b struct {
		id string
		bl bool
	}) int {
		return strings.Compare(b.id, a.id)
	})
	str := ""
	for _, bl := range zds {
		if bl.bl {
			str += "1"
		} else {
			str += "0"
		}
	}
	res, _ := strconv.ParseInt(str, 2, 64)
	return res
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}

package solution

import (
	"fmt"
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
		new_hot = append(hot, new_hot...)
		hot = []string{}
	outer:
		for _, ht := range new_hot {
			for _, gt := range gates {
				if gt.Left == wiring[ht] || gt.Right == wiring[ht] {
					hot = append(hot, ht)
					continue outer
				}
			}
		}
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

func compute(wiring map[string]*Wire, gates map[string]*Gate, hot []string, tgt int64) error {
	// our "out register" then is all the Z wires
	zds := []struct {
		id string
		bl *Wire
	}{}
	for id, str := range wiring {
		if id[0] == 'z' {
			zds = append(zds, struct {
				id string
				bl *Wire
			}{id, str})
		}
	}
	slices.SortFunc(zds, func(a, b struct {
		id string
		bl *Wire
	}) int {
		return strings.Compare(b.id, a.id)
	})
	out_register := []*Wire{}
	for _, z := range zds {
		out_register = append(out_register, z.bl)
	}

	// we then turn tgt into it's bitmask
	bmsk := strconv.FormatInt(tgt, 2)
	expct := []bool{}
	for _, rn := range bmsk {
		expct = append(expct, rn == '1')
	}

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
		new_hot = append(hot, new_hot...)
		hot = []string{}
	outer:
		for _, ht := range new_hot {
			for _, gt := range gates {
				if gt.Left == wiring[ht] || gt.Right == wiring[ht] {
					hot = append(hot, ht)
					continue outer
				}
			}
		}
		if len(gates) == 0 {
			break
		}
	}

	output := ""
	for _, oo := range out_register {
		if oo.State == nil {
			output += "X"
			continue
		}
		if *oo.State {

			output += "1"
		} else {

			output += "0"
		}
	}
	nmsm := ""
	hasError := false
	for i, o := range out_register {
		if o.State == nil {
			nmsm += "_"
			hasError = true
		} else if *o.State == expct[i] {
			nmsm += "O"
		} else {
			nmsm += "X"
			hasError = true
		}
	}
	if hasError {
		return fmt.Errorf("bmk: %s\nout: %s\nmsm: %s", bmsk, output, nmsm)
	}
	return nil
}

func ComputeSolutionTwo(data []byte, swaps int) string {
	// We want an "out register"

	// get the initial x/y values, and add them - that's our target
	wiring_init, _, _ := parse(append([]byte{}, data...))
	x, y, tgt := int64(0), int64(0), int64(0)

	xds := []struct {
		id string
		bl bool
	}{}
	yds := []struct {
		id string
		bl bool
	}{}
	for id, str := range wiring_init {
		if id[0] == 'x' && str.State != nil {
			xds = append(xds, struct {
				id string
				bl bool
			}{id, *str.State})
		}
		if id[0] == 'y' && str.State != nil {
			yds = append(yds, struct {
				id string
				bl bool
			}{id, *str.State})
		}
	}
	slices.SortFunc(xds, func(a, b struct {
		id string
		bl bool
	}) int {
		return strings.Compare(a.id, b.id)
	})
	slices.SortFunc(yds, func(a, b struct {
		id string
		bl bool
	}) int {
		return strings.Compare(a.id, b.id)
	})
	str := ""
	for _, bl := range xds {
		if bl.bl {
			str += "1"
		} else {
			str += "0"
		}
	}
	x, _ = strconv.ParseInt(str, 2, 64)

	str = ""
	for _, bl := range yds {
		if bl.bl {
			str += "1"
		} else {
			str += "0"
		}
	}
	y, _ = strconv.ParseInt(str, 2, 64)

	tgt = x + y

	fmt.Printf("X: %d\n", x)
	fmt.Printf("Y: %d\n", y)
	fmt.Printf("T: %d\n", tgt)

	result := make(chan []string)
	go func() {
		i := 0
		swps := []string{}
		for {
			i += 1
			ii := i
			wiring, gates, hot := parse(append([]byte{}, data...))

			err := compute(wiring, gates, hot, tgt)
			if err == nil {
				result <- swps
			} else {
				fmt.Printf("%d: \n%s\n", ii, err)
			}
		}
	}()
	for res := range result {
		// we only want the first anyway
		slices.Sort(res)
		return strings.Join(res, ",")
	}
	panic("supposedly impossible")
}

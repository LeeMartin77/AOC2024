package solution

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Computer struct {
	RegisterA int64
	RegisterB int64
	RegisterC int64

	Ptr int

	Output []int64
}

func ParseInput(data []byte) (*Computer, []int8) {
	cmp := Computer{
		Ptr:    0,
		Output: []int64{},
	}
	prg := []int8{}
	for x, ln := range strings.Split(string(data), "\n") {
		prts := strings.Split(ln, " ")
		switch x {
		case 0:
			ra, _ := strconv.ParseInt(prts[2], 10, 64)
			cmp.RegisterA = ra
		case 1:
			ra, _ := strconv.ParseInt(prts[2], 10, 64)
			cmp.RegisterB = ra
		case 2:
			ra, _ := strconv.ParseInt(prts[2], 10, 64)
			cmp.RegisterC = ra
		case 4:
			inps := strings.Split(prts[1], ",")
			for _, in := range inps {
				i, _ := strconv.ParseInt(in, 10, 8)
				prg = append(prg, int8(i))
			}
		}
	}
	return &cmp, prg
}

type previous struct {
	cmp Computer
	op  int8
}

func (cmp *Computer) CloneComputer() Computer {
	return Computer{
		RegisterA: cmp.RegisterA,
		RegisterB: cmp.RegisterB,
		RegisterC: cmp.RegisterC,

		Ptr: cmp.Ptr,
	}
}

func (cmp *Computer) RunProgram(prg []int8) bool {
	// TODO: Check for "infinite loop"
	seenCombos := map[int8][]previous{}
	for {
		if cmp.Ptr+1 > len(prg)-1 {
			// halt
			break
		}
		inst, op := prg[cmp.Ptr], prg[cmp.Ptr+1]
		if slices.ContainsFunc(seenCombos[inst], func(prv previous) bool {
			return prv.op == op &&
				prv.cmp.RegisterA == cmp.RegisterA &&
				prv.cmp.RegisterB == cmp.RegisterB &&
				prv.cmp.RegisterC == cmp.RegisterC &&
				prv.cmp.Ptr == cmp.Ptr
		}) {
			return false
		} else {
			seenCombos[inst] = append(seenCombos[inst], previous{
				cmp.CloneComputer(), op,
			})
		}
		switch inst {
		case 0: //cv
			cmp.Adv(op)
		case 1: //cv
			cmp.Bxl(op)
		case 2: //cv
			cmp.Bst(op)
		case 3: //cv
			cmp.Jnz(op)
		case 4: //cv
			cmp.Bxc(op)
		case 5: //cv
			cmp.Out(op)
		case 6: //ncv
			cmp.Bdv(op)
		case 7: //ncv
			cmp.Cdv(op)
		}
		cmp.Ptr += 2
	}
	return true
}

func (cmp *Computer) ComboOp(op int8) int64 {
	if op < 4 {
		return int64(op)
	}
	switch op {
	case 4:
		return cmp.RegisterA
	case 5:
		return cmp.RegisterB
	case 6:
		return cmp.RegisterC
	}
	panic("Valid programs don't do this")
}

func (cmp *Computer) Adv(op int8) {
	out := float64(cmp.RegisterA) / math.Pow(2, float64(cmp.ComboOp(op)))
	// we make this gnarly
	str := strings.Split(fmt.Sprintf("%f", out), ".")[0]
	prs, _ := strconv.ParseInt(str, 10, 64)
	cmp.RegisterA = prs
}
func (cmp *Computer) Bxl(op int8) {
	cmp.RegisterB = cmp.RegisterB ^ int64(op)
}
func (cmp *Computer) Bst(op int8) {
	cmp.RegisterB = cmp.ComboOp(op) % 8
}
func (cmp *Computer) Jnz(op int8) {
	if cmp.RegisterA == 0 {
		//nop
		return
	}
	cmp.Ptr = int(op) - 2
}
func (cmp *Computer) Bxc(_ int8) {
	// FIXME
	cmp.RegisterB = cmp.RegisterB ^ cmp.RegisterC
}
func (cmp *Computer) Out(op int8) {
	cmp.Output = append(cmp.Output, cmp.ComboOp(op)%8)
}
func (cmp *Computer) Bdv(op int8) {
	out := float64(cmp.RegisterA) / math.Pow(2, float64(cmp.ComboOp(op)))
	// we make this gnarly
	str := strings.Split(fmt.Sprintf("%f", out), ".")[0]
	prs, _ := strconv.ParseInt(str, 10, 64)
	cmp.RegisterB = prs
}
func (cmp *Computer) Cdv(op int8) {
	out := float64(cmp.RegisterA) / math.Pow(2, float64(cmp.ComboOp(op)))
	// we make this gnarly
	str := strings.Split(fmt.Sprintf("%f", out), ".")[0]
	prs, _ := strconv.ParseInt(str, 10, 64)
	cmp.RegisterC = prs
}

func (cmp *Computer) PrintOutput() string {
	str := []string{}
	for _, v := range cmp.Output {
		str = append(str, fmt.Sprintf("%d", v))
	}
	return strings.Join(str, ",")
}

func PrintProgram(prg []int8) string {
	str := []string{}
	for _, v := range prg {
		str = append(str, fmt.Sprintf("%d", v))
	}
	return strings.Join(str, ",")
}

func ComputeSolutionOne(data []byte) string {
	cmp, prg := ParseInput(data)
	cmp.RunProgram(prg)
	return cmp.PrintOutput()
}

func ComputeSolutionTwo(data []byte) int64 {

	cmp, prg := ParseInput(data)
	prnt := PrintProgram(prg)
	solution := int64(0)
	attempt := int64(0)
	solchan := make(chan int64)
	soldone := make(chan struct{})
	ch := make(chan struct{}, 1000000)
	wg := sync.WaitGroup{}
	go func() {
		for sol := range solchan {
			if solution == 0 || sol < solution {
				solution = sol
			}
		}
		soldone <- struct{}{}
	}()
	for solution == 0 {
		wg.Add(1)
		ch <- struct{}{}
		go func(atmp int64, tstcmp Computer) {
			defer wg.Done()
			tstcmp.RegisterA = atmp
			halted := tstcmp.RunProgram(prg)
			if !halted {
				fmt.Printf("%d:false\n", atmp)
				<-ch
				return
			}
			out := tstcmp.PrintOutput()
			//fmt.Printf("%d:true\n", attempt)
			if out == prnt {
				solchan <- atmp
			}
			<-ch
		}(attempt, cmp.CloneComputer())
		attempt += 1
	}
	wg.Wait()
	close(solchan)
	<-soldone
	return solution
}

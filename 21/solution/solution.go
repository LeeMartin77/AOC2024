package solution

import (
	"fmt"
	"strconv"
	"strings"
)

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A*|
//     +---+---+

// type numberpad
// input -> command sequence
var keymap map[rune][]int = map[rune][]int{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'E': {0, 3},
	'0': {1, 3},
	'A': {2, 3},
}

func TypeNumberPad(cmd string) string {
	return typeIntoPad(cmd, keymap, []int{keymap['A'][0], keymap['A'][1]}, nil, "")

}

// type commandpad
// input comseq -> command sequence

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

var cmdmap map[rune][]int = map[rune][]int{
	'E': {0, 0},
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

type ComCache struct {
	cache map[string]*string
}

func (cc *ComCache) TypeCommmandPad(cmd string) string {
	return typeIntoPad(cmd, cmdmap, []int{cmdmap['A'][0], cmdmap['A'][1]}, cc.cache, "")
}

func typeIntoPad(cmdstr string, keyp map[rune][]int, pos []int, cache map[string]*string, retval string) string {

	cache_str := cmdstr + fmt.Sprintf(":%d,%d", pos[0], pos[1])
	if cache != nil && cache[cache_str] != nil {
		return retval + *cache[cache_str]
	}
	var cmd byte

	if len(cmdstr) == 0 {
		return retval
	}

	if len(cmdstr) == 1 {
		cmd = cmdstr[0]
		cmdstr = ""
	} else {
		cmd, cmdstr = cmdstr[0], cmdstr[1:]
	}
	target_pos := keyp[rune(cmd)]

	xadj := pos[0] - target_pos[0]
	yadj := pos[1] - target_pos[1]

	// this is for efficient operation orders
	// and avoiding going over the error key
	vert_first := false
	if yadj > 0 && !(keyp['E'][0] == target_pos[0] && keyp['E'][1] == pos[1]) {
		vert_first = false
	} else if !(keyp['E'][0] == pos[0] && keyp['E'][1] == target_pos[1]) {
		vert_first = true
	}

	vert := ""
	horiz := ""
	for xadj != 0 {
		if xadj > 0 {
			xadj -= 1
			pos[0] -= 1
			horiz += "<"
		} else if xadj < 0 {
			xadj += 1
			pos[0] += 1
			horiz += ">"
		}
	}
	for yadj != 0 {
		if yadj > 0 {
			yadj -= 1
			pos[1] -= 1
			vert += "^"
		} else if yadj < 0 {
			yadj += 1
			pos[1] += 1
			vert += "v"
		}
	}
	new_com := ""
	if vert_first {
		new_com += vert + horiz + "A"
	} else {
		new_com += horiz + vert + "A"
	}
	retval += new_com

	final := typeIntoPad(cmdstr, keyp, pos, cache, retval)
	if cache != nil {
		cache[cache_str] = &final
	}
	return final
}

func GetNumber(cmd string) int64 {
	vl, _ := strconv.ParseInt(cmd[:len(cmd)-1], 10, 16)
	return vl
}

func GoThroughRobotsAndGetComplexity(cmd string, num_bots int) int64 {
	comput := TypeNumberPad(cmd)
	cc := ComCache{
		cache: map[string]*string{},
	}
	comput = cc.TypeCommmandPad(comput)
	commcache := ComCache{
		cache: map[string]*string{},
	}
	for range num_bots - 1 {
		comput = commcache.TypeCommmandPad(comput)
	}
	return GetNumber(cmd) * int64(len(comput))
}

func ComputeSolutionOne(data []byte) int64 {
	inpts := strings.Split(string(data), "\n")
	acc := int64(0)
	for _, inp := range inpts {
		// we'll just do it in the loop for P1
		acc += GoThroughRobotsAndGetComplexity(inp, 2)
	}
	return acc
}

func ComputeSolutionTwo(data []byte) int64 {
	inpts := strings.Split(string(data), "\n")
	acc := int64(0)
	for _, inp := range inpts {
		// we'll just do it in the loop for P1
		acc += GoThroughRobotsAndGetComplexity(inp, 25)
	}
	return acc
}

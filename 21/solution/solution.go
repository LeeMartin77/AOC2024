package solution

import (
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
	return typeIntoPad(cmd, keymap, []int{keymap['A'][0], keymap['A'][1]}, "")

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

func TypeCommmandPad(cmd string) string {
	return typeIntoPad(cmd, cmdmap, []int{cmdmap['A'][0], cmdmap['A'][1]}, "")
}

func getCommandSequence(cmd rune, keyp map[rune][]int, pos []int) string {
	target_pos := keyp[cmd]

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
	return new_com
}

func typeIntoPad(cmdstr string, keyp map[rune][]int, pos []int, retval string) string {

	if len(cmdstr) == 0 {
		return retval
	}
	var cmd byte
	if len(cmdstr) == 1 {
		cmd = cmdstr[0]
		cmdstr = ""
	} else {
		cmd, cmdstr = cmdstr[0], cmdstr[1:]
	}

	retval += getCommandSequence(rune(cmd), keyp, pos)
	return typeIntoPad(cmdstr, keyp, pos, retval)
}

func GetNumber(cmd string) int64 {
	vl, _ := strconv.ParseInt(cmd[:len(cmd)-1], 10, 16)
	return vl
}

type Memo struct {
	seq   string
	depth int
}

var dp = map[Memo]int{} // Memoization map

func dfs(memo Memo) int {
	if v, ok := dp[memo]; ok {
		return v
	}
	if memo.depth == 0 {
		return len(memo.seq)
	}

	path := TypeCommmandPad(memo.seq)
	res := 0
	for _, p := range path {
		res += dfs(Memo{string(p), memo.depth - 1})
	}
	dp[memo] = res
	return res
}
func GoThroughRobotsAndGetComplexityDfs(cmd string, num_bots int) int64 {
	comput := TypeNumberPad(cmd)
	acc := 0
	for _, cm := range comput {
		acc += dfs(Memo{string(cm), num_bots})
	}
	return GetNumber(cmd) * int64(len(comput)+acc)
}

func GoThroughRobotsAndGetComplexity(cmd string, num_bots int) int64 {
	comput := TypeNumberPad(cmd)
	for range num_bots {
		comput = TypeCommmandPad(comput)
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
		acc += GoThroughRobotsAndGetComplexityDfs(inp, 25)
	}
	return acc
}

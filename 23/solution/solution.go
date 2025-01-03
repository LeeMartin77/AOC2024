package solution

import (
	"slices"
	"strings"
)

func parseComputers(data []byte) map[string][]string {
	connection_map := map[string][]string{}
	for _, ln := range strings.Split(string(data), "\n") {
		cmps := strings.Split(ln, "-")
		cmpa, cmpb := cmps[0], cmps[1]
		connection_map[cmpa] = append(connection_map[cmpa], cmpb)
		connection_map[cmpb] = append(connection_map[cmpb], cmpa)
	}
	return connection_map
}

func ComputeSolutionOne(data []byte) int64 {
	connection_map := parseComputers(data)
	// we are trying to find triangles...
	triangles := map[string]bool{}
	for cmp, sbnt := range connection_map {
		// triangle
		for _, sb := range sbnt {
			for _, sbb := range connection_map[sb] {
				if slices.Contains(sbnt, sbb) {
					sbnd := []string{cmp, sb, sbb}
					slices.Sort(sbnd)
					triangles[strings.Join(sbnd, ",")] = true
				}
			}
		}
	}
	acc := int64(0)

	for k := range triangles {
		cmps := strings.Split(k, ",")
		if slices.ContainsFunc(cmps, func(c string) bool {
			return c[0] == 't'
		}) {
			acc += 1
		}
	}

	return acc
}

func buildClique(acc []string, con_map map[string][]string) []string {
	for _, cm := range acc {
		for _, npm := range con_map[cm] {
			if slices.Contains(acc, npm) {
				continue
			}
			if slices.ContainsFunc(acc, func(ac string) bool {
				return !slices.Contains(con_map[npm], ac)
			}) {
				continue
			}
			return buildClique(append(acc, npm), con_map)
		}
	}
	return acc
}

func ComputeSolutionTwo(data []byte) string {
	connection_map := parseComputers(data)
	sets := [][]string{}
	for cmp := range connection_map {
		sets = append(sets, buildClique([]string{cmp}, connection_map))
	}
	slices.SortFunc(sets, func(a, b []string) int {
		return len(b) - len(a)
	})
	slices.Sort(sets[0])
	return strings.Join(sets[0], ",")
}

package adventofcode2025

import (
	"strings"
)

type Day10Machine struct {
	target   uint64   // target pattern as bitmask (part 1)
	buttons  []uint64 // each button as bitmask of lights it toggles
	n        int      // number of lights
	joltage  []int    // joltage requirements (part 2)
	btnLists [][]int  // button indices as lists (part 2)
}

type Day10Puzzle struct {
	machines []Day10Machine
}

func NewDay10(lines []string) (Day10Puzzle, error) {
	machines := make([]Day10Machine, len(lines))
	for i, line := range lines {
		machines[i] = parseDay10Line(line)
	}
	return Day10Puzzle{machines: machines}, nil
}

func parseDay10Line(line string) Day10Machine {
	var m Day10Machine

	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	pattern := line[start+1 : end]
	m.n = len(pattern)
	for i, c := range pattern {
		if c == '#' {
			m.target |= 1 << i
		}
	}

	rest := line[end+1:]
	braceIdx := strings.Index(rest, "{")
	buttonsPart := rest
	joltPart := ""
	if braceIdx >= 0 {
		buttonsPart = rest[:braceIdx]
		joltPart = rest[braceIdx:]
	}

	for len(buttonsPart) > 0 {
		pStart := strings.Index(buttonsPart, "(")
		if pStart < 0 {
			break
		}
		pEnd := strings.Index(buttonsPart, ")")
		nums := buttonsPart[pStart+1 : pEnd]
		buttonsPart = buttonsPart[pEnd+1:]

		var button uint64
		var btnList []int
		parts := strings.Split(nums, ",")
		for _, p := range parts {
			var idx int
			for _, c := range p {
				if c >= '0' && c <= '9' {
					idx = idx*10 + int(c-'0')
				}
			}
			button |= 1 << idx
			btnList = append(btnList, idx)
		}
		m.buttons = append(m.buttons, button)
		m.btnLists = append(m.btnLists, btnList)
	}

	if joltPart != "" {
		jStart := strings.Index(joltPart, "{")
		jEnd := strings.Index(joltPart, "}")
		if jStart >= 0 && jEnd > jStart {
			nums := joltPart[jStart+1 : jEnd]
			parts := strings.Split(nums, ",")
			for _, p := range parts {
				var val int
				for _, c := range p {
					if c >= '0' && c <= '9' {
						val = val*10 + int(c-'0')
					}
				}
				m.joltage = append(m.joltage, val)
			}
		}
	}

	return m
}

func Day10(p Day10Puzzle, part1 bool) uint {
	if part1 {
		return day10Part1(p)
	}
	return day10Part2(p)
}

func day10Part1(p Day10Puzzle) uint {
	var total uint
	for _, m := range p.machines {
		total += minPresses(m)
	}
	return total
}

func minPresses(m Day10Machine) uint {
	target := m.target
	if len(m.buttons) == 0 {
		if target == 0 {
			return 0
		}
		return 0
	}

	visited := make(map[uint64]bool)
	queue := [][2]uint64{{0, 0}}
	visited[0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		state, presses := cur[0], cur[1]

		if state == target {
			return uint(presses)
		}

		for _, btn := range m.buttons {
			next := state ^ btn
			if !visited[next] {
				visited[next] = true
				queue = append(queue, [2]uint64{next, presses + 1})
			}
		}
	}

	return 0
}

func day10Part2(p Day10Puzzle) uint {
	var total uint
	for _, m := range p.machines {
		total += minPressesJoltage(m)
	}
	return total
}

func minPressesJoltage(m Day10Machine) uint {
	if len(m.joltage) == 0 {
		return 0
	}

	nCounters := len(m.joltage)
	nButtons := len(m.btnLists)

	// Build matrix A where A[counter][button] = 1 if button affects counter
	A := make([][]int, nCounters)
	for i := range A {
		A[i] = make([]int, nButtons)
	}
	for bi, bl := range m.btnLists {
		for _, idx := range bl {
			if idx < nCounters {
				A[idx][bi] = 1
			}
		}
	}

	_, minSum, ok := ILPSolve(A, m.joltage)
	if !ok {
		return 0
	}
	return uint(minSum)
}

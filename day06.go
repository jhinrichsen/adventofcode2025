package adventofcode2025

func Day06(lines []string, part1 bool) uint {
	if part1 {
		return day06Part1(lines)
	}
	return day06Part2(lines)
}

func day06Part1(lines []string) uint {
	h := len(lines)

	// track position in each row
	pos := make([]int, h)

	nextToken := func(row int) (start, end int, ok bool) {
		line := lines[row]
		p := pos[row]
		for p < len(line) && line[p] == ' ' {
			p++
		}
		if p >= len(line) {
			return 0, 0, false
		}
		start = p
		for p < len(line) && line[p] != ' ' {
			p++
		}
		pos[row] = p
		return start, p, true
	}

	var total uint
	for {
		// get operator from last row
		opStart, _, ok := nextToken(h - 1)
		if !ok {
			break
		}
		op := lines[h-1][opStart]

		var result uint
		if op == '*' {
			result = 1
		}

		for row := range h - 1 {
			start, end, ok := nextToken(row)
			if !ok {
				continue
			}
			var n uint
			for i := start; i < end; i++ {
				c := lines[row][i]
				if c >= '0' && c <= '9' {
					n = n*10 + uint(c-'0')
				}
			}
			if op == '+' {
				result += n
			} else {
				result *= n
			}
		}
		total += result
	}
	return total
}

func day06Part2(lines []string) uint {
	h := len(lines)
	w := 0
	for _, line := range lines {
		if len(line) > w {
			w = len(line)
		}
	}

	getChar := func(row, x int) byte {
		if x < len(lines[row]) {
			return lines[row][x]
		}
		return ' '
	}

	isSpaceCol := func(x int) bool {
		for row := range h {
			if getChar(row, x) != ' ' {
				return false
			}
		}
		return true
	}

	var total uint
	x := 0
	for x < w {
		for x < w && isSpaceCol(x) {
			x++
		}
		if x >= w {
			break
		}

		x0 := x
		for x < w && !isSpaceCol(x) {
			x++
		}
		x1 := x

		// find operator
		var op byte
		for i := x0; i < x1; i++ {
			c := getChar(h-1, i)
			if c == '+' || c == '*' {
				op = c
				break
			}
		}

		// process character columns right-to-left
		var result uint
		if op == '*' {
			result = 1
		}
		for charCol := x1 - 1; charCol >= x0; charCol-- {
			var n uint
			hasDigit := false
			for row := range h - 1 {
				c := getChar(row, charCol)
				if c >= '0' && c <= '9' {
					n = n*10 + uint(c-'0')
					hasDigit = true
				}
			}
			if hasDigit {
				if op == '+' {
					result += n
				} else {
					result *= n
				}
			}
		}
		total += result
	}
	return total
}

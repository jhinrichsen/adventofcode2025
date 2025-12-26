package adventofcode2025

func Day07(lines []string, part1 bool) uint {
	var sx int
	for x := range lines[0] {
		if lines[0][x] == 'S' {
			sx = x
			break
		}
	}

	w := len(lines[0])
	counts := make([]uint, w)
	next := make([]uint, w)
	counts[sx] = 1

	var splits uint
	for row := 1; row < len(lines); row++ {
		line := lines[row]
		for col := range w {
			next[col] = 0
		}
		for col := range w {
			if counts[col] == 0 {
				continue
			}
			if col < len(line) && line[col] == '^' {
				splits++
				if col > 0 {
					next[col-1] += counts[col]
				}
				if col+1 < w {
					next[col+1] += counts[col]
				}
			} else {
				next[col] += counts[col]
			}
		}
		counts, next = next, counts
	}

	if part1 {
		return splits
	}
	var total uint
	for _, c := range counts {
		total += c
	}
	return total
}

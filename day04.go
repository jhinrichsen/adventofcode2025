package adventofcode2025

func Day04(buf []byte, part1 bool) (uint, error) {
	var w int
	for i := range buf {
		if buf[i] == '\n' {
			w = i + 1 // include newline in width
			break
		}
	}
	h := len(buf) / w
	g := Grid{w, h}

	if part1 {
		var count uint
		for idx, neighbors := range g.C8Indices() {
			if buf[idx] == '\n' { // skip newline column
				continue
			}
			if buf[idx] != '@' {
				continue
			}
			var rolls uint
			for ni := range neighbors {
				if buf[ni] == '@' {
					rolls++
				}
			}
			if rolls < 4 {
				count++
			}
		}
		return count, nil
	}

	var total uint
	toRemove := make([]int, w*h)

	for {
		n := 0

		for idx, neighbors := range g.C8Indices() {
			if buf[idx] == '\n' { // skip newline column
				continue
			}
			if buf[idx] != '@' {
				continue
			}
			var rolls uint
			for ni := range neighbors {
				if buf[ni] == '@' {
					rolls++
				}
			}
			if rolls < 4 {
				toRemove[n] = idx
				n++
			}
		}

		if n == 0 {
			break
		}

		for i := range n {
			buf[toRemove[i]] = '.'
		}
		total += uint(n)
	}

	return total, nil
}

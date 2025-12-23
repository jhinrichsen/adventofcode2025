package adventofcode2025

func Day03(buf []byte, part1 bool) (uint, error) {
	var sum uint

	// we need to trigger on newline, but also ignore the last character for l
	// single pass processing is a bit tricky and error prone, so we split
	// processing into a line parser and a second valud parser

	start := 0
	for i := range buf {
		if buf[i] == '\n' {
			line := buf[start:i:i]

			var l, r byte
			n := len(line)
			if n > 0 {
				// Scan up to n-1 to leave room for r
				_ = line[n-1] // bounds check hint
				for j := 0; j < n-1; j++ {
					b := line[j]
					if b > l {
						l = b
						r = 0
					} else if b > r {
						r = b
					}
				}
				// Always consider last digit for r
				last := line[n-1]
				if last > r {
					r = last
				}

				sum += 10*uint(l-'0') + uint(r-'0')
			}
			start = i + 1
		}
	}

	return sum, nil
}

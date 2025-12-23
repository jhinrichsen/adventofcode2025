package adventofcode2025

func Day03(buf []byte, part1 bool) (uint, error) {
	var sum uint

	if part1 {
		// Optimized path for finding 2 largest digits
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
	} else {
		// Part 2: Keep 12 digits to maximize the number
		// Use monotonic stack to remove k digits greedily
		// Preallocate stack buffer to avoid per-line allocations
		stack := make([]byte, 0, 128)
		start := 0
		for i := range buf {
			if buf[i] == '\n' {
				line := buf[start:i:i]
				n := len(line)
				if n > 0 {
					value := maximizeNumberWithStack(line, 12, stack[:0])
					sum += value
				}
				start = i + 1
			}
		}
	}

	return sum, nil
}

// maximizeNumberWithStack removes (len(digits) - keep) digits to produce the maximum number
// Reuses the provided stack buffer to avoid allocations
func maximizeNumberWithStack(digits []byte, keep int, stack []byte) uint {
	n := len(digits)
	if n <= keep {
		// Already at or below target, convert to number
		var result uint
		for _, d := range digits {
			result = result*10 + uint(d-'0')
		}
		return result
	}

	toRemove := n - keep

	for _, digit := range digits {
		// Remove smaller digits from stack if next digit is larger
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}

	// If we still need to remove digits, remove from the end
	stack = stack[:len(stack)-toRemove]

	// Convert to number
	var result uint
	for _, d := range stack {
		result = result*10 + uint(d-'0')
	}
	return result
}

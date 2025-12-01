package adventofcode2025

import "fmt"

func Day01(buf []byte, part1 bool) (uint, error) {
	const (
		LEFT    = 'L'
		RIGHT   = 'R'
		RANGE   = 100 // 0..99
		NEWLINE = 0x0a
	)

	var count uint

	dial := 50 // initial position

	// conventional for loop because index is mutated inside the loop
	for i := 0; i < len(buf); i++ {
		// direction
		var sign int
		if buf[i] == LEFT {
			sign = -1
		} else if buf[i] == RIGHT {
			sign = 1
		} else {
			return 0, fmt.Errorf("Want %c or %c but got %c", LEFT, RIGHT, buf[i])
		}
		i++

		// number
		var n int
		for buf[i] != NEWLINE {
			n = 10*n + int(buf[i]-'0')
			i++
		}

		// turn, consider both left and right wrap, branchless
		dial = ((dial+sign*n)%RANGE + RANGE) % RANGE

		// count
		if dial == 0 {
			count++
		}
	}
	return count, nil
}

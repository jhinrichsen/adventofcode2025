package adventofcode2025

import (
	"strconv"
	"strings"
)

func Day01(buf []byte, part1 bool) (uint, error) {
	const (
		LEFT    = 'L'
		RIGHT   = 'R'
		RANGE   = 100 // 0..99
		NEWLINE = 0x0a
		BIT     = 2 // L and R differ in bit position 2
	)

	var count uint

	dial := 50 // initial position

	// conventional for loop because index is mutated inside the loop
	for i := 0; i < len(buf); i++ {
		// direction
		sign := (int(buf[i] & 2)) - 1
		i++

		// number
		var n int
		for buf[i] != NEWLINE {
			n = 10*n + int(buf[i]-'0')
			i++
		}

		// part2: count crossings during movement (O(1) instead of O(n))
		if !part1 {
			if sign == 1 { // RIGHT
				// Going right: count multiples of 100 in range (dial, dial+n]
				count += uint((dial+n)/RANGE - dial/RANGE)
			} else { // LEFT
				// Going left: count multiples of 100 in range [dial-n, dial)
				// Use floor division for negative numbers
				oldFloor := (dial - 1) / RANGE
				if dial-1 < 0 && (dial-1)%RANGE != 0 {
					oldFloor--
				}
				newFloor := (dial - n - 1) / RANGE
				if (dial-n-1) < 0 && (dial-n-1)%RANGE != 0 {
					newFloor--
				}
				count += uint(oldFloor - newFloor)
			}
		}

		// turn, consider both left and right wrap, branchless
		dial = ((dial+sign*n)%RANGE + RANGE) % RANGE

		// count final position (part1)
		if part1 && dial == 0 {
			count++
		}
	}
	return count, nil
}

func Day01BruteForce(buf []byte, part1 bool) (uint, error) {
	var count uint

	lines := strings.Split(string(buf), "\n")
	position := 50

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := 1
		if line[0] == 'L' {
			direction = -1
		}

		n, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		for range n {
			position += direction
			if position == -1 {
				position = 99
			} else if position == 100 {
				position = 0
			}
			if (!part1) && position == 0 {
				count++
			}
		}

		if part1 {
			if position == 0 {
				count++
			}
		}
	}
	return count, nil
}

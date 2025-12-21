package adventofcode2025

import (
	"fmt"
)

func Day02(buf []byte, part1 bool) (uint, error) {
	var sum, i uint
	num := func() (n uint) {
		for {
			n = 10*n + uint(buf[i]-'0')
			i++
			if buf[i] < '0' || buf[i] > '9' {
				return
			}
		}
	}

	var seq func(uint, uint) func(func(uint) bool)
	if part1 {
		seq = twinSeq
	} else {
		seq = repeatedPatternSeq
	}

	for {
		// start of range
		j := num()
		if buf[i] == '-' {
			i++
		} else {
			return 0, fmt.Errorf("want '-' at index %d but got %c", i, buf[i])
		}

		// end of range
		k := num()

		// process the range
		for n := range seq(j, k) {
			sum += n
		}

		// check for end of input or next pair
		if buf[i] == '\n' {
			break
		}
		if buf[i] == ',' {
			i++
		} else {
			return 0, fmt.Errorf("want ',' or newline at index %d but got %c", i, buf[i])
		}
	}
	return sum, nil
}

// iterate twin nibbles [j..k]
func twinSeq(j, k uint) func(func(uint) bool) {
	return func(yield func(uint) bool) {
		// Find the starting number of digits (must be even)
		startDigits := numberOfDigits(j)
		if startDigits%2 == 1 {
			startDigits++
		}

		endDigits := numberOfDigits(k)

		// Iterate through each even number of digits
		for numDigits := startDigits; numDigits <= endDigits; numDigits += 2 {
			halfDigits := numDigits / 2
			minHalf := pow(halfDigits - 1) // e.g., 10 for 2 digits, 100 for 3 digits
			maxHalf := pow(halfDigits) - 1 // e.g., 99 for 2 digits, 999 for 3 digits

			// For each possible half value, construct the twin
			for half := minHalf; half <= maxHalf; half++ {
				twinNum := half*pow(halfDigits) + half

				// Only yield if in range [j, k]
				if twinNum >= j && twinNum <= k {
					if !yield(twinNum) {
						return
					}
				} else if twinNum > k {
					return
				}
			}
		}
	}
}

func numberOfDigits(n uint) uint {
	var p uint
	for n > 0 {
		p++
		n /= 10
	}
	return p
}

func evenNumberOfDigits(n uint) bool {
	return numberOfDigits(n)%2 == 0
}

func twin(n uint) uint {
	a := numberOfDigits(n)
	a /= 2
	p := pow(a)
	n /= p
	n = n*p + n
	return n
}

// inverse function of numberOfDigits.
func pow(n uint) uint {
	p := uint(1)
	for range n {
		p *= 10
	}
	return p
}

// hasRepeatedPattern checks if n is made only of a pattern repeated at least twice
func hasRepeatedPattern(n uint) bool {
	digits := numberOfDigits(n)
	// Try all pattern lengths from 1 to digits/2
	for patternLen := uint(1); patternLen <= digits/2; patternLen++ {
		// Pattern length must divide total digits evenly
		if digits%patternLen != 0 {
			continue
		}

		// Extract the pattern (first patternLen digits)
		divisor := pow(digits - patternLen)
		pattern := n / divisor

		// Reconstruct number by repeating the pattern
		repetitions := digits / patternLen
		reconstructed := uint(0)
		for i := uint(0); i < repetitions; i++ {
			reconstructed = reconstructed*pow(patternLen) + pattern
		}

		if reconstructed == n {
			return true
		}
	}
	return false
}

// repeatedPatternSeq iterates over all numbers with repeated patterns in [j..k]
func repeatedPatternSeq(j, k uint) func(func(uint) bool) {
	return func(yield func(uint) bool) {
		for n := j; n <= k; n++ {
			if hasRepeatedPattern(n) {
				if !yield(n) {
					return
				}
			}
		}
	}
}

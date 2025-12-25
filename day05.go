package adventofcode2025

import "fmt"

type day05Range struct {
	min, max uint
}

type Day05Puzzle struct {
	merged      []day05Range // sorted, non-overlapping
	ingredients []uint
}

func NewDay05(lines []string) (Day05Puzzle, error) {
	p := Day05Puzzle{}
	var ranges []day05Range

	// scan ranges
	var i int
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		line := lines[i]
		var j int
		var min uint
		for ; j < len(line) && line[j] != '-'; j++ {
			min = min*10 + uint(line[j]-'0')
		}
		if j >= len(line) {
			return p, fmt.Errorf("bad range: %s", line)
		}
		j++ // skip '-'
		var max uint
		for ; j < len(line); j++ {
			max = max*10 + uint(line[j]-'0')
		}
		ranges = append(ranges, day05Range{min, max})
	}

	// scan ingredients
	for ; i < len(lines); i++ {
		line := lines[i]
		var n uint
		for j := 0; j < len(line); j++ {
			n = n*10 + uint(line[j]-'0')
		}
		p.ingredients = append(p.ingredients, n)
	}

	// sort and merge ranges
	p.merged = mergeDay05Ranges(ranges)
	return p, nil
}

func Day05(p Day05Puzzle, part1 bool) uint {
	if part1 {
		var fresh uint
		for _, ing := range p.ingredients {
			if inDay05MergedRanges(p.merged, ing) {
				fresh++
			}
		}
		return fresh
	}
	var total uint
	for _, r := range p.merged {
		total += r.max - r.min + 1
	}
	return total
}

func mergeDay05Ranges(ranges []day05Range) []day05Range {
	if len(ranges) == 0 {
		return nil
	}
	// insertion sort by min
	for i := 1; i < len(ranges); i++ {
		for j := i; j > 0 && ranges[j].min < ranges[j-1].min; j-- {
			ranges[j], ranges[j-1] = ranges[j-1], ranges[j]
		}
	}
	merged := make([]day05Range, 0, len(ranges))
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if current.max >= ranges[i].min || current.max+1 >= ranges[i].min {
			if ranges[i].max > current.max {
				current.max = ranges[i].max
			}
		} else {
			merged = append(merged, current)
			current = ranges[i]
		}
	}
	return append(merged, current)
}

func inDay05MergedRanges(merged []day05Range, val uint) bool {
	lo, hi := 0, len(merged)
	for lo < hi {
		mid := (lo + hi) / 2
		if merged[mid].min <= val {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo == 0 {
		return false
	}
	return val <= merged[lo-1].max
}

// --- brute force version: O(n*m) for benchmarking comparison ---

type day05BruteForcePuzzle struct {
	ranges      []day05Range
	ingredients []uint
}

func newDay05BruteForce(lines []string) (day05BruteForcePuzzle, error) {
	p := day05BruteForcePuzzle{}

	// scan ranges
	var i int
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		line := lines[i]
		var j int
		var min uint
		for ; j < len(line) && line[j] != '-'; j++ {
			min = min*10 + uint(line[j]-'0')
		}
		if j >= len(line) {
			return p, fmt.Errorf("bad range: %s", line)
		}
		j++ // skip '-'
		var max uint
		for ; j < len(line); j++ {
			max = max*10 + uint(line[j]-'0')
		}
		p.ranges = append(p.ranges, day05Range{min, max})
	}

	// scan ingredients
	for ; i < len(lines); i++ {
		line := lines[i]
		var n uint
		for j := 0; j < len(line); j++ {
			n = n*10 + uint(line[j]-'0')
		}
		p.ingredients = append(p.ingredients, n)
	}

	return p, nil
}

func day05BruteForce(p day05BruteForcePuzzle) uint {
	var fresh uint
	for i := range p.ingredients {
		for j := range p.ranges {
			if p.ingredients[i] >= p.ranges[j].min && p.ingredients[i] <= p.ranges[j].max {
				fresh++
				break
			}
		}
	}
	return fresh
}

package adventofcode2025

type Day06Puzzle struct {
	nums []uint // flat 2D: nums[row*cols+col], 1 alloc, contiguous
	ops  []byte
}

func NewDay06(lines []string) (Day06Puzzle, error) {
	h := len(lines)

	countTokens := func(line string) int {
		count := 0
		inToken := false
		for i := range len(line) {
			if line[i] == ' ' {
				inToken = false
			} else if !inToken {
				inToken = true
				count++
			}
		}
		return count
	}

	parseOps := func(line string, ops []byte) {
		col := 0
		for i := range len(line) {
			c := line[i]
			if c == '+' || c == '*' {
				ops[col] = c
				col++
			}
		}
	}

	parseNums := func(line string, nums []uint) {
		col := 0
		var n uint
		inNum := false
		for i := range len(line) {
			c := line[i]
			if c >= '0' && c <= '9' {
				n = n*10 + uint(c-'0')
				inNum = true
			} else if inNum {
				nums[col] = n
				col++
				n = 0
				inNum = false
			}
		}
		if inNum {
			nums[col] = n
		}
	}

	cols := countTokens(lines[h-1])
	rows := h - 1

	nums := make([]uint, rows*cols)
	ops := make([]byte, cols)

	parseOps(lines[h-1], ops)

	for i := range rows {
		parseNums(lines[i], nums[i*cols:(i+1)*cols])
	}

	return Day06Puzzle{nums: nums, ops: ops}, nil
}

func Day06(p Day06Puzzle, part1 bool) uint {
	cols := len(p.ops)
	rows := len(p.nums) / cols
	var total uint
	for col := range cols {
		var result uint
		if p.ops[col] == '*' {
			result = 1
		}
		for row := range rows {
			n := p.nums[row*cols+col]
			if p.ops[col] == '+' {
				result += n
			} else {
				result *= n
			}
		}
		total += result
	}
	return total
}

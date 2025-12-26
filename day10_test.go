package adventofcode2025

import "testing"

func TestDay10Part1Example(t *testing.T) {
	testWithParser(t, 10, exampleFilename, true, NewDay10, Day10, 7)
}

func TestDay10Part1(t *testing.T) {
	lines := linesFromFilename(t, filename(10))
	p, _ := NewDay10(lines)
	got := Day10(p, true)
	t.Logf("Day10 Part1: %d", got)
}

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParser(b, 10, true, NewDay10, Day10)
}

func TestDay10Part2Example(t *testing.T) {
	testWithParser(t, 10, exampleFilename, false, NewDay10, Day10, 33)
}

// Part 2 tries:
// | Attempt | Answer | Result   | Issue                                      |
// |---------|--------|----------|--------------------------------------------|
// | 1       | 16289  | too low  | Missing lower bounds in free var search   |
// | 2       | 16463  | ?        |                                            |

func TestDay10Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(10))
	p, _ := NewDay10(lines)
	got := Day10(p, false)
	t.Logf("Day10 Part2: %d", got)
}

func BenchmarkDay10Part2(b *testing.B) {
	benchWithParser(b, 10, false, NewDay10, Day10)
}

package adventofcode2025

import "testing"

func TestDay09Part1Example(t *testing.T) {
	testWithParser(t, 9, exampleFilename, true, NewDay09, Day09, 50)
}

func TestDay09Part1(t *testing.T) {
	testWithParser(t, 9, filename, true, NewDay09, Day09, 4755429952)
}

func BenchmarkDay09Part1(b *testing.B) {
	benchWithParser(b, 9, true, NewDay09, Day09)
}

func TestDay09Part2Example(t *testing.T) {
	testWithParser(t, 9, exampleFilename, false, NewDay09, Day09, 24)
}

func TestDay09Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(9))
	p, _ := NewDay09(lines)
	got := Day09(p, false)
	t.Logf("Day09 Part2: %d", got)
}

func BenchmarkDay09Part2(b *testing.B) {
	benchWithParser(b, 9, false, NewDay09, Day09)
}

package adventofcode2025

import "testing"

func TestDay11Part1Example(t *testing.T) {
	testWithParser(t, 11, exampleFilename, true, NewDay11, Day11, 5)
}

func TestDay11Part1(t *testing.T) {
	lines := linesFromFilename(t, filename(11))
	p, _ := NewDay11(lines)
	got := Day11(p, true)
	t.Logf("Day11 Part1: %d", got)
}

func BenchmarkDay11Part1(b *testing.B) {
	benchWithParser(b, 11, true, NewDay11, Day11)
}

func TestDay11Part2Example(t *testing.T) {
	lines := linesFromFilename(t, "testdata/day11_example2.txt")
	p, _ := NewDay11(lines)
	got := Day11(p, false)
	if got != 2 {
		t.Errorf("Expected 2, got %d", got)
	}
}

func TestDay11Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(11))
	p, _ := NewDay11(lines)
	got := Day11(p, false)
	t.Logf("Day11 Part2: %d", got)
}

func BenchmarkDay11Part2(b *testing.B) {
	benchWithParser(b, 11, false, NewDay11, Day11)
}

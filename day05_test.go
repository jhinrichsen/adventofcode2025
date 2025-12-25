package adventofcode2025

import "testing"

func TestDay05Part1Example(t *testing.T) {
	testWithParser(t, 5, exampleFilename, true, NewDay05, Day05, 3)
}

func TestDay05Part1(t *testing.T) {
	testWithParser(t, 5, filename, true, NewDay05, Day05, 821)
}

func BenchmarkDay05Part1(b *testing.B) {
	benchWithParser(b, 5, true, NewDay05, Day05)
}

func TestDay05Part2Example(t *testing.T) {
	testWithParser(t, 5, exampleFilename, false, NewDay05, Day05, 14)
}

func TestDay05Part2(t *testing.T) {
	testWithParser(t, 5, filename, false, NewDay05, Day05, 344771884978261)
}

func BenchmarkDay05Part2(b *testing.B) {
	benchWithParser(b, 5, false, NewDay05, Day05)
}

// brute force version benchmark: O(n*m) for comparison
func BenchmarkDay05Part1BruteForce(b *testing.B) {
	lines := linesFromFilename(b, filename(5))
	b.ResetTimer()
	for b.Loop() {
		p, _ := newDay05BruteForce(lines)
		day05BruteForce(p)
	}
}

func TestDay05BruteForceCorrectness(t *testing.T) {
	lines := linesFromFilename(t, filename(5))
	p, _ := newDay05BruteForce(lines)
	got := day05BruteForce(p)
	if got != 821 {
		t.Fatalf("want 821, got %d", got)
	}
}

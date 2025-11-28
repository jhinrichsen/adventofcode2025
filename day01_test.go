package adventofcode2025

import "testing"

func TestDay01Part1Example(t *testing.T) {
	testLines(t, 1, exampleFilename, true, Day01, 0)
}

func TestDay01Part1(t *testing.T) {
	testLines(t, 1, filename, true, Day01, 0)
}

func BenchmarkDay01Part1(b *testing.B) {
	benchLines(b, 1, true, Day01)
}

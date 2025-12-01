package adventofcode2025

import "testing"

func TestDay01Part1Example(t *testing.T) {
	testSolver(t, 1, exampleFilename, true, Day01, 3)
}

func TestDay01Part1(t *testing.T) {
	testSolver(t, 1, filename, true, Day01, 1105)
}

func BenchmarkDay01Part1(b *testing.B) {
	bench(b, 1, true, Day01)
}

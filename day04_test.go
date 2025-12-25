package adventofcode2025

import "testing"

func TestDay04Part1Example(t *testing.T) {
	testSolver(t, 4, exampleFilename, true, Day04, 13)
}

func TestDay04Part1(t *testing.T) {
	testSolver(t, 4, filename, true, Day04, 1480)
}

func BenchmarkDay04Part1(b *testing.B) {
	bench(b, 4, true, Day04)
}

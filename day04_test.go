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

func TestDay04Part2Example(t *testing.T) {
	testSolver(t, 4, exampleFilename, false, Day04, 43)
}

func TestDay04Part2(t *testing.T) {
	testSolver(t, 4, filename, false, Day04, 8899)
}

func BenchmarkDay04Part2(b *testing.B) {
	bench(b, 4, false, Day04)
}

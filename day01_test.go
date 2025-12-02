package adventofcode2025

import "testing"

func TestDay01Part1Example(t *testing.T) {
	testSolver(t, 1, exampleFilename, true, Day01, 3)
}

func TestDay01Part2Example(t *testing.T) {
	testSolver(t, 1, exampleFilename, false, Day01BruteForce, 6)
}

func TestDay01Part2ExamplePerformance(t *testing.T) {
	testSolver(t, 1, exampleFilename, false, Day01, 6)
}

func TestDay01Part1(t *testing.T) {
	testSolver(t, 1, filename, true, Day01, 1105)
}

func TestDay01Part1BruteForce(t *testing.T) {
	testSolver(t, 1, filename, true, Day01BruteForce, 1105)
}

func TestDay01Part2(t *testing.T) {
	testSolver(t, 1, filename, false, Day01, 6599)
}

func TestDay01Part2BruteForce(t *testing.T) {
	testSolver(t, 1, filename, false, Day01BruteForce, 6599)
}

func BenchmarkDay01Part1(b *testing.B) {
	bench(b, 1, true, Day01)
}

func BenchmarkDay01Part2(b *testing.B) {
	bench(b, 1, false, Day01)
}

func BenchmarkDay01Part1BruteForce(b *testing.B) {
	bench(b, 1, true, Day01BruteForce)
}

func BenchmarkDay01Part2BruteForce(b *testing.B) {
	bench(b, 1, false, Day01BruteForce)
}

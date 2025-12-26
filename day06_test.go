package adventofcode2025

import "testing"

func TestDay06Part1Example(t *testing.T) {
	testLines(t, 6, exampleFilename, true, Day06, 4277556)
}

func TestDay06Part1(t *testing.T) {
	testLines(t, 6, filename, true, Day06, 3525371263915)
}

func BenchmarkDay06Part1(b *testing.B) {
	benchLines(b, 6, true, Day06)
}

func TestDay06Part2Example(t *testing.T) {
	testLines(t, 6, exampleFilename, false, Day06, 3263827)
}

func TestDay06Part2(t *testing.T) {
	testLines(t, 6, filename, false, Day06, 6846480843636)
}

func BenchmarkDay06Part2(b *testing.B) {
	benchLines(b, 6, false, Day06)
}

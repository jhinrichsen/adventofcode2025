package adventofcode2025

import "testing"

func TestDay07Part1Example(t *testing.T) {
	testLines(t, 7, exampleFilename, true, Day07, 21)
}

func TestDay07Part1(t *testing.T) {
	testLines(t, 7, filename, true, Day07, 1638)
}

func BenchmarkDay07Part1(b *testing.B) {
	benchLines(b, 7, true, Day07)
}

func TestDay07Part2Example(t *testing.T) {
	testLines(t, 7, exampleFilename, false, Day07, 40)
}

func TestDay07Part2(t *testing.T) {
	testLines(t, 7, filename, false, Day07, 7759107121385)
}

func BenchmarkDay07Part2(b *testing.B) {
	benchLines(b, 7, false, Day07)
}

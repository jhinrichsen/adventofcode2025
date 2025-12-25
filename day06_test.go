package adventofcode2025

import "testing"

func TestDay06Part1Example(t *testing.T) {
	testWithParser(t, 6, exampleFilename, true, NewDay06, Day06, 4277556)
}

func TestDay06Part1(t *testing.T) {
	testWithParser(t, 6, filename, true, NewDay06, Day06, 3525371263915)
}

func BenchmarkDay06Part1(b *testing.B) {
	benchWithParser(b, 6, true, NewDay06, Day06)
}

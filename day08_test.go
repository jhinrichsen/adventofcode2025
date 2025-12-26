package adventofcode2025

import "testing"

func TestDay08Part1Example(t *testing.T) {
	testWithParser(t, 8, exampleFilename, true, NewDay08, Day08, 40)
}

func TestDay08Part1(t *testing.T) {
	testWithParser(t, 8, filename, true, NewDay08, Day08, 163548)
}

func BenchmarkDay08Part1(b *testing.B) {
	benchWithParser(b, 8, true, NewDay08, Day08)
}

func TestDay08Part2Example(t *testing.T) {
	testWithParser(t, 8, exampleFilename, false, NewDay08, Day08, 25272)
}

func TestDay08Part2(t *testing.T) {
	// testWithParser(t, 8, filename, false, NewDay08, Day08, 692420076) // wrong: Prim's finds different last edge
	testWithParser(t, 8, filename, false, NewDay08, Day08, 772452514)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchWithParser(b, 8, false, NewDay08, Day08)
}

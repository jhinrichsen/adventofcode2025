package adventofcode2025

import "testing"

func TestDay12Part1Example(t *testing.T) {
	testWithParser(t, 12, exampleFilename, true, NewDay12, Day12, 2)
}

func TestDay12ExampleRegions(t *testing.T) {
	lines := linesFromFilename(t, "testdata/day12_example.txt")
	p, _ := NewDay12(lines)

	tests := []struct {
		name string
		idx  int
		want bool
	}{
		{"4x4 with 2 shape-4", 0, true},
		{"12x5 with [1,0,1,0,2,2]", 1, true},
		{"12x5 with [1,0,1,0,3,2]", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := p.regions[tt.idx]
			got := canFit(p.shapes, r)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	lines := linesFromFilename(t, filename(12))
	p, _ := NewDay12(lines)
	got := Day12(p, true)
	t.Logf("Day12 Part1: %d", got)
}

func BenchmarkDay12Part1(b *testing.B) {
	benchWithParser(b, 12, true, NewDay12, Day12)
}

func TestDay12Part2Example(t *testing.T) {
	testWithParser(t, 12, exampleFilename, false, NewDay12, Day12, 0)
}

func TestDay12Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(12))
	p, _ := NewDay12(lines)
	got := Day12(p, false)
	t.Logf("Day12 Part2: %d", got)
}

func BenchmarkDay12Part2(b *testing.B) {
	benchWithParser(b, 12, false, NewDay12, Day12)
}

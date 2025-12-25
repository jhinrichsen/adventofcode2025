package adventofcode2025

import "testing"

func TestNumberOfDigits(t *testing.T) {
	tests := []struct {
		n    uint
		want uint
		even bool
	}{
		{1, 1, false},
		{13, 2, true},
		{123, 3, false},
		{1234, 4, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := numberOfDigits(tt.n)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
			b := evenNumberOfDigits(tt.n)
			if tt.even != b {
				t.Fatalf("want %t but got %t", tt.even, b)
			}

		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		n    uint
		want uint
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := pow(tt.n)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestTwin(t *testing.T) {
	tests := []struct {
		n    uint
		want uint
	}{
		{11, 11},
		{13, 11},
		{99, 99},
		{1234, 1212},
		{5678, 5656},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := twin(tt.n)
			if tt.want != got {
				t.Fatalf("twin(%d): want %d but got %d", tt.n, tt.want, got)
			}
		})
	}
}

func TestHasRepeatedPattern(t *testing.T) {
	tests := []struct {
		n    uint
		want bool
	}{
		// Part 2 examples from problem
		{11, true},      // "1" x 2
		{22, true},      // "2" x 2
		{99, true},      // "9" x 2
		{111, true},     // "1" x 3
		{999, true},     // "9" x 3
		{1010, true},    // "10" x 2
		{222222, true},  // "2" x 6 or "22" x 3 or "222" x 2
		{565656, true},  // "56" x 3
		{824824824, true}, // "824" x 3
		{2121212121, true}, // "21" x 5
		{12341234, true},   // "1234" x 2
		{123123123, true},  // "123" x 3
		{1212121212, true}, // "12" x 5
		{1111111, true},    // "1" x 7

		// Non-patterns
		{1188511885, true}, // "11885" x 2
		{123, false},
		{1234, false},
		{12323, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := hasRepeatedPattern(tt.n)
			if tt.want != got {
				t.Fatalf("hasRepeatedPattern(%d): want %t but got %t", tt.n, tt.want, got)
			}
		})
	}
}

func TestDay02Part1Example(t *testing.T) {
	testSolver(t, 2, exampleFilename, true, Day02, 1227775554)
}

func TestDay02Part2Example(t *testing.T) {
	testSolver(t, 2, exampleFilename, false, Day02, 4174379265)
}

func TestDay02Part1(t *testing.T) {
	// 21831176757 is too low
	testSolver(t, 2, filename, true, Day02, 21898734247)
}

func TestDay02Part2(t *testing.T) {
	testSolver(t, 2, filename, false, Day02, 28915664389)
}

func BenchmarkDay02Part1(b *testing.B) {
	bench(b, 2, true, Day02)
}

func BenchmarkDay02Part2(b *testing.B) {
	bench(b, 2, false, Day02)
}

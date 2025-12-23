package adventofcode2025

import "testing"

func TestDay03Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{"987654321111111\n", 98},
		{"811111111111119\n", 89},
		{"234234234234278\n", 78},
		{"818181911112111\n", 92},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Day03([]byte(tt.input), true)
			if err != nil {
				t.Fatalf("Day03(%s) error = %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Day03(%s) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDay03Part1(t *testing.T) {
	testSolver(t, 3, filename, true, Day03, 17343)
}

func BenchmarkDay03Part1(b *testing.B) {
	bench(b, 3, true, Day03)
}

func TestDay03Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{"987654321111111\n", 987654321111},
		{"811111111111119\n", 811111111119},
		{"234234234234278\n", 434234234278},
		{"818181911112111\n", 888911112111},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Day03([]byte(tt.input), false)
			if err != nil {
				t.Fatalf("Day03(%s) error = %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Day03(%s) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDay03Part2(t *testing.T) {
	testSolver(t, 3, filename, false, Day03, 172664333119298)
}

func BenchmarkDay03Part2(b *testing.B) {
	bench(b, 3, false, Day03)
}

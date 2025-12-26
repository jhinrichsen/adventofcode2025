package adventofcode2025

import (
	"testing"
)

// ILP = Integer Linear Programming
//
// Problem: Find non-negative integers x₁, x₂, ... xₙ that:
//   - Satisfy: A·x = b  (system of linear equations)
//   - Minimize: Σxᵢ     (sum of all variables)
//
// This appears in MANY real-world problems!

// =============================================================================
// EXAMPLE 1: Coin Change (Euro)
// =============================================================================
// You have Euro coins of 1, 2, 5, 10, 20, and 50 cents.
// What's the minimum number of coins to make exactly 83 cents?
//
// Variables: x₀ = #1ct, x₁ = #2ct, x₂ = #5ct, x₃ = #10ct, x₄ = #20ct, x₅ = #50ct
// Constraint: 1·x₀ + 2·x₁ + 5·x₂ + 10·x₃ + 20·x₄ + 50·x₅ = 83
// Minimize: x₀ + x₁ + x₂ + x₃ + x₄ + x₅

func TestILP_CoinChange(t *testing.T) {
	// Matrix A: one row per constraint, one column per coin type
	// The values are the coin denominations in Euro cents
	A := [][]int{
		{1, 2, 5, 10, 20, 50}, // Euro cent coins
	}
	b := []int{83} // target amount in cents

	x, numCoins, ok := ILPSolve(A, b)
	if !ok {
		t.Fatal("Expected solution")
	}

	// Optimal: 50 + 20 + 10 + 2 + 1 = 83 cents with 5 coins
	t.Logf("Coins: %d×1ct, %d×2ct, %d×5ct, %d×10ct, %d×20ct, %d×50ct = %d coins",
		x[0], x[1], x[2], x[3], x[4], x[5], numCoins)

	if numCoins != 5 {
		t.Errorf("Expected 5 coins, got %d", numCoins)
	}

	// Verify: check the sum equals 83
	total := x[0]*1 + x[1]*2 + x[2]*5 + x[3]*10 + x[4]*20 + x[5]*50
	if total != 83 {
		t.Errorf("Coins don't sum to 83: got %d", total)
	}
}

// =============================================================================
// EXAMPLE 2: Recipe Mixing
// =============================================================================
// A bakery makes cookies that need exactly:
//   - 10 cups flour
//   - 6 cups sugar
//
// They have two pre-mixed bags:
//   - Bag A: 2 cups flour + 1 cup sugar
//   - Bag B: 1 cup flour + 2 cups sugar
//
// What's the minimum number of bags to buy?
//
// Variables: x₀ = #BagA, x₁ = #BagB
// Constraints:
//   2·x₀ + 1·x₁ = 10  (flour)
//   1·x₀ + 2·x₁ = 6   (sugar)
// Minimize: x₀ + x₁

func TestILP_RecipeMixing_NoSolution(t *testing.T) {
	// This example has NO integer solution!
	// Solving: 2a + b = 10, a + 2b = 6
	// → a = 14/3, b = 2/3 (not integers)
	A := [][]int{
		{2, 1}, // flour: 2·BagA + 1·BagB = 10
		{1, 2}, // sugar: 1·BagA + 2·BagB = 6
	}
	b := []int{10, 6}

	_, _, ok := ILPSolve(A, b)
	if ok {
		t.Error("Expected no solution - requires fractional bags")
	}
}

// =============================================================================
// EXAMPLE 2b: Recipe Mixing (solvable version)
// =============================================================================
// Same bakery, different recipe that HAS an integer solution:
//   - 8 cups flour
//   - 7 cups sugar
//
// Bag A: 2 cups flour + 1 cup sugar
// Bag B: 1 cup flour + 2 cups sugar
//
// Solution: 3×BagA + 2×BagB = (6+2) flour + (3+4) sugar = 8 flour, 7 sugar ✓

func TestILP_RecipeMixing_WithSolution(t *testing.T) {
	A := [][]int{
		{2, 1}, // flour: 2·BagA + 1·BagB = 8
		{1, 2}, // sugar: 1·BagA + 2·BagB = 7
	}
	b := []int{8, 7}

	x, numBags, ok := ILPSolve(A, b)
	if !ok {
		t.Fatal("Expected solution")
	}

	t.Logf("Buy: %d×BagA, %d×BagB = %d bags total", x[0], x[1], numBags)

	// Verify
	flour := 2*x[0] + 1*x[1]
	sugar := 1*x[0] + 2*x[1]
	if flour != 8 || sugar != 7 {
		t.Errorf("Got %d flour, %d sugar; want 8, 7", flour, sugar)
	}
}

// =============================================================================
// EXAMPLE 3: Shift Scheduling
// =============================================================================
// A store needs workers during 4 time slots:
//   - Slot 0 (morning):   3 workers
//   - Slot 1 (midday):    5 workers
//   - Slot 2 (afternoon): 4 workers
//   - Slot 3 (evening):   2 workers
//
// Available shift patterns:
//   - Shift A: works slots 0,1
//   - Shift B: works slots 1,2
//   - Shift C: works slots 2,3
//   - Shift D: works slots 0,1,2
//
// Minimum workers to hire?

func TestILP_ShiftScheduling(t *testing.T) {
	// Rows = time slots, Cols = shift patterns
	// A[slot][shift] = 1 if shift covers that slot
	A := [][]int{
		{1, 0, 0, 1}, // slot 0: shifts A, D
		{1, 1, 0, 1}, // slot 1: shifts A, B, D
		{0, 1, 1, 1}, // slot 2: shifts B, C, D
		{0, 0, 1, 0}, // slot 3: shift C only
	}
	b := []int{3, 5, 4, 2} // workers needed per slot

	x, numWorkers, ok := ILPSolve(A, b)
	if !ok {
		t.Fatal("Expected solution")
	}

	t.Logf("Hire: %d×ShiftA, %d×ShiftB, %d×ShiftC, %d×ShiftD = %d workers",
		x[0], x[1], x[2], x[3], numWorkers)

	// Verify each slot has enough coverage
	for slot := 0; slot < 4; slot++ {
		coverage := 0
		for shift := 0; shift < 4; shift++ {
			coverage += A[slot][shift] * x[shift]
		}
		if coverage != b[slot] {
			t.Errorf("Slot %d: need %d, got %d", slot, b[slot], coverage)
		}
	}
}

// =============================================================================
// EXAMPLE 4: Resource Allocation (Button Pressing - like AoC Day 10!)
// =============================================================================
// You have buttons that each activate certain machines.
// Goal: activate each machine exactly the required number of times.
//
// Button 0: activates machines 0, 2
// Button 1: activates machines 0, 1, 3
// Target: machine 0 needs 28 activations
//         machine 1 needs 20 activations
//         machine 2 needs 8 activations
//         machine 3 needs 20 activations

func TestILP_ButtonPressing(t *testing.T) {
	// Rows = machines, Cols = buttons
	A := [][]int{
		{1, 1}, // machine 0: buttons 0,1
		{0, 1}, // machine 1: button 1 only
		{1, 0}, // machine 2: button 0 only
		{0, 1}, // machine 3: button 1 only
	}
	b := []int{28, 20, 8, 20}

	x, numPresses, ok := ILPSolve(A, b)
	if !ok {
		t.Fatal("Expected solution")
	}

	t.Logf("Press button 0 %d times, button 1 %d times = %d total",
		x[0], x[1], numPresses)

	// From constraints:
	// machine 1: x₁ = 20
	// machine 2: x₀ = 8
	// machine 3: x₁ = 20 ✓
	// machine 0: x₀ + x₁ = 8 + 20 = 28 ✓
	if x[0] != 8 || x[1] != 20 {
		t.Errorf("Expected [8, 20], got %v", x)
	}
	if numPresses != 28 {
		t.Errorf("Expected 28 presses, got %d", numPresses)
	}
}

// =============================================================================
// EXAMPLE 5: No Solution Case
// =============================================================================
// Sometimes the constraints are impossible to satisfy with integers.

func TestILP_NoSolution(t *testing.T) {
	// Two buttons, both activate the same machine
	// Machine 0: buttons 0,1 both activate it
	// But we want machine 0 at 3 AND machine 0 at 5 - impossible!
	A := [][]int{
		{1, 0}, // constraint 1: button 0 = 3
		{1, 0}, // constraint 2: button 0 = 5  (contradiction!)
	}
	b := []int{3, 5}

	_, _, ok := ILPSolve(A, b)
	if ok {
		t.Error("Expected no solution for contradictory constraints")
	}
}

// =============================================================================
// EXAMPLE 6: Day 10 Part 2 Example
// =============================================================================
// From the actual AoC puzzle example

func TestILP_Day10Example(t *testing.T) {
	// Machine 1: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
	// 4 counters, 6 buttons
	// Button 0=(3), 1=(1,3), 2=(2), 3=(2,3), 4=(0,2), 5=(0,1)
	A := [][]int{
		{0, 0, 0, 0, 1, 1}, // counter 0: buttons 4,5
		{0, 1, 0, 0, 0, 1}, // counter 1: buttons 1,5
		{0, 0, 1, 1, 1, 0}, // counter 2: buttons 2,3,4
		{1, 1, 0, 1, 0, 0}, // counter 3: buttons 0,1,3
	}
	b := []int{3, 5, 4, 7}

	x, sum, ok := ILPSolve(A, b)
	if !ok {
		t.Fatal("Expected solution")
	}

	t.Logf("Button presses: %v = %d total", x, sum)

	if sum != 10 {
		t.Errorf("Expected 10 presses, got %d", sum)
	}
}

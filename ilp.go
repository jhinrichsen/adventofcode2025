package adventofcode2025

// ILPSolve finds non-negative integer solution x minimizing sum(x) where Ax = b.
// A is an m×n matrix (m equations, n variables).
// Returns solution x, minimum sum, and whether a valid solution exists.
func ILPSolve(A [][]int, b []int) (x []int, minSum int, ok bool) {
	if len(A) == 0 || len(b) == 0 {
		return nil, 0, false
	}

	m := len(A)    // rows (equations)
	n := len(A[0]) // cols (variables)

	// Rational number for exact arithmetic
	type frac struct{ n, d int }

	gcd := func(a, b int) int {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	norm := func(f frac) frac {
		if f.n == 0 {
			return frac{0, 1}
		}
		if f.d < 0 {
			f.n, f.d = -f.n, -f.d
		}
		g := gcd(f.n, f.d)
		return frac{f.n / g, f.d / g}
	}

	sub := func(a, b frac) frac {
		return norm(frac{a.n*b.d - b.n*a.d, a.d * b.d})
	}

	mul := func(a, b frac) frac {
		return norm(frac{a.n * b.n, a.d * b.d})
	}

	div := func(a, b frac) frac {
		return norm(frac{a.n * b.d, a.d * b.n})
	}

	// Build augmented matrix [A|b]
	mat := make([][]frac, m)
	for i := range mat {
		mat[i] = make([]frac, n+1)
		for j := 0; j < n; j++ {
			mat[i][j] = frac{A[i][j], 1}
		}
		mat[i][n] = frac{b[i], 1}
	}

	// Gaussian elimination with partial pivoting
	pivotRow := 0
	pivotCols := []int{}

	for col := 0; col < n && pivotRow < m; col++ {
		// Find pivot
		bestRow := -1
		for row := pivotRow; row < m; row++ {
			if mat[row][col].n != 0 {
				bestRow = row
				break
			}
		}
		if bestRow < 0 {
			continue
		}

		mat[pivotRow], mat[bestRow] = mat[bestRow], mat[pivotRow]
		pivotCols = append(pivotCols, col)

		// Scale pivot row
		pivot := mat[pivotRow][col]
		for j := range mat[pivotRow] {
			mat[pivotRow][j] = div(mat[pivotRow][j], pivot)
		}

		// Eliminate column
		for row := 0; row < m; row++ {
			if row == pivotRow || mat[row][col].n == 0 {
				continue
			}
			factor := mat[row][col]
			for j := range mat[row] {
				mat[row][j] = sub(mat[row][j], mul(factor, mat[pivotRow][j]))
			}
		}
		pivotRow++
	}

	// Check for inconsistency
	for row := pivotRow; row < m; row++ {
		if mat[row][n].n != 0 {
			return nil, 0, false
		}
	}

	// Identify free variables
	pivotSet := make(map[int]bool)
	for _, c := range pivotCols {
		pivotSet[c] = true
	}
	freeVars := []int{}
	for j := 0; j < n; j++ {
		if !pivotSet[j] {
			freeVars = append(freeVars, j)
		}
	}

	// No free variables: unique solution
	if len(freeVars) == 0 {
		solution := make([]int, n)
		total := 0
		for i, col := range pivotCols {
			v := mat[i][n]
			if v.d != 1 || v.n < 0 {
				return nil, 0, false
			}
			solution[col] = v.n
			total += v.n
		}
		return solution, total, true
	}

	// With free variables: search for minimum sum
	// Compute upper bound for each free variable from the original constraint
	maxFree := make([]int, len(freeVars))
	for fi, fv := range freeVars {
		maxFree[fi] = 0
		for i := 0; i < m; i++ {
			if A[i][fv] > 0 {
				bound := b[i] / A[i][fv]
				if maxFree[fi] == 0 || bound < maxFree[fi] {
					maxFree[fi] = bound
				}
			}
		}
		if maxFree[fi] == 0 {
			// Fallback: use max of b
			for _, v := range b {
				if v > maxFree[fi] {
					maxFree[fi] = v
				}
			}
		}
	}

	best := -1
	var bestSolution []int
	freeVals := make([]int, len(freeVars))

	var search func(idx, partialSum int)
	search = func(idx, partialSum int) {
		// Prune if partial sum already >= best
		if best >= 0 && partialSum >= best {
			return
		}

		if idx == len(freeVars) {
			// Compute full solution
			solution := make([]int, n)
			total := partialSum

			// Set free variables
			for fi, fv := range freeVars {
				solution[fv] = freeVals[fi]
			}

			// Compute pivot variables
			for i, col := range pivotCols {
				val := mat[i][n]
				for fi, fv := range freeVars {
					val = sub(val, mul(mat[i][fv], frac{freeVals[fi], 1}))
				}
				if val.d != 1 || val.n < 0 {
					return // Invalid
				}
				solution[col] = val.n
				total += val.n
			}

			if best < 0 || total < best {
				best = total
				bestSolution = make([]int, n)
				copy(bestSolution, solution)
			}
			return
		}

		for v := 0; v <= maxFree[idx]; v++ {
			freeVals[idx] = v
			search(idx+1, partialSum+v)
		}
	}

	search(0, 0)

	if best < 0 {
		return nil, 0, false
	}
	return bestSolution, best, true
}

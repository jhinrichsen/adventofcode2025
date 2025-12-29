package adventofcode2025

import (
	"strings"
)

type Day12Shape struct {
	cells [][2]int // relative positions of cells
}

type Day12Region struct {
	width, height int
	quantities    []int // how many of each shape
}

type Day12Puzzle struct {
	shapes  [][]Day12Shape // shapes[i] = all orientations of shape i
	regions []Day12Region
}

func NewDay12(lines []string) (Day12Puzzle, error) {
	var p Day12Puzzle
	p.shapes = make([][]Day12Shape, 0)

	i := 0
	// Parse shapes
	for i < len(lines) {
		line := lines[i]
		if strings.Contains(line, "x") {
			break // start of regions
		}
		if strings.HasSuffix(line, ":") {
			// Shape header
			var grid []string
			i++
			for i < len(lines) && len(lines[i]) > 0 && !strings.HasSuffix(lines[i], ":") && !strings.Contains(lines[i], "x") {
				grid = append(grid, lines[i])
				i++
			}
			orientations := generateOrientations(grid)
			p.shapes = append(p.shapes, orientations)
		} else {
			i++
		}
	}

	// Parse regions
	for i < len(lines) {
		line := lines[i]
		if strings.Contains(line, "x") {
			var r Day12Region
			parts := strings.Split(line, ": ")
			dims := strings.Split(parts[0], "x")
			r.width = atoi(dims[0])
			r.height = atoi(dims[1])
			nums := strings.Fields(parts[1])
			for _, n := range nums {
				r.quantities = append(r.quantities, atoi(n))
			}
			p.regions = append(p.regions, r)
		}
		i++
	}

	return p, nil
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}

func generateOrientations(grid []string) []Day12Shape {
	// Get base shape cells
	var cells [][2]int
	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				cells = append(cells, [2]int{r, c})
			}
		}
	}

	seen := make(map[string]bool)
	var result []Day12Shape

	// Generate all 8 orientations (4 rotations × 2 flips)
	current := cells
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			normalized := normalize(current)
			key := shapeKey(normalized)
			if !seen[key] {
				seen[key] = true
				result = append(result, Day12Shape{cells: normalized})
			}
			current = rotate90(current)
		}
		current = flipH(cells)
	}

	return result
}

func normalize(cells [][2]int) [][2]int {
	if len(cells) == 0 {
		return cells
	}
	minR, minC := cells[0][0], cells[0][1]
	for _, c := range cells {
		if c[0] < minR {
			minR = c[0]
		}
		if c[1] < minC {
			minC = c[1]
		}
	}
	result := make([][2]int, len(cells))
	for i, c := range cells {
		result[i] = [2]int{c[0] - minR, c[1] - minC}
	}
	return result
}

func rotate90(cells [][2]int) [][2]int {
	result := make([][2]int, len(cells))
	for i, c := range cells {
		result[i] = [2]int{c[1], -c[0]}
	}
	return result
}

func flipH(cells [][2]int) [][2]int {
	result := make([][2]int, len(cells))
	for i, c := range cells {
		result[i] = [2]int{c[0], -c[1]}
	}
	return result
}

func shapeKey(cells [][2]int) string {
	// Sort and stringify
	sorted := make([][2]int, len(cells))
	copy(sorted, cells)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i][0] > sorted[j][0] || (sorted[i][0] == sorted[j][0] && sorted[i][1] > sorted[j][1]) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	var sb strings.Builder
	for _, c := range sorted {
		sb.WriteByte(byte(c[0] + 'A'))
		sb.WriteByte(byte(c[1] + 'A'))
	}
	return sb.String()
}

func Day12(p Day12Puzzle, part1 bool) uint {
	if part1 {
		return day12Part1(p)
	}
	return day12Part2(p)
}

func day12Part1(p Day12Puzzle) uint {
	var count uint
	for _, region := range p.regions {
		if canFit(p.shapes, region) {
			count++
		}
	}
	return count
}

func canFit(shapes [][]Day12Shape, region Day12Region) bool {
	// Copy quantities as remaining counts
	remaining := make([]int, len(region.quantities))
	copy(remaining, region.quantities)

	// Check total area
	totalShapeArea := 0
	hasAny := false
	for i, q := range remaining {
		if q > 0 {
			hasAny = true
			if len(shapes[i]) > 0 {
				totalShapeArea += q * len(shapes[i][0].cells)
			}
		}
	}
	if !hasAny {
		return true
	}

	gridArea := region.width * region.height
	if totalShapeArea > gridArea {
		return false
	}

	// Grid
	grid := make([][]bool, region.height)
	for i := range grid {
		grid[i] = make([]bool, region.width)
	}

	return backtrack(shapes, grid, remaining, region.width, region.height, 0, 0)
}

func backtrack(shapes [][]Day12Shape, grid [][]bool, remaining []int, w, h, shapeIdx, minPos int) bool {
	// Find next shape type to place
	for shapeIdx < len(remaining) && remaining[shapeIdx] == 0 {
		shapeIdx++
		minPos = 0 // Reset position constraint for new shape type
	}

	// All shapes placed
	if shapeIdx >= len(remaining) {
		return true
	}

	// Try placing one instance of this shape type
	// Start from minPos to avoid symmetric duplicates within same shape type
	for _, orient := range shapes[shapeIdx] {
		for pos := minPos; pos < w*h; pos++ {
			pr, pc := pos/w, pos%w
			if canPlace(grid, orient, pr, pc, w, h) {
				place(grid, orient, pr, pc, true)
				remaining[shapeIdx]--
				// Next shape of same type must be placed after this position
				if backtrack(shapes, grid, remaining, w, h, shapeIdx, pos+1) {
					return true
				}
				remaining[shapeIdx]++
				place(grid, orient, pr, pc, false)
			}
		}
	}

	return false
}

func canPlace(grid [][]bool, shape Day12Shape, pr, pc, w, h int) bool {
	for _, cell := range shape.cells {
		r, c := pr+cell[0], pc+cell[1]
		if r < 0 || r >= h || c < 0 || c >= w || grid[r][c] {
			return false
		}
	}
	return true
}

func place(grid [][]bool, shape Day12Shape, pr, pc int, fill bool) {
	for _, cell := range shape.cells {
		grid[pr+cell[0]][pc+cell[1]] = fill
	}
}


func day12Part2(p Day12Puzzle) uint {
	return 0
}

package adventofcode2025

import (
	"image"
	"slices"
)

type Day09Puzzle struct {
	points []image.Point
}

func NewDay09(lines []string) (Day09Puzzle, error) {
	points := make([]image.Point, len(lines))
	for i, line := range lines {
		var x, y int
		j := 0
		for ; j < len(line) && line[j] != ','; j++ {
			x = x*10 + int(line[j]-'0')
		}
		j++
		for ; j < len(line); j++ {
			y = y*10 + int(line[j]-'0')
		}
		points[i] = image.Point{x, y}
	}
	return Day09Puzzle{points: points}, nil
}

func Day09(p Day09Puzzle, part1 bool) uint {
	if part1 {
		return day09Part1(p)
	}
	return day09Part2(p)
}

func day09Part1(p Day09Puzzle) uint {
	var maxArea uint
	n := len(p.points)
	for i := range n {
		for j := i + 1; j < n; j++ {
			r := image.Rectangle{p.points[i], p.points[j]}.Canon()
			area := uint((r.Dx() + 1) * (r.Dy() + 1))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func day09Part2(p Day09Puzzle) uint {
	n := len(p.points)
	if n == 0 {
		return 0
	}

	// collect unique X and Y coordinates for compression
	xSet := make(map[int]struct{})
	ySet := make(map[int]struct{})
	for _, pt := range p.points {
		xSet[pt.X] = struct{}{}
		ySet[pt.Y] = struct{}{}
	}

	// sort and create mappings
	xs := make([]int, 0, len(xSet))
	for x := range xSet {
		xs = append(xs, x)
	}
	ys := make([]int, 0, len(ySet))
	for y := range ySet {
		ys = append(ys, y)
	}
	slices.Sort(xs)
	slices.Sort(ys)

	xIdx := make(map[int]int, len(xs))
	for i, x := range xs {
		xIdx[x] = i
	}
	yIdx := make(map[int]int, len(ys))
	for i, y := range ys {
		yIdx[y] = i
	}

	w := len(xs) + 2 // +2 for border
	h := len(ys) + 2

	// 0 = unknown, 1 = boundary, 2 = exterior
	grid := make([]byte, w*h)
	idx := func(xi, yi int) int { return yi*w + xi }

	// mark red tiles as boundary (offset by 1 for border)
	for _, pt := range p.points {
		grid[idx(xIdx[pt.X]+1, yIdx[pt.Y]+1)] = 1
	}

	// mark lines between consecutive red tiles
	for i := range n {
		p1 := p.points[i]
		p2 := p.points[(i+1)%n]
		xi1, yi1 := xIdx[p1.X]+1, yIdx[p1.Y]+1
		xi2, yi2 := xIdx[p2.X]+1, yIdx[p2.Y]+1

		if xi1 == xi2 {
			// vertical line - mark all y indices between
			if yi1 > yi2 {
				yi1, yi2 = yi2, yi1
			}
			for yi := yi1; yi <= yi2; yi++ {
				grid[idx(xi1, yi)] = 1
			}
		} else {
			// horizontal line - mark all x indices between
			if xi1 > xi2 {
				xi1, xi2 = xi2, xi1
			}
			for xi := xi1; xi <= xi2; xi++ {
				grid[idx(xi, yi1)] = 1
			}
		}
	}

	// flood fill exterior from (0,0)
	stack := []image.Point{{0, 0}}
	grid[0] = 2
	for len(stack) > 0 {
		pt := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, d := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			np := pt.Add(d)
			if np.X >= 0 && np.X < w && np.Y >= 0 && np.Y < h {
				i := np.Y*w + np.X
				if grid[i] == 0 {
					grid[i] = 2
					stack = append(stack, np)
				}
			}
		}
	}

	// valid cells are boundary (1) or interior (still 0)
	isValidCell := func(xi, yi int) bool {
		v := grid[idx(xi, yi)]
		return v == 0 || v == 1
	}

	// check each pair of red tiles
	var maxArea uint
	for i := range n {
		for j := i + 1; j < n; j++ {
			xi1, yi1 := xIdx[p.points[i].X]+1, yIdx[p.points[i].Y]+1
			xi2, yi2 := xIdx[p.points[j].X]+1, yIdx[p.points[j].Y]+1
			if xi1 > xi2 {
				xi1, xi2 = xi2, xi1
			}
			if yi1 > yi2 {
				yi1, yi2 = yi2, yi1
			}

			// check all compressed cells in rectangle
			valid := true
			for yi := yi1; yi <= yi2 && valid; yi++ {
				for xi := xi1; xi <= xi2 && valid; xi++ {
					if !isValidCell(xi, yi) {
						valid = false
					}
				}
			}

			if valid {
				r := image.Rectangle{p.points[i], p.points[j]}.Canon()
				area := uint((r.Dx() + 1) * (r.Dy() + 1))
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}


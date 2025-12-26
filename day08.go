package adventofcode2025

import "slices"

type Day08Puzzle struct {
	points [][3]int
}

func NewDay08(lines []string) (Day08Puzzle, error) {
	points := make([][3]int, len(lines))
	for i, line := range lines {
		var x, y, z int
		for k := range line {
			if line[k] == ',' {
				x = y
				y = z
				z = 0
			} else {
				z = z*10 + int(line[k]-'0')
			}
		}
		points[i] = [3]int{x, y, z}
	}
	return Day08Puzzle{points: points}, nil
}

func Day08(p Day08Puzzle, part1 bool) uint {
	n := len(p.points)
	k := 1000
	if n <= 100 {
		k = n / 2
	}

	// compute and sort edges
	type edge struct {
		i, j   int
		distSq int64
	}
	edges := make([]edge, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := int64(p.points[i][0] - p.points[j][0])
			dy := int64(p.points[i][1] - p.points[j][1])
			dz := int64(p.points[i][2] - p.points[j][2])
			edges = append(edges, edge{i, j, dx*dx + dy*dy + dz*dz})
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		if a.distSq < b.distSq {
			return -1
		}
		if a.distSq > b.distSq {
			return 1
		}
		return 0
	})

	// union-find
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var lastI, lastJ int
	unions := 0
	pairs := 0

	for idx := range edges {
		e := edges[idx]
		rx, ry := find(e.i), find(e.j)
		if rx != ry {
			if size[rx] < size[ry] {
				rx, ry = ry, rx
			}
			parent[ry] = rx
			size[rx] += size[ry]
			unions++
			lastI, lastJ = e.i, e.j
		}
		pairs++

		if part1 && pairs >= k {
			break
		}
		if !part1 && unions >= n-1 {
			break
		}
	}

	if part1 {
		sizes := make([]int, 0, n)
		for i := range n {
			if parent[i] == i {
				sizes = append(sizes, size[i])
			}
		}
		slices.Sort(sizes)
		m := len(sizes)
		if m < 3 {
			return 0
		}
		return uint(sizes[m-1]) * uint(sizes[m-2]) * uint(sizes[m-3])
	}

	return uint(p.points[lastI][0]) * uint(p.points[lastJ][0])
}

package adventofcode2025

import (
	"strings"
)

type Day11Puzzle struct {
	graph map[string][]string
}

func NewDay11(lines []string) (Day11Puzzle, error) {
	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		src := parts[0]
		dests := strings.Fields(parts[1])
		graph[src] = append(graph[src], dests...)
	}
	return Day11Puzzle{graph: graph}, nil
}

func Day11(p Day11Puzzle, part1 bool) uint {
	if part1 {
		return day11Part1(p)
	}
	return day11Part2(p)
}

func day11Part1(p Day11Puzzle) uint {
	return countPaths(p.graph, "you", "out", make(map[string]bool))
}

func countPaths(graph map[string][]string, current, target string, visited map[string]bool) uint {
	if current == target {
		return 1
	}
	if visited[current] {
		return 0
	}

	visited[current] = true
	var count uint
	for _, next := range graph[current] {
		count += countPaths(graph, next, target, visited)
	}
	visited[current] = false

	return count
}

func day11Part2(p Day11Puzzle) uint {
	// DAG: use memoization with state (node, hasDac, hasFft)
	type state struct {
		node   string
		hasDac bool
		hasFft bool
	}
	memo := make(map[state]uint)

	var count func(node string, hasDac, hasFft bool) uint
	count = func(node string, hasDac, hasFft bool) uint {
		if node == "dac" {
			hasDac = true
		}
		if node == "fft" {
			hasFft = true
		}

		if node == "out" {
			if hasDac && hasFft {
				return 1
			}
			return 0
		}

		s := state{node, hasDac, hasFft}
		if v, ok := memo[s]; ok {
			return v
		}

		var total uint
		for _, next := range p.graph[node] {
			total += count(next, hasDac, hasFft)
		}
		memo[s] = total
		return total
	}

	return count("svr", false, false)
}

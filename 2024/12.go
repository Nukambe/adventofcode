package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
)

func calcRegion(plot vector, grid []string, visited map[vector]struct{}) (int, int) {
	plant := grid[plot.y][plot.x]
	visited[plot] = struct{}{}
	area := 1
	perimeter := 0
	a := 0
	p := 0
	if plot.y > 0 { // up
		up := vector{plot.x, plot.y - 1}
		if grid[plot.y-1][plot.x] == plant {
			if _, ok := visited[up]; !ok {
				a, p = calcRegion(up, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}
	if plot.y < len(grid)-1 { // down
		down := vector{plot.x, plot.y + 1}
		if grid[plot.y+1][plot.x] == plant {
			if _, ok := visited[down]; !ok {
				a, p = calcRegion(down, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}
	if plot.x > 0 { // left
		left := vector{plot.x - 1, plot.y}
		if grid[plot.y][plot.x-1] == plant {
			if _, ok := visited[left]; !ok {
				a, p = calcRegion(left, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}
	if plot.x < len(grid[plot.y])-1 { // right
		right := vector{plot.x + 1, plot.y}
		if grid[plot.y][plot.x+1] == plant {
			if _, ok := visited[right]; !ok {
				a, p = calcRegion(right, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}
	return area, perimeter
}

func SolveDay12Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "12.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	cost := 0

	visited := map[vector]struct{}{}
	for y, plots := range lines {
		for x, plot := range plots {
			v := vector{x, y}
			if _, ok := visited[v]; ok {
				continue
			}
			fmt.Println("Plot:", string(plot))
			area, perimeter := calcRegion(v, lines, visited)
			cost += area * perimeter
		}
	}

	return fmt.Sprintf("%d", cost)
}

func calcRegion2(plot vector, grid []string, visited map[vector]struct{}) (int, int) {
	plant := grid[plot.y][plot.x]
	visited[plot] = struct{}{}
	area := 1
	sides := 4
	a := 0
	s := 0
	if plot.y > 0 { // up
		up := vector{plot.x, plot.y - 1}
		if grid[plot.y-1][plot.x] == plant {
			if _, ok := visited[up]; !ok {
				a, s = calcRegion(up, grid, visited)
				area += a
				sides -= 2 + s
			}
		}
	}
	if plot.y < len(grid)-1 { // down
		down := vector{plot.x, plot.y + 1}
		if grid[plot.y+1][plot.x] == plant {
			if _, ok := visited[down]; !ok {
				a, s = calcRegion(down, grid, visited)
				area += a
				sides -= 2 + s
			}
		}
	}
	if plot.x > 0 { // left
		left := vector{plot.x - 1, plot.y}
		if grid[plot.y][plot.x-1] == plant {
			if _, ok := visited[left]; !ok {
				a, s = calcRegion(left, grid, visited)
				area += a
				sides -= 2 + s
			}
		}
	}
	if plot.x < len(grid[plot.y])-1 { // right
		right := vector{plot.x + 1, plot.y}
		if grid[plot.y][plot.x+1] == plant {
			if _, ok := visited[right]; !ok {
				a, s = calcRegion(right, grid, visited)
				area += a
				sides -= 2 + s
			}
		}
	}
	return area, sides
}

func SolveDay12Puzzle2() string {
	return SolveDay12Puzzle1()
}

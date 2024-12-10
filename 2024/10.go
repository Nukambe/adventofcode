package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
)

func getTrailScore(start vector, grid []string, peaks map[vector]struct{}) map[vector]struct{} {
	height, _ := strconv.Atoi(string(grid[start.y][start.x]))
	if height == 9 {
		peaks[start] = struct{}{}
		return peaks
	}
	target := height + 1

	if start.y > 0 { // up
		next := helpers.CharToInt(grid[start.y-1][start.x])
		if next == target {
			peaks = getTrailScore(vector{start.x, start.y - 1}, grid, peaks)
		}
	}
	if start.y < len(grid)-1 { // down
		next := helpers.CharToInt(grid[start.y+1][start.x])
		if next == target {
			peaks = getTrailScore(vector{start.x, start.y + 1}, grid, peaks)
		}
	}
	if start.x > 0 { // left
		next := helpers.CharToInt(grid[start.y][start.x-1])
		if next == target {
			peaks = getTrailScore(vector{start.x - 1, start.y}, grid, peaks)
		}
	}
	if start.x < len(grid[start.y])-1 { // right
		next := helpers.CharToInt(grid[start.y][start.x+1])
		if next == target {
			peaks = getTrailScore(vector{start.x + 1, start.y}, grid, peaks)
		}
	}
	return peaks
}

func SolveDay10Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "10.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for y, heights := range lines {
		for x, height := range heights {
			if height != '0' {
				continue
			}
			total += len(getTrailScore(vector{x, y}, lines, map[vector]struct{}{}))
		}
	}

	return fmt.Sprintf("%d", total)
}

func getTrailRating(start vector, grid []string) int {
	height, _ := strconv.Atoi(string(grid[start.y][start.x]))
	if height == 9 {
		return 1
	}
	target := height + 1
	up := 0
	down := 0
	left := 0
	right := 0

	if start.y > 0 { // up
		next := helpers.CharToInt(grid[start.y-1][start.x])
		if next == target {
			up = getTrailRating(vector{start.x, start.y - 1}, grid)
		}
	}
	if start.y < len(grid)-1 { // down
		next := helpers.CharToInt(grid[start.y+1][start.x])
		if next == target {
			down = getTrailRating(vector{start.x, start.y + 1}, grid)
		}
	}
	if start.x > 0 { // left
		next := helpers.CharToInt(grid[start.y][start.x-1])
		if next == target {
			left = getTrailRating(vector{start.x - 1, start.y}, grid)
		}
	}
	if start.x < len(grid[start.y])-1 { // right
		next := helpers.CharToInt(grid[start.y][start.x+1])
		if next == target {
			right = getTrailRating(vector{start.x + 1, start.y}, grid)
		}
	}
	return up + down + left + right
}

func SolveDay10Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "10.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for y, heights := range lines {
		for x, height := range heights {
			if height != '0' {
				continue
			}
			total += getTrailRating(vector{x, y}, lines)
		}
	}

	return fmt.Sprintf("%d", total)
}

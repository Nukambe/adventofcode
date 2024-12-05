package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
)

func checkMAS(m, a, s uint8) bool {
	return m == 'M' && a == 'A' && s == 'S'
}

func SolveDay4Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "4.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for y := 0; y < len(grid); y++ {
		row := grid[y]
		for x := 0; x < len(row); x++ {
			if row[x] != 'X' {
				continue
			}

			if x < len(row)-3 { // right
				if row[x:x+4] == "XMAS" {
					total++
				}
				if y > 2 { // right + up
					if checkMAS(grid[y-1][x+1], grid[y-2][x+2], grid[y-3][x+3]) {
						total++
					}
				}
				if y < len(grid)-3 { // right + down
					if checkMAS(grid[y+1][x+1], grid[y+2][x+2], grid[y+3][x+3]) {
						total++
					}
				}
			}
			if x > 2 { // left
				if row[x-3:x+1] == "SAMX" {
					total++
				}
				if y > 2 { // left + up
					if checkMAS(grid[y-1][x-1], grid[y-2][x-2], grid[y-3][x-3]) {
						total++
					}
				}
				if y < len(grid)-3 { // left + down
					if checkMAS(grid[y+1][x-1], grid[y+2][x-2], grid[y+3][x-3]) {
						total++
					}
				}
			}
			if y < len(grid)-3 { // down
				if checkMAS(grid[y+1][x], grid[y+2][x], grid[y+3][x]) {
					total++
				}
			}
			if y > 2 { // up
				if checkMAS(grid[y-1][x], grid[y-2][x], grid[y-3][x]) {
					total++
				}
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

func SolveDay4Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "4.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for y := 1; y < len(grid)-1; y++ {
		row := grid[y]
		for x := 1; x < len(row)-1; x++ {
			if row[x] != 'A' {
				continue
			}

			if (checkMAS(grid[y-1][x-1], grid[y][x], grid[y+1][x+1]) ||
				checkMAS(grid[y+1][x+1], grid[y][x], grid[y-1][x-1])) &&
				(checkMAS(grid[y+1][x-1], grid[y][x], grid[y-1][x+1]) ||
					checkMAS(grid[y-1][x+1], grid[y][x], grid[y+1][x-1])) {
				total++
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
)

func findLines(grid []string) map[int32][]vector {
	frequencies := map[int32][]vector{}
	for y, row := range grid {
		for x, f := range row {
			if f == '.' {
				continue
			}

			_, ok := frequencies[f]
			if !ok {
				frequencies[f] = []vector{}
			}
			frequencies[f] = append(frequencies[f], vector{x, y})
		}
	}
	return frequencies
}

func SolveDay8Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "8.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)

	antinodes := map[vector]struct{}{}

	frequencies := findLines(grid)
	for _, vectors := range frequencies {
		for i, start := range vectors {
			for j, end := range vectors {
				if i == j {
					continue
				}

				m := vector{
					x: end.x - start.x,
					y: end.y - start.y,
				}
				antinode := addVectors(end, m)
				if antinode.x < 0 || antinode.x >= len(grid[0]) ||
					antinode.y < 0 || antinode.y >= len(grid) {
					continue
				}
				antinodes[antinode] = struct{}{}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

func SolveDay8Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "8.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)

	antinodes := map[vector]struct{}{}

	frequencies := findLines(grid)
	for _, vectors := range frequencies {
		for i, start := range vectors {
			for j, end := range vectors {
				if i == j {
					continue
				}
				m := vector{
					x: end.x - start.x,
					y: end.y - start.y,
				}

				antinodes[start] = struct{}{}
				antinodes[end] = struct{}{}
				antinode := end
				for {
					antinode = addVectors(antinode, m)
					if antinode.x < 0 || antinode.x >= len(grid[0]) ||
						antinode.y < 0 || antinode.y >= len(grid) {
						break
					}
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

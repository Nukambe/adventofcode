package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
)

type vector struct {
	x int
	y int
}

type guard struct {
	location  vector
	direction vector
	path      map[vector]int      // set
	loop      map[vector]struct{} // set
}

func addVectors(v1 vector, v2 vector) vector {
	return vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func (g *guard) patrol(grid []string, extra vector) bool {
	next := addVectors(g.location, g.direction)
	if next.y >= len(grid) || next.y < 0 || next.x >= len(grid[g.location.y]) || next.x < 0 {
		return false
	}

	if grid[next.y][next.x] == '#' || next == extra {
		g.turn()
		return g.patrol(grid, extra)
	}
	g.location = next // move
	_, ok := g.path[g.location]
	if !ok {
		g.path[g.location] = 0
	}
	g.path[g.location]++ // save location
	if g.path[g.location] > 3 {
		g.loop[g.location] = struct{}{}
	}
	return true
}

func (g *guard) turn() {
	up := vector{0, -1}
	down := vector{0, 1}
	left := vector{-1, 0}
	right := vector{1, 0}

	switch g.direction {
	case up:
		g.direction = right
	case down:
		g.direction = left
	case left:
		g.direction = up
	case right:
		g.direction = down
	}
}

func SolveDay6Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "6.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)

	gallivant := guard{
		location:  vector{},
		direction: vector{0, -1},
		path:      map[vector]int{},
	}

	for y, line := range grid {
		for x, tile := range line {
			if tile != '^' {
				continue
			}
			gallivant.location = vector{x, y}
			gallivant.path[gallivant.location] = 1
		}
	}

	for gallivant.patrol(grid, vector{-1, -1}) {
		// patrolling...
	}

	return fmt.Sprintf("%d", len(gallivant.path))
}

func SolveDay6Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "6.txt")
	grid, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	start := vector{}
	for y, line := range grid { // find start
		for x, tile := range line {
			if tile != '^' {
				continue
			}
			start = vector{x, y}
		}
	}

	for y, row := range grid {
		for x, col := range row {
			if col == '#' || col == '^' {
				continue
			}

			gallivant := guard{
				location:  start,
				direction: vector{0, -1},
				path:      map[vector]int{start: 1},
				loop:      map[vector]struct{}{},
			}

			for gallivant.patrol(grid, vector{x, y}) {
				if len(gallivant.loop) > 1 {
					total++
					break
				}
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

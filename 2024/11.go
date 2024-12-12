package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
	"strings"
)

func splitStone(stone string) []int {
	left, _ := strconv.Atoi(stone[:len(stone)/2])
	right, _ := strconv.Atoi(stone[len(stone)/2:])
	return []int{left, right}
}

func blink(stones []int) []int {
	var stns []int

	for _, stone := range stones {
		if stone == 0 {
			stns = append(stns, 1)
			continue
		}

		s := fmt.Sprintf("%d", stone)
		if len(s)%2 == 0 {
			stns = append(stns, splitStone(s)...)
			continue
		}

		stns = append(stns, stone*2024)
	}

	return stns
}

func SolveDay11Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "11.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	engravings := strings.Split(lines[0], " ")

	stones := make([]int, len(engravings))
	for i, engraving := range engravings {
		n, _ := strconv.Atoi(engraving)
		stones[i] = n
	}

	for range 25 {
		stones = blink(stones)
	}

	return fmt.Sprintf("%d", len(stones))
}

func SolveDay11Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "11.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	engravings := strings.Split(lines[0], " ")

	stones := make([]int, len(engravings))
	for i, engraving := range engravings {
		n, _ := strconv.Atoi(engraving)
		stones[i] = n
	}

	var stns []int
	for i, stone := range stones {
		stn := []int{stone}
		for j := range 75 {
			fmt.Println(i, j)
			stn = blink(stn)
		}
		stns = append(stns, stn...)
	}

	return fmt.Sprintf("%d", len(stns))
}

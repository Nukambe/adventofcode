package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"slices"
)

func reportIsSafe(levels []int) bool {
	asc := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		thisLevel := levels[i]
		nextLevel := levels[i+1]

		if thisLevel == nextLevel || helpers.AbsoluteInt(thisLevel-nextLevel) > 3 {
			return false
		}

		if (asc && thisLevel > nextLevel) || (!asc && thisLevel < nextLevel) {
			return false
		}
	}
	return true
}

func SolveDay2Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "2-1.txt")
	lines, err := helpers.ReadPuzzleInput(puzzlePath)
	if err != nil {
		fmt.Println("unable to open puzzle input:", puzzlePath)
	}
	reports := helpers.LinesToInts(lines)

	safe := len(reports)
	for _, report := range reports {
		if !reportIsSafe(report) {
			safe--
		}
	}
	return fmt.Sprintf("%d", safe)
}

func SolveDay2Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "2-1.txt")
	lines, err := helpers.ReadPuzzleInput(puzzlePath)
	if err != nil {
		fmt.Println("unable to open puzzle input:", puzzlePath)
	}
	reports := helpers.LinesToInts(lines)

	safe := 0
	for _, report := range reports {
		cont := false
		for i := 0; i < len(report)-1; i++ {
			if reportIsSafe(slices.Concat(report[:i], report[i+1:])) {
				safe++
				cont = true
				break
			}
		}
		if cont {
			continue
		}
		if reportIsSafe(report[:len(report)-1]) {
			safe++
		}
	}
	return fmt.Sprintf("%d", safe)
}

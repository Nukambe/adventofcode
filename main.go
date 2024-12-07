package main

import (
	"fmt"
	_024 "github.com/Nukambe/adventofcode/2024"
	"github.com/Nukambe/adventofcode/helpers"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("usage: <year> <day> <puzzle>")
		return
	}

	puzzles := helpers.Puzzles{}
	puzzles.RegisterPuzzle("2024", "1", 0, _024.SolveDay1Puzzle1)
	puzzles.RegisterPuzzle("2024", "1", 1, _024.SolveDay1Puzzle2)
	puzzles.RegisterPuzzle("2024", "2", 0, _024.SolveDay2Puzzle1)
	puzzles.RegisterPuzzle("2024", "2", 1, _024.SolveDay2Puzzle2)
	puzzles.RegisterPuzzle("2024", "3", 0, _024.SolveDay3Puzzle1)
	puzzles.RegisterPuzzle("2024", "3", 1, _024.SolveDay3Puzzle2)
	puzzles.RegisterPuzzle("2024", "4", 0, _024.SolveDay4Puzzle1)
	puzzles.RegisterPuzzle("2024", "4", 1, _024.SolveDay4Puzzle2)
	puzzles.RegisterPuzzle("2024", "5", 0, _024.SolveDay5Puzzle1)
	puzzles.RegisterPuzzle("2024", "5", 1, _024.SolveDay5Puzzle2)
	puzzles.RegisterPuzzle("2024", "6", 0, _024.SolveDay6Puzzle1)
	puzzles.RegisterPuzzle("2024", "6", 1, _024.SolveDay6Puzzle2)
	puzzles.RegisterPuzzle("2024", "7", 0, _024.SolveDay7Puzzle1)
	puzzles.RegisterPuzzle("2024", "7", 1, _024.SolveDay7Puzzle2)

	answer, err := puzzles.SolvePuzzle(args[1], args[2], args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(answer)
}

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

	answer, err := puzzles.SolvePuzzle(args[1], args[2], args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(answer)
}

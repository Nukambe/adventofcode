package main

import (
	"fmt"
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
	register2024(puzzles)

	answer, err := puzzles.SolvePuzzle(args[1], args[2], args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(answer)
}

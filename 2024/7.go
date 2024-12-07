package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
	"strings"
)

func testEquation(y int, equation []int) bool {
	if len(equation) < 3 {
		return y == equation[0]+equation[1] ||
			y == equation[0]*equation[1]
	}
	return testEquation(y, append([]int{equation[0] + equation[1]}, equation[2:]...)) ||
		testEquation(y, append([]int{equation[0] * equation[1]}, equation[2:]...))
}

func testEquation2(y int, equation []int) bool {
	if len(equation) < 3 {
		return y == equation[0]+equation[1] ||
			y == equation[0]*equation[1] ||
			fmt.Sprintf("%d", y) == fmt.Sprintf("%d%d", equation[0], equation[1])
	}
	i, _ := strconv.Atoi(fmt.Sprintf("%d%d", equation[0], equation[1]))
	return testEquation2(y, append([]int{equation[0] + equation[1]}, equation[2:]...)) ||
		testEquation2(y, append([]int{equation[0] * equation[1]}, equation[2:]...)) ||
		testEquation2(y, append([]int{i}, equation[2:]...))
}

func SolveDay7Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "7.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for _, line := range lines {
		equation := strings.Split(line, ": ")
		y, _ := strconv.Atoi(equation[0])
		var nums []int
		for _, n := range strings.Split(equation[1], " ") {
			i, _ := strconv.Atoi(n)
			nums = append(nums, i)
		}
		if testEquation(y, nums) {
			total += y
		}
	}
	return fmt.Sprintf("%d", total)
}

func SolveDay7Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "7.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for _, line := range lines {
		equation := strings.Split(line, ": ")
		y, _ := strconv.Atoi(equation[0])
		var nums []int
		for _, n := range strings.Split(equation[1], " ") {
			i, _ := strconv.Atoi(n)
			nums = append(nums, i)
		}
		if testEquation2(y, nums) {
			total += y
		}
	}
	return fmt.Sprintf("%d", total)
}

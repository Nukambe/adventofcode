package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"sort"
	"strconv"
	"strings"
)

func splitLists() (left []int, right []int) {
	puzzlePath := path.Join("2024", "puzzle-inputs", "1-1.txt")
	lines, err := helpers.ReadPuzzleInput(puzzlePath)
	if err != nil {
		fmt.Println("unable to open puzzle input:", puzzlePath)
	}

	left = make([]int, len(lines))
	right = make([]int, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, "   ")

		n, _ := strconv.Atoi(nums[0])
		left[i] = n

		n, _ = strconv.Atoi(nums[1])
		right[i] = n
	}
	return
}

func SolveDay1Puzzle1() string {
	left, right := splitLists()

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += helpers.AbsoluteInt(left[i] - right[i])
	}

	return fmt.Sprintf("%d", distance)
}

func SolveDay1Puzzle2() string {
	left, right := splitLists()

	rightCount := map[int]int{}
	for _, r := range right {
		if _, ok := rightCount[r]; !ok {
			rightCount[r] = 0
		}
		rightCount[r]++
	}

	similarity := 0
	for _, l := range left {
		count, ok := rightCount[l]
		if ok {
			similarity += l * count
		}
	}

	return fmt.Sprintf("%d", similarity)
}

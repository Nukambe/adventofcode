package helpers

import (
	"strconv"
	"strings"
)

func LinesToInts(lines []string) [][]int {
	nums := make([][]int, len(lines))
	for i, line := range lines {
		texts := strings.Split(line, " ")
		nums[i] = make([]int, len(texts))
		for j, text := range texts {
			n, _ := strconv.Atoi(text)
			nums[i][j] = n
		}
	}
	return nums
}

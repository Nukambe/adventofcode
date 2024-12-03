package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
	"strings"
	"unicode"
)

func multiply(mul string) int { // found 0 nested muls...
	nums := strings.Split(mul, ",")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	return a * b
}

func SolveDay3Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "3-1.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	for _, line := range lines {
		i := 0
		j := 0
		for i < len(line)-4 {
			if line[i:i+4] == "mul(" { // [mul(]0,0)
				j = i + 3 // mul([0],0)
				i += 4    // mul(0[,]0)
				open := 1 // mul[(]0,0)
				for i < len(line) {
					r := rune(line[i])
					if unicode.IsDigit(r) { // mul([0],[0])
						i++
						continue
					} else if r == ')' { // mul(0,0[)]
						open--
						if open == 0 {
							total += multiply(line[j+1 : i]) // mul([0,0])
							break
						} else if open < 0 { // mul(0,mul)))))
							break
						} else {
							i++
							continue
						}
					} else if r == ',' { // mul(0[,]0)
						if unicode.IsDigit(rune(line[i+1])) { // mul(0,[0])
							i += 2
							continue
						} else if line[i+1:i+5] == "mul(" { // mul(0,[mul(]0,0))
							i += 5
							open++
							continue
						} else {
							break
						}
					} else if r == 'm' { // mul([m]ul(0,0),0)
						if line[i:i+4] == "mul(" { // mul([mul(]0,0),mul(0,0))
							i += 5
							open++
							continue
						} else {
							break
						}
					} else {
						break
					}
				}
			}
			i++
		}
	}
	return fmt.Sprintf("%d", total)
}

func SolveDay3Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "3-1.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	total := 0

	disabled := false
	for _, line := range lines {
		i := 0
		j := 0
		for i < len(line)-4 {
			if i+4 < len(line) {
				if line[i:i+4] == "do()" {
					disabled = false
					i += 4
				}
			}
			if i+7 < len(line) {
				if line[i:i+7] == "don't()" {
					disabled = true
					i += 7
				}
			}
			if !disabled && line[i:i+4] == "mul(" { // [mul(]0,0)
				j = i + 3 // mul([0],0)
				i += 4    // mul(0[,]0)
				open := 1 // mul[(]0,0)
				for i < len(line) {
					r := rune(line[i])
					if unicode.IsDigit(r) { // mul([0],[0])
						i++
						continue
					} else if r == ')' { // mul(0,0[)]
						open--
						if open == 0 {
							total += multiply(line[j+1 : i]) // mul([0,0])
							break
						} else if open < 0 { // mul(0,mul)))))
							break
						} else {
							i++
							continue
						}
					} else if r == ',' { // mul(0[,]0)
						if unicode.IsDigit(rune(line[i+1])) { // mul(0,[0])
							i += 2
							continue
						} else if line[i+1:i+5] == "mul(" { // mul(0,[mul(]0,0))
							i += 5
							open++
							continue
						} else {
							break
						}
					} else if r == 'm' { // mul([m]ul(0,0),0)
						if line[i:i+4] == "mul(" { // mul([mul(]0,0),mul(0,0))
							i += 5
							open++
							continue
						} else {
							break
						}
					} else {
						break
					}
				}
			}
			i++
		}
	}
	return fmt.Sprintf("%d", total)
}

package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
)

func getFileBlocks(diskMap string) []int {
	var block []int
	id := 0

	for i, f := range diskMap {
		n, _ := strconv.Atoi(string(f))
		for _ = range n {
			if i%2 == 0 {
				block = append(block, id)
				continue
			}
			block = append(block, -1)
		}
		if i%2 == 0 {
			id++
		}
	}

	return block
}

func SolveDay9Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "9.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	diskMap := lines[0]
	checksum := 0

	fileBlocks := getFileBlocks(diskMap)
	left := 0
	right := len(fileBlocks) - 1
	for left < right {
		if fileBlocks[left] == -1 {
			for fileBlocks[right] == -1 {
				right--
			}
			if right > left {
				fileBlocks[left] = fileBlocks[right]
				fileBlocks[right] = -1
			}
		}
		if fileBlocks[left] != -1 {
			checksum += left * fileBlocks[left]
		}
		left++
	}

	return fmt.Sprintf("%d", checksum)
}

func findNextFile(i int, fileBlocks []int) (int, []int) {
	for fileBlocks[i] == -1 {
		i--
		if i <= 0 {
			return 0, nil
		}
	}
	n := 0
	id := fileBlocks[i]
	for i > -1 {
		if fileBlocks[i] != id {
			break
		}
		n++
		i--
	}
	i++
	return i, fileBlocks[i : i+n]
}

func findFreeSpace(fileIndex, space int, fileBlocks []int) []int {
	i := 0
	for i < fileIndex {
		if fileBlocks[i] != -1 {
			i++
			continue
		}
		n := 0
		for fileBlocks[i] == -1 {
			n++
			i++
			if n == space {
				return fileBlocks[i-n : i+1]
			}
		}
		i++
	}
	return nil
}

func SolveDay9Puzzle2() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "9.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	diskMap := lines[0]
	checksum := 0

	fileBlocks := getFileBlocks(diskMap)
	i := len(fileBlocks) - 1

	for i > -1 {
		fileIndex, file := findNextFile(i, fileBlocks)
		if fileIndex <= 0 {
			break
		}
		fileSpace := len(file)
		free := findFreeSpace(i-fileSpace, fileSpace, fileBlocks)
		if free != nil {
			for n, _ := range file {
				free[n] = file[n]
				file[n] = -1
			}
		}

		i = fileIndex
		i--
	}

	for i := 0; i < len(fileBlocks); i++ {
		if fileBlocks[i] == -1 {
			continue
		}
		checksum += i * fileBlocks[i]
	}

	return fmt.Sprintf("%d", checksum)
}

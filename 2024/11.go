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
	puzzlePath := path.Join("2024", "puzzle-inputs", "test.txt")
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

	return fmt.Sprintf("%v", stones)
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

	fmt.Println("working to 40...")
	for range 40 {
		stones = blink(stones)
	}

	var stns []int
	stoneCh := make(chan []int)
	for _, stone := range stones {
		go func() {
			stn := []int{stone}
			for range 35 {
				stn = blink(stn)
			}
			stoneCh <- stn
		}()
	}

	i := 0
	go func() {
		for addStones := range stoneCh {
			stns = append(stns, addStones...)
			i++
			fmt.Println(i)
		}
	}()

	for {
		if i == len(stones) {
			break
		}
	}

	return fmt.Sprintf("%d", len(stns))
}

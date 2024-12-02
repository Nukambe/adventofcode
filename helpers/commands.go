package helpers

import (
	"fmt"
	"strconv"
)

type Puzzles map[string]map[string][]func() string

func (years Puzzles) RegisterPuzzle(year, day string, puzzle int, f func() string) {
	_, ok := years[year]
	if !ok {
		years[year] = map[string][]func() string{}
	}

	_, ok = years[year][day]
	if !ok {
		years[year][day] = make([]func() string, 2)
	}

	years[year][day][puzzle] = f
}

func (years Puzzles) SolvePuzzle(year, day, puzzle string) (string, error) {
	puzzleMap, ok := years[year]
	if !ok {
		return "", fmt.Errorf("invalid year: %s", year)
	}

	i, err := strconv.Atoi(puzzle)
	if err != nil {
		return "", err
	}

	var solvers []func() string
	solvers, ok = puzzleMap[day]
	if !ok {
		return "", fmt.Errorf("invalid day: %s", day)
	}

	solver := solvers[i-1]
	answer := solver()
	return answer, nil
}

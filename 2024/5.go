package _024

import (
	"fmt"
	"github.com/Nukambe/adventofcode/helpers"
	"path"
	"strconv"
	"strings"
)

type ruleSet map[string][]string

func readRulesAndPages(lines []string) (ruleSet, [][]string) {
	rules := ruleSet{}
	updates := make([][]string, 0)

	for _, line := range lines {
		pages := strings.Split(line, "|")
		if len(pages) == 2 { // is this a rule?
			_, ok := rules[pages[0]]
			if !ok {
				rules[pages[0]] = make([]string, 1)
			}
			rules[pages[0]] = append(rules[pages[0]], pages[1])
		} else { // no, these are these pages
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return rules, updates
}

func checkRuleset(ruleset ruleSet, update []string) bool {
	pages := map[string]bool{}
	for _, page := range update {
		pages[page] = true
		rules, ok := ruleset[page]
		if !ok { // no rules, ignore
			continue
		}
		for _, rule := range rules {
			if _, ok = pages[rule]; ok {
				return false
			}
		}
	}
	return true
}

func reorderUpdate(ruleset ruleSet, update []string) {
	for i := 0; i < len(update)-1; i++ {
		a := update[i]
		b := update[i+1]
		rules := ruleset[b]

		if n := helpers.SliceContains(rules, a); n > -1 {
			update[i] = b
			update[i+1] = a
		}
	}
	if !checkRuleset(ruleset, update) {
		reorderUpdate(ruleset, update)
	}
}

func SolveDay5Puzzle1() string {
	puzzlePath := path.Join("2024", "puzzle-inputs", "5.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	rules, updates := readRulesAndPages(lines)
	total := 0

	for _, update := range updates {
		if !checkRuleset(rules, update) {
			continue
		}
		mid, _ := strconv.Atoi(update[len(update)/2])
		total += mid
	}
	return fmt.Sprintf("%d", total)
}

func SolveDay5Puzzle2() string { // ~3800 swaps
	puzzlePath := path.Join("2024", "puzzle-inputs", "5.txt")
	lines, _ := helpers.ReadPuzzleInput(puzzlePath)
	rules, updates := readRulesAndPages(lines)
	total := 0

	for _, update := range updates {
		if checkRuleset(rules, update) {
			continue
		}

		reorderUpdate(rules, update)

		mid, _ := strconv.Atoi(update[len(update)/2])
		total += mid
	}

	return fmt.Sprintf("%d", total)
}

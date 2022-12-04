package day3

import (
	"sort"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func groupElves(input []string, groupSize int) [][]string {
	elves := [][]string{}

	for {
		if len(input) == 0 {
			break
		}

		tmp := input[0:groupSize]
		sort.Slice(tmp, func(i, j int) bool {
			return len(tmp[i]) < len(tmp[j])
		})

		elves = append(elves, tmp)
		input = input[groupSize:]
	}
	return elves
}

func findCommons(elves [][]string) []string {
	all := []string{}
	for _, group := range elves {
		all = append(all, findCommonsPerGroup(group)...)
	}
	return all
}

func findCommonsPerGroup(elves []string) []string {
	commons := map[string]bool{}
	first := elves[0]
	for _, char := range first {
		if strings.IndexByte(elves[1], byte(char)) != -1 && strings.IndexByte(elves[2], byte(char)) != -1 {
			commons[string(char)] = true
		}
	}
	return maps.Keys(commons)
}

// find shortest string
// check which val exists in both others
// get Value of that

func findDups(lines []string) []string {
	dups := []string{}
	for _, val := range lines {
		first, second := val[:len(val)/2], val[len(val)/2:]

		intermediate := []string{}
		for _, a := range first {
			if strings.Contains(second, string(a)) && !slices.Contains(intermediate, string(a)) {
				intermediate = append(intermediate, string(a))
			}
		}
		dups = append(dups, intermediate...)
	}
	return dups
}

func translateValue(val string) int {
	_rune := int(val[0])
	if _rune >= 97 && _rune <= 122 {
		return _rune - 96
	} else if _rune >= 65 && _rune <= 90 {
		return _rune - 38
	}
	return 0
}

func toPrio(dups []string) int {
	totals := 0
	for _, k := range dups {
		totals += translateValue(k)
	}
	return totals
}

// a-z 97-122 1-26
// A-Z 65-90 27-52

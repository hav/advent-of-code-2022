package day4

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"besharati.se/advent-of-code-2022/utils"
)

func convertToInt(s string) int {
	x, err := strconv.Atoi(s)
	utils.Check(err)
	return x
}

func createArea2(input string) []int {
	area := strings.Split(input, "-")
	tmp := []int{}

	for i := convertToInt(area[0]); i < convertToInt(area[1]); i++ {
		tmp = append(tmp, i)
	}

	return tmp
}

func overlaps2(input []string) int {
	overlaps := 0
	for _, inputPair := range input {
		pair := strings.Split(inputPair, ",")
		a1, a2 := createArea2(pair[0]), createArea2(pair[1])

		fmt.Println(a1, a2)
		fmt.Printf("start: %d\n", sort.SearchInts(a1, a2[0]))
		fmt.Printf("end: %d\n", sort.SearchInts(a1, a2[len(a2)-1]))
	}

	return overlaps
}

type Area struct {
	start string
	end   string
}

func createArea(input string) Area {
	area := strings.Split(input, "-")

	return Area{start: area[0], end: area[1]}
}

func (a *Area) contains(b Area) bool {
	return a.start >= b.start && a.end <= b.end
}

func overlaps(input []string) int {
	overlaps := 0
	for _, inputPair := range input {
		pair := strings.Split(inputPair, ",")
		a1, a2 := createArea(pair[0]), createArea(pair[1])
		if a1.contains(a2) || a2.contains(a1) {
			overlaps += 1
		}
	}

	return overlaps
}

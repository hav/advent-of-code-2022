package day1

import (
	"fmt"
	"sort"
	"strconv"

	utils "besharati.se/advent-of-code-2022/utils"
)

func buildElves() []int {
	input := utils.ReadFileAsSlice("input_1_1.txt")

	elves := []int{}
	elfItems := []int{}
	for _, txt := range input {
		if txt == "" {
			elves = append(elves, utils.SumIntSlice(elfItems))
			elfItems = []int{}
			continue
		}

		val, err := strconv.Atoi(txt)
		utils.Check(err)
		elfItems = append(elfItems, val)
	}
	return elves
}

func day1() {
	elves := buildElves()
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1:][0])
	fmt.Println(utils.SumIntSlice(elves[len(elves)-3:]))
}

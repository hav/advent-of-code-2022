package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	utils "besharati.se/advent-of-code-2022/utils"
)

func sumIntSlice(slice []int) int {
	total := 0
	for _, s := range slice {
		total += s
	}
	return total
}

func buildElves() []int {

	input, err := os.Open("input_1_1.txt")
	utils.Check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	elves := []int{}
	elfItems := []int{}
	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			elves = append(elves, sumIntSlice(elfItems))
			elfItems = []int{}
			continue
		}

		val, err := strconv.Atoi(txt)
		utils.Check(err)
		elfItems = append(elfItems, val)
	}
	return elves
}

func main() {
	elves := buildElves()
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1:][0])
	fmt.Println(sumIntSlice(elves[len(elves)-3:]))
}

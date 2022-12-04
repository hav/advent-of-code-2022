package utils

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func SumIntSlice(slice []int) int {
	total := 0
	for _, s := range slice {
		total += s
	}
	return total
}

func ReadFile(fileName string) []string {
	input, err := os.Open(fileName)
	Check(err)
	defer input.Close()

	lines := []string{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

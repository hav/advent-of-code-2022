package day4

import (
	"testing"

	"besharati.se/advent-of-code-2022/utils"
)

func TestDay4(t *testing.T) {

	t.Run("test input", func(t *testing.T) {
		input := []string{
			"2-4,6-8", "2-3,4-5",
			"5-7,7-9", "2-8,3-7",
			"6-6,4-6", "2-6,4-8",
		}

		overlap := overlaps2(input)
		if overlap != 2 {
			t.Errorf("was %d", overlap)
		}
	})

	t.Run("first puzzle", func(t *testing.T) {
		input := utils.ReadFile("input_4_1.txt")

		overlap := overlaps(input)
		// this should fail and print the actual value
		if overlap != 482 {
			t.Errorf("was %d", overlap)
		}
	})
}

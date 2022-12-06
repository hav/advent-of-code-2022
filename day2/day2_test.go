package day2

import (
	"testing"

	"besharati.se/advent-of-code-2022/utils"
)

func TestDay2(t *testing.T) {
	t.Run("test data", func(t *testing.T) {
		input := []string{"A Y", "B X", "C Z"}

		result := rockPaperScissors(input, false)

		if result != 15 {
			t.Errorf("was: %d\n", result)
		}
	})

	t.Run("first puzzle", func(t *testing.T) {
		input := utils.ReadFileAsSlice("input_2_1.txt")
		result := rockPaperScissors(input, false)

		if result != 9651 {
			t.Errorf("was: %d\n", result)
		}
	})

	t.Run("second puzzle test data", func(t *testing.T) {
		input := []string{"A Y", "B X", "C Z"}
		result := rockPaperScissors(input, true)

		if result != 12 {
			t.Errorf("was: %d\n", result)
		}
	})

	t.Run("second puzzle", func(t *testing.T) {
		input := utils.ReadFileAsSlice("input_2_1.txt")
		result := rockPaperScissors(input, true)

		if result != 10560 {
			t.Errorf("was: %d\n", result)
		}
	})
}

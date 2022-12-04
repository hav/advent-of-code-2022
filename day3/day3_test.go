package day3

import (
	"testing"

	"besharati.se/advent-of-code-2022/utils"
)

func TestDay3(t *testing.T) {
	t.Run("first puzzle test input", func(t *testing.T) {
		input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw"}
		dups := findDups(input)
		val := toPrio(dups)
		if val != 157 {
			t.Errorf("was %d", val)
		}
	})

	t.Run("first puzzle", func(t *testing.T) {
		input := utils.ReadFile("input_3_1.txt")
		dups := findDups(input)
		val := toPrio(dups)

		if val != 8123 {
			t.Errorf("Was %d", val)
		}
	})

	t.Run("second puzzle test data", func(t *testing.T) {
		input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}

		groups := groupElves(input, 3)
		commons := findCommons(groups)
		prios := toPrio(commons)
		if prios != 70 {
			t.Error()
		}
	})

	t.Run("second puzzle", func(t *testing.T) {
		input := utils.ReadFile("input_3_1.txt")
		groups := groupElves(input, 3)
		commons := findCommons(groups)
		prios := toPrio(commons)
		if prios != 2620 {
			t.Errorf("was %d\n", prios)
		}
	})
}

func BenchmarkFindCommons(b *testing.B) {
	input := utils.ReadFile("input_3_1.txt")
	groups := groupElves(input, 3)
	b.Run("findCommons", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			findCommons(groups)
		}
	})
}

func BenchmarkFindDups(b *testing.B) {
	input := utils.ReadFile("input_3_1.txt")

	b.Run("findDups", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			findDups(input)
		}
	})
}

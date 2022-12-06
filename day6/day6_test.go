package day6

import (
	"testing"

	"besharati.se/advent-of-code-2022/utils"
)

func TestDay6(t *testing.T) {
	firstMarkerLength := 4
	secondMarkerLength := 14

	t.Run("test data", func(t *testing.T) {
		testData := map[string]int{
			"bvwbjplbgvbhsrlpgdmjqwftvncz": 5, "nppdvjthqldpwncqszvftbrmjlhg": 6, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": 11,
		}
		for buffer, markerIndex := range testData {
			index := findMarkerIndex(buffer, firstMarkerLength)
			if index != markerIndex {
				t.Errorf("wanted %d, got %d", markerIndex, index)
			}
		}
	})

	t.Run("first puzzle", func(t *testing.T) {
		buffer := utils.ReadFile("input_6_1.txt")
		index := findMarkerIndex(buffer, firstMarkerLength)
		if index != 1623 {
			t.Errorf("got %d", index)
		}
	})

	t.Run("second puzzle test", func(t *testing.T) {
		testData := map[string]int{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
			"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
			"nppdvjthqldpwncqszvftbrmjlhg":      23,
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
		}

		for buffer, markerIndex := range testData {
			index := findMarkerIndex(buffer, secondMarkerLength)
			if index != markerIndex {
				t.Errorf("wanted %d, got %d", markerIndex, index)
			}
		}
	})

	t.Run("second puzzle", func(t *testing.T) {
		buffer := utils.ReadFile("input_6_1.txt")
		index := findMarkerIndex(buffer, secondMarkerLength)
		if index != 3774 {
			t.Errorf("got %d", index)
		}
	})
}

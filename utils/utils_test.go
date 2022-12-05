package utils

import "testing"

func TestSumIntSlice(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		res := SumIntSlice([]int{1, 2, 3})

		if res != 6 {
			t.Error()
		}
	})
}

func BenchmarkSumIntSlice(b *testing.B) {
	for i := 0; i < 100000; i++ {
		SumIntSlice([]int{1, 2, 3})
	}
}

// 0.0002311 ns/op for loop

// 0.0000232 ns/op

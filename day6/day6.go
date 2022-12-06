package day6

func unique(s string) bool {
	m := map[rune]bool{}

	for _, v := range s {
		if !m[v] {
			m[v] = true
		} else {
			return false
		}
	}

	return true
}

func findMarkerIndex(buffer string, markerLength int) int {
	for i := markerLength; i < len(buffer); i++ {
		// fmt.Printf("buffer %s, %t, index %d\n", string(buffer[i-4:i]), unique(string(buffer[i-4:i])), i)
		if unique(string(buffer[i-markerLength : i])) {
			return i
		}
	}
	return 0
}

package strings

// Checks if b is a permutation of a
func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	var strInfo map[rune]int
	strInfo = make(map[rune]int)

	for _, cr := range a {
		n := strInfo[cr]
		strInfo[cr] = n + 1
	}

	for _, cr := range b {
		n := strInfo[cr]
		if n == 0 {
			// This means that cr not same kind of chars
			return false
		}
		if n == 1 {
			delete(strInfo, cr)
		} else {
			strInfo[cr] = n - 1
		}
	}
	return len(strInfo) == 0
}

package lev

// Calculate the Levenshtein distance between given strings
// Uses iterative matrix algorithm
func Calc(a, b string) int {
	var (
		m [][]int
		c int
	)

	// Build initial matrix
	m = make([][]int, len(a)+1)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, len(b)+1)
		if i == 0 {
			for j := 0; j < len(m[i]); j++ {
				m[i][j] = j
			}
		} else {
			m[i][0] = i
		}
	}

	// Calculate distance matrix
	for ai := 1; ai < len(m); ai++ {
		for bi := 1; bi < len(m[ai]); bi++ {

			if a[ai-1] == b[bi-1] {
				c = 0
			} else {
				c = 1
			}

			del := m[ai-1][bi] + 1
			ins := m[ai][bi-1] + 1
			sub := m[ai-1][bi-1] + c

			m[ai][bi] = min(del, ins, sub)
		}
	}

	return m[len(m)-1][len(m[0])-1]
}

package chapter1

// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0.

func ZeroColumnRow(m [][]int) [][]int {

	row := make([]bool, len(m))
	col := make([]bool, len(m[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if row[i] == true || col[j] {
				m[i][j] = 0
			}
		}
	}
	return m
}

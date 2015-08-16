package chapter1

import (
	"testing"
)

type zeroColumnRowTestPair struct {
	orig   [][]int
	result [][]int
}

var zeroColumnRowTests = []zeroColumnRowTestPair{
	{[][]int{[]int{2, 2, 2, 2, 2},
		[]int{2, 2, 3, 0, 3},
		[]int{2, 2, 4, 3, 4},
		[]int{2, 2, 0, 1, 0}},

		[][]int{[]int{2, 2, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{2, 2, 0, 0, 0},
			[]int{0, 0, 0, 0, 0}}},
}

func TestZeroColumnRow(t *testing.T) {
	for _, test := range zeroColumnRowTests {
		elem := ZeroColumnRow(test.orig)
		for i := 0; i < len(test.result); i++ {
			for j := 0; j < len(test.result[0]); j++ {
				if test.result[i][j] != elem[i][j] {
					t.Error("For ", test.orig,
						"Expected", test.result,
						"got", elem)
				}
			}
		}
	}
}

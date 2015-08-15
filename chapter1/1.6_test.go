package chapter1

import (
	"testing"
)

type rotateTestPair struct {
	orig   [4][4]int
	result [4][4]int
}

var rotateTests = []rotateTestPair{
	{[4][4]int{[4]int{0, 1, 2, 3},
		[4]int{4, 5, 6, 7},
		[4]int{8, 9, 10, 11},
		[4]int{12, 13, 14, 15}},

		[4][4]int{[4]int{12, 8, 4, 0},
			[4]int{13, 9, 5, 1},
			[4]int{14, 10, 6, 2},
			[4]int{15, 11, 7, 3}}},
}

func TestRotate(t *testing.T) {
	for _, test := range rotateTests {

		if elem := Rotate(test.orig); elem != test.result {
			t.Error("For ", test.orig,
				"Expected", test.result,
				"got", elem)
		}
	}
}

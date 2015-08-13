package chapter1

import (
	"testing"
)

type IsAllUniqueTestPair struct {
	s      string
	result bool
}

var isAllUniqueTests = []IsAllUniqueTestPair{
	{"murat", true},
	{"yildirim", false},
	{"12345678a~'`ZXVBNM[]{}:?<>", true},
}

func TestIsAllUnique(t *testing.T) {
	for _, test := range isAllUniqueTests {
		if elem := IsAllUnique(test.s); elem != test.result {
			t.Error("For ", test.s,
				"Expected", test.result,
				"got", elem)
		}
	}
}

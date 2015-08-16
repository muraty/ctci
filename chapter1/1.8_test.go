package chapter1

import (
	"testing"
)

type isRotationTestPair struct {
	orig    string
	rotated string
	result  bool
}

var isRotationTests = []isRotationTestPair{
	{"murat", "atmur", true},
	{"murat", "murat", true},
	{"murat", "tarum", false},
	{"golang", " golang", false},
	{"golang", "golang  ", false},
}

func TestIsRotation(t *testing.T) {
	for _, test := range isRotationTests {
		if elem := IsRotation(test.orig, test.rotated); elem != test.result {
			t.Error("For ", test.orig, "and", test.rotated,
				"Expected", test.result,
				"got", false)
		}
	}
}

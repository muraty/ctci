package chapter1

import (
	"testing"
)

type permutationTestPair struct {
	source string
	target string
	result bool
}

var permutationTests = []permutationTestPair{
	{"abc", "cba", true},
	{"golang", "langgo", true},
	{"gopher", "Gopher    ", false},
	{"python", "python ", false},
}

func TestIsPermutation(t *testing.T) {
	for _, test := range permutationTests {
		if r := isPermutation(test.source, test.target); r != test.result {
			t.Error("For ", test.source, "and", test.target,
				"Expected", test.result,
				"got", r)
		}
	}
}

package chapter1

import (
	"testing"
)

type reverseTestPair struct {
	str    string
	result string
}

var reverseTests = []reverseTestPair{
	{"murat", "tarum"},
	{"omer", "remo"},
	{"yildirim", "miridliy"},
}

func TestReverse(t *testing.T) {
	for _, test := range reverseTests {
		orig := test.str
		reverse(&(test.str))
		if test.str != test.result {
			t.Error("For ", orig,
				"Expected", test.result,
				"got", test.str)
		}
	}
}

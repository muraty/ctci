package chapter1

import (
	"testing"
)

type replaceSpacesTestPair struct {
	orig   string
	result string
}

var replaceSpacesTests = []replaceSpacesTestPair{

	{"Mr John Smith      ", "Mr%20John%20Smith"},
	{"    Mr John Smith      ", "%20%20%20%20Mr%20John%20Smith"},
	{"Murat", "Murat"},
}

func TestReplaceSpaces(t *testing.T) {
	for _, test := range replaceSpacesTests {
		orig := test.orig
		ReplaceSpaces(&test.orig)
		if test.orig != test.result {
			t.Error("For ", orig,
				"Expected", test.result,
				"got", test.orig)
		}
	}
}

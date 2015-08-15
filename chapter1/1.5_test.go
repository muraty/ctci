package chapter1

import (
	"testing"
)

type compressionTestPair struct {
	orig   string
	result string
}

var compressionTests = []compressionTestPair{
	{"murat", "murat"},
	{"aabcccccaaa", "a2b1c5a3"},
	{"aAbcCcccaaaaaa", "a1A1b1c1C1c3a6"},
	{"muuurrrrat", "m1u3r4a1"},
}

func TestCompression(t *testing.T) {
	for _, test := range compressionTests {

		if elem := Compression(test.orig); elem != test.result {
			t.Error("For ", test.orig,
				"Expected", test.result,
				"got", elem)
		}
	}
}

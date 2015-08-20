package chapter2

import (
	datastructures "ctci/datastructures"
	"testing"
)

type sliceTestPair struct {
	orig   datastructures.LinkedList
	result *datastructures.LinkedList
	kth    int
}

func initializeSlice() []sliceTestPair {
	orig := datastructures.LinkedList{}
	for i := 0; i < 10; i++ {
		n := &datastructures.Node{}
		n.SetValue(i)
		orig.Add(n)
	}

	target := datastructures.LinkedList{}
	for i := 3; i < 10; i++ {
		n := &datastructures.Node{}
		n.SetValue(i)
		target.Add(n)
	}
	var sliceTests = []sliceTestPair{
		{orig, &target, 3},
		{orig, nil, 30},
		{orig, nil, -3},
	}
	return sliceTests
}

func TestSlice(t *testing.T) {
	sliceTests := initializeSlice()
	for _, test := range sliceTests {
		elem := Slice(test.orig, test.kth)
		if elem == nil {
			if elem == test.result {
				continue
			} else {
				t.Error("For ", test.orig.GetElements(),
					"Expected", test.result.GetElements(),
					"got", elem.GetElements())
			}
		} else {
			for i := test.kth; i < elem.Len(); i++ {
				// from kth to last element must be identical for both of
				// test.result and return value of Slice() function.
				if elem.Find(i).GetValue() != test.result.Find(i).GetValue() {
					t.Error("For ", test.orig.GetElements(),
						"Expected", test.result.GetElements(),
						"got", elem.GetElements())
				}
				// from 0 to kth element must be nil since we want to slice
				// from kth to last element only.
				for i := 0; i < test.kth; i++ {
					if test.result.Find(i) != nil {
						t.Error("For ", test.orig.GetElements(),
							"Expected", test.result.GetElements(),
							"got", elem.GetElements())
					}

				}
			}
		}
	}
}

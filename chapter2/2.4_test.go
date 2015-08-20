package chapter2

import (
	datastructures "ctci/datastructures"
	"testing"
)

type partitionListTestPair struct {
	orig   datastructures.LinkedList
	result datastructures.LinkedList
	value  int
}

func initializePartitionList() []partitionListTestPair {
	orig := datastructures.LinkedList{}
	result := datastructures.LinkedList{}
	value := 4
	for i := 0; i < 10; i++ {
		n := datastructures.Node{}
		n.SetValue(i)

		// Set before values
		if i < value {
			orig.Add(&n)
			// Set after values
		} else {
			result.Add(&n)
		}
	}
	// we can expand test cases
	var partitionListTests = []partitionListTestPair{
		{orig, result, value},
	}
	return partitionListTests
}

func TestPartitionList(t *testing.T) {
	partitionListTests := initializePartitionList()
	for _, test := range partitionListTests {
		partitioned := PartitionList(test.orig, test.value)

		// traverse through the result for function under test
		head := partitioned.Head()

		// test before values
		for i := 0; i < test.value; i++ {
			if head.GetValue() >= test.value {
				t.Error("Before ", test.orig.GetElements(),
					"After", test.result.GetElements(),
					"Value", test.value,
					"Got", partitioned.GetElements(),
					"for before values")
			}
			head = head.Next()
		}
		// test after values
		for i := test.value; i < partitioned.Len(); i++ {
			if head.GetValue() < test.value {
				t.Error("For Before ", test.orig.GetElements(),
					"After", test.result.GetElements(),
					"Value", test.value,
					"Got", partitioned.GetElements(),
					"for after values")
			}
			head = head.Next()
		}
	}
}

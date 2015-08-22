package chapter2

import (
	ds "ctci/datastructures"
	"testing"
)

type findCorruptNodeTestPair struct {
	orig        ds.LinkedList
	corruptNode *ds.Node
}

func initializeLoopedList() []findCorruptNodeTestPair {
	l1 := ds.LinkedList{}
	l2 := ds.LinkedList{}
	node := &ds.Node{}

	// Create a looped linked list
	// 0->1->2->3->4->5->6->3->4->5->6->3->4->5->6...
	for i := 0; i < 6; i++ {
		n := ds.Node{}
		n.SetValue(i)
		l1.Add(&n)

		if i == 5 {
			temp := l1.Head().Next().Next().Next()
			node = temp
			l1.Add(temp)
		}
	}
	// Create a non-looped linked list
	for i := 0; i < 10; i++ {
		n := ds.Node{}
		n.SetValue(i)
		l2.Add(&n)
	}
	// we can expand test cases
	var loopTests = []findCorruptNodeTestPair{
		{l1, node},
		{l2, nil}, // no loop
	}
	return loopTests
}

func TestFindCorruptNode(t *testing.T) {
	loopTests := initializeLoopedList()
	for _, test := range loopTests {
		if elem := FindCorruptNode(test.orig); elem != test.corruptNode {
			t.Error("For list ", test.orig.GetPartialElements(10),
				"Expect", test.corruptNode,
				"Got", elem)
		}
	}
}

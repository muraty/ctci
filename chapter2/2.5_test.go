package chapter2

import (
	datastructures "ctci/datastructures"
	"testing"
)

type sumLinkedListReverseTestPair struct {
	l1     datastructures.LinkedList
	l2     datastructures.LinkedList
	result int
}

type sumLinkedListForwardTestPair struct {
	l1     datastructures.LinkedList
	l2     datastructures.LinkedList
	result int
}

func initializeReverseList() []sumLinkedListReverseTestPair {
	l1 := datastructures.LinkedList{}
	l2 := datastructures.LinkedList{}
	// l2 : 1->2->3    = 321
	// l1 : 1->2->3->4 = 4321
	// summed value    = 4642
	value := 4642

	// Create l1 : 1->2->3->4 = 4321
	for i := 1; i < 5; i++ {
		n := datastructures.Node{}
		n.SetValue(i)
		if i == 4 {
			l1.Add(&n)
			continue
		}
		l1.Add(&n)
	}
	// Create l2 : 1->2->3    = 321
	for i := 1; i < 4; i++ {
		n := datastructures.Node{}
		n.SetValue(i)
		l2.Add(&n)
	}
	// we can expand test cases
	var reverseTests = []sumLinkedListReverseTestPair{
		{l1, l2, value},
	}
	return reverseTests
}
func initializeForwardList() []sumLinkedListForwardTestPair {
	l1 := datastructures.LinkedList{}
	l2 := datastructures.LinkedList{}
	// l1 : 1->2->3->4 = 1234
	// l2 : 1->2->3    = 123
	// summed value    = 1357
	value := 1357

	// Create l1 : 1->2->3->4 = 1234
	for i := 1; i < 5; i++ {
		n := datastructures.Node{}
		n.SetValue(i)

		l1.Add(&n)
	}
	// Create l2 : 1->2->3    = 123
	for i := 1; i < 4; i++ {
		n := datastructures.Node{}
		n.SetValue(i)
		l2.Add(&n)
	}

	// we can expand test cases
	var forwardTests = []sumLinkedListForwardTestPair{
		{l1, l2, value},
	}
	return forwardTests
}
func TestSumLinkedListReverse(t *testing.T) {
	reverseListTests := initializeReverseList()
	for _, test := range reverseListTests {
		if elem := SumLinkedListReverse(test.l1, test.l2); elem != test.result {
			t.Error("l1 ", test.l1.GetElements(), "and",
				"l2", test.l2.GetElements(),
				"Expect", test.result,
				"Got", elem)
		}
	}
}

func TestSumLinkedListForward(t *testing.T) {
	forwardListTests := initializeForwardList()
	for _, test := range forwardListTests {
		if elem := SumLinkedListForward(test.l1, test.l2); elem != test.result {
			t.Error("l1 ", test.l1.GetElements(), "and",
				"l2", test.l2.GetElements(),
				"Expect", test.result,
				"Got", elem)
		}
	}
}

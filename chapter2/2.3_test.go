package chapter2

import (
	datastructures "ctci/datastructures"
	"testing"
)

type deleteNodeTestPair struct {
	orig   datastructures.LinkedList
	result datastructures.LinkedList
	node   datastructures.Node
}

func initializeDeleteNode() []deleteNodeTestPair {
	node := datastructures.Node{}
	orig := datastructures.LinkedList{}
	target := datastructures.LinkedList{}
	for i := 0; i < 10; i++ {
		n := datastructures.Node{}
		n.SetValue(i)
		orig.Add(&n)

		// will delete this node
		// so, don't add this node to result list
		if i == 4 {
			node = n
		} else {
			target.Add(&n)
		}
	}
	// we can expand test cases
	var deleteNodeTests = []deleteNodeTestPair{
		{orig, target, node},
	}
	return deleteNodeTests
}

func TestDeleteNode(t *testing.T) {
	deleteNodeTests := initializeDeleteNode()
	for _, test := range deleteNodeTests {
		// save original list in a temporary list
		temp := datastructures.LinkedList{}
		for elem := range test.orig.GetElements() {
			n := datastructures.Node{}
			n.SetValue(elem)
			temp.Add(&n)
		}

		DeleteNode(&test.node)
		// item by item check through two lists
		for i := 0; i < test.orig.Len(); i++ {
			if test.orig.Find(i) != test.result.Find(i) {
				t.Error("For ", temp.GetElements(),
					"Expected", test.result.GetElements(),
					"got", test.orig.GetElements())
			}
		}
	}
}

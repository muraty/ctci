package chapter2

import ds "ctci/datastructures"

// Given a circular linked list, implement an algorithm which returns
// the node at the beginning of the loop.
// DEFINITION
// Circular linked list: A (corrupt) linked list in which a node's
// next pointer points to an earlier node, so as to make a loop in
// the linked list.
// EXAMPLE
// Input: A->B->C->D->E->C [thesameCasearlier] Output:C

// Finds the first node of the loop
func FindCorruptNode(l ds.LinkedList) *ds.Node {
	head := l.Head()
	slowRunner := head
	fastRunner := head.Next()
	corruptNode := &ds.Node{}

	// Find the Collusion
	for {
		// There is no loop
		if slowRunner.Next() == nil || fastRunner.Next() == nil {
			return nil
		}
		// Move slow runner one point
		slowRunner = slowRunner.Next()
		// Move fast runner two points
		fastRunner = fastRunner.Next().Next()
		// Are they matched ?
		if slowRunner == fastRunner {
			fastRunner = fastRunner.Next()
			break
		}
	}
	// Find the collision point
	slowRunner = head
	for {
		if slowRunner == fastRunner {
			corruptNode = slowRunner
			break
		}
		// Move one point ahead for two of runners
		slowRunner = slowRunner.Next()
		fastRunner = fastRunner.Next()
	}
	return corruptNode
}

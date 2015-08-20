package chapter2

import datastructures "ctci/datastructures"

// Implement an algorithm to delete a node in the middle of a singly linked
// list, given only access to that node. EXAMPLE;
// Input: the node c from the linked list a->b->c->d->e
// Result: nothing isreturned, but the new linked list looks like a->b->d->e

// deletes a node inside the linked list.
// we assume that, we are not given access to the head of the linked list.
// You only have access to that node.
func DeleteNode(n *datastructures.Node) {
	if n.Next() != nil {
		n.SetValue(n.Next().GetValue())
		n.SetNext(n.Next().Next())
	}
}

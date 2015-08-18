package chapter2

import datastructures "ctci/datastructures"

// Implement an algorithm to find the kth to last element of a singly
// linked list.

// Gives sliced linked list from kth element to last element.
// Also it assumes that, we don't know size&last element of the linkedlist.
func Slice(l datastructures.LinkedList, kth int) *datastructures.LinkedList {
	if kth <= 0 {
		return nil
	}
	current := l.Head()
	runner := current
	sliced_list := datastructures.LinkedList{}
	count := 0
	for current != nil {
		runner = current
		if count == kth {
			sliced_list.SetHead(current)
			for runner != nil {
				runner = runner.Next()
			}
			break
			sliced_list.SetTail(current)
		}
		count++
		current = current.Next()
	}
	if sliced_list.Head() != nil {
		return &sliced_list
	} else {
		return nil
	}
}

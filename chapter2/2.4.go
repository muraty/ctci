package chapter2

import datastructures "ctci/datastructures"

// Write code to partition a linked list around a value x, such
// that all nodes less than x come before all nodes greater
// than or equal to x.

// partition given list into two parts due to the given value
// and then merge them together.
func PartitionList(l datastructures.LinkedList, value int) datastructures.LinkedList {
	beforeValue := datastructures.LinkedList{}
	afterValue := datastructures.LinkedList{}

	head := l.Head()
	for head != nil {
		n := datastructures.Node{}
		n.SetValue(head.GetValue())
		if head.GetValue() < value {
			beforeValue.Add(&n)
		} else {
			afterValue.Add(&n)
		}
		head = head.Next()
	}
	beforeValue.SetTail(afterValue.Head())
	return beforeValue
}

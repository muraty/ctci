package chapter2

import datastructures "ctci/datastructures"

// Write code to remove duplicates from an unsorted linked list.
// How would you solve this problem if a temporary buffer is not allowed?

func RemoveDuplicates(l *datastructures.LinkedList) {
	current := l.Head()
	for current != nil {
		runner := current
		for runner.Next() != nil {
			if runner.Next().GetValue() == current.GetValue() {
				runner.SetNext(runner.Next().Next())
			} else {
				runner = runner.Next()
			}
		}
		current = current.Next()
	}
}

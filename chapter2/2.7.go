package chapter2

import ds "ctci/datastructures"

// Implement a function to check if a linked list is a palindrome.

// checks wheter given list is palindrome
func IsPalindrome(l ds.LinkedList) bool {

	head := l.Head()
	// reverse given list
	reversed_list := reverseList(l)
	reversed_head := reversed_list.Head()

	// checks items until half of the list
	for i := 0; i < l.Len()/2; i++ {
		if head.GetValue() != reversed_head.GetValue() {
			return false
		}
		head = head.Next()
		reversed_head = reversed_head.Next()
	}
	return true
}

// reverse given list
func reverseList(l ds.LinkedList) ds.LinkedList {
	temp := l.Copy()
	reversed_list := ds.LinkedList{}
	head := temp.Head()
	current_node := head
	var prev_node *ds.Node
	for {
		if current_node != nil {
			next_node := current_node.Next()
			current_node.SetNext(prev_node)
			prev_node = current_node
			current_node = next_node
		} else {
			reversed_list.SetHead(prev_node)
			break
		}
	}
	return reversed_list
}

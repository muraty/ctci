package ctci

import (
	"fmt"
)

// unit element of the list
type Node struct {
	value int
	next  *Node
}

// get next element of the Node
func (n Node) Next() *Node {
	return n.next
}

// set next node element of the Node
func (n *Node) SetNext(node *Node) {
	n.next = node
}

// set value of the Node
func (n *Node) SetValue(i int) {
	n.value = i
}

// get value of the Node
func (n Node) GetValue() int {
	return n.value
}

// linked list data structure
type LinkedList struct {
	head, tail *Node
}

// set head of the list
func (l *LinkedList) SetHead(n *Node) {
	l.head = n
}

// set tail of the list
func (l *LinkedList) SetTail(n *Node) {
	l.tail.SetNext(n)
}

// adds element to the tail
func (l *LinkedList) Add(n *Node) {
	if l.head != nil {
		l.tail.next = n
		l.tail = n
	} else {
		l.head = n
		l.tail = n
	}
}

// reverses all the elements
func (l *LinkedList) Reverse() {
	head := l.Head()
	current_node := head
	var prev_node *Node
	for {
		if current_node != nil {
			next_node := current_node.Next()
			current_node.next = prev_node
			prev_node = current_node
			current_node = next_node

		} else {
			l.head = prev_node
			break
		}
	}
}

// Find an integer element from the list
func (l *LinkedList) Find(i int) *Node {
	head_node := l.Head()
	for {
		if head_node != nil {
			if i == head_node.value {
				return head_node
			} else {
				head_node = head_node.Next()
			}
		} else {
			return nil
		}

	}
}

func (l *LinkedList) Delete(i int) bool {
	head_node := l.Head()
	tail_node := l.Tail()
	current_node := head_node
	previous_node := current_node
	for {
		if current_node.GetValue() == i {
			if current_node == head_node {
				if l.head.next != nil {
					l.head = l.head.next
					return true
				}
			} else {
				if current_node == tail_node {
					previous_node.next = nil
					return true
				} else {
					previous_node.next = current_node.Next()
					return true
				}
			}
		}
		if current_node.Next() != nil {
			previous_node = current_node
			current_node = current_node.Next()
		} else {
			return false
		}
	}
	return false
}

// get head of the list
func (l *LinkedList) Head() *Node {
	return l.head
}

// get tail of the list
func (l *LinkedList) Tail() *Node {
	return l.tail
}

// get the length of the list
func (l *LinkedList) Len() int {
	count := 0
	current_node := l.Head()

	if l.DetectLoop() {
		return -1
	}
	for {
		if current_node != nil {
			count++
			current_node = current_node.Next()
		} else {
			return count
		}
	}
}

// detects whether linkedlist has an infinite loop
func (l *LinkedList) DetectLoop() bool {
	head := l.Head()
	slowRunner := head
	fastRunner := head.Next()

	// Move ahead two different pointers; slow & fast
	for {
		// There is no loop
		if slowRunner == nil || fastRunner == nil {
			return false
		}
		// There is no loop
		if slowRunner.Next() == nil || fastRunner.Next() == nil {
			return false
		}
		// Move slow runner one point
		slowRunner = slowRunner.Next()
		// Move fast runner two points
		fastRunner = fastRunner.Next().Next()
		// Are they matched ?
		if slowRunner == fastRunner {
			return true
		}
	}
}

// get all the elements from the list
func (l *LinkedList) GetElements() []int {
	head_node := l.Head()
	var node_slice []int
	for {
		node_slice = append(node_slice, head_node.value)
		if head_node.next != nil {
			head_node = head_node.next
		} else {
			return node_slice
		}
	}
}

// get partial elements of the list in order to given size parameter
func (l *LinkedList) GetPartialElements(size int) []int {
	head := l.Head()
	var node_slice []int
	length := 0

	// Handle negative values
	if size < 0 {
		size = 0
	}

	// is there loop inside the list ?
	if l.Len() < 0 {
		// assign a random value as length for looped linkedlist cases
		length = 100
	} else {
		length = l.Len()
	}

	// Does size parameter exceed the length of the list ?
	if size > length {
		size = length
	}

	for i := 0; i < size; i++ {
		node_slice = append(node_slice, head.value)
		if head.next != nil {
			head = head.next
		}
	}
	return node_slice
}

// print all the elements inside the list
func (l *LinkedList) PrintElements() {
	// print all elements
	n := l.GetElements()
	for _, elem := range n {
		fmt.Println(elem)
	}
}

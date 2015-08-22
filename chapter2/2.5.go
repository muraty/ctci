package chapter2

import ds "ctci/datastructures"

// You have two numbers represented by a linked list, where each
// node contains a single digit. The digits are stored in reverse order,
// such that the Ts digit is at the head of the list. Write a function
// that adds the two numbers and returns the sum as a linked list.
// EXAMPLE
// Input:(7-> 1 -> 6) + (5 -> 9 -> 2).Thatis,617 + 295.
// Output: 2 -> 1 -> 9.That is, 912.
// FOLLOW UP
// Suppose the digits are stored in forward order. Repeat the above
// problem. EXAMPLE
// Input:(6 -> 1 -> 7) + (2 -> 9 -> 5).Thatis,617 + 295.
// Output: 9 -> 1 -> 2.That is, 912.

// Sum up two linked list in reverse order
// Input:(7-> 1 -> 6) + (5 -> 9 -> 2).Thatis,617 + 295.
// Output: 2 -> 1 -> 9.That is, 912.
// we assume that, result integer does not exceed integer size
func SumLinkedListReverse(l1 ds.LinkedList, l2 ds.LinkedList) int {
	head1 := l1.Head()
	head2 := l2.Head()
	num1 := 0
	num2 := 0
	multiplier := 1

	// convert l1 list to int
	for head1 != nil {
		num1 += multiplier * head1.GetValue()
		multiplier = multiplier * 10
		head1 = head1.Next()
	}

	multiplier = 1
	// convert l2 list to int
	for head2 != nil {
		num2 += multiplier * head2.GetValue()
		multiplier = multiplier * 10
		head2 = head2.Next()
	}
	return num1 + num2
}

// Sum up two linked list in forward order
// Input:(6 -> 1 -> 7) + (2 -> 9 -> 5).Thatis,617 + 295.
// Output: 9 -> 1 -> 2.That is, 912.
// we assume that, result integer does not exceed integer size
func SumLinkedListForward(l1 ds.LinkedList, l2 ds.LinkedList) int {
	head1 := l1.Head()
	head2 := l2.Head()
	num1 := 0
	num2 := 0
	multiplier1 := exp(10, l1.Len()-1)
	multiplier2 := exp(10, l2.Len()-1)

	// convert l1 list to int
	for head1 != nil {
		num1 += multiplier1 * head1.GetValue()
		multiplier1 = multiplier1 / 10
		head1 = head1.Next()
	}
	// convert l2 list to int
	for head2 != nil {
		num2 += multiplier2 * head2.GetValue()
		multiplier2 = multiplier2 / 10
		head2 = head2.Next()
	}
	return num1 + num2
}

func exp(base int, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

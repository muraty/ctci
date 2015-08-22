package chapter2

import (
	ds "ctci/datastructures"
	"testing"
)

type palindromeTestPair struct {
	orig   ds.LinkedList
	result bool
}

func initializePalindromeTests() []palindromeTestPair {
	l1 := ds.LinkedList{} // palindrome
	l2 := ds.LinkedList{} // non-palindrome

	// create a palindrome list
	// move until half of the list to create a palindrome list
	for i := 0; i < 5; i++ {
		n := ds.Node{}
		n.SetValue(i)
		l1.Add(&n)
	}
	// move from half of the list to create a palindrome list
	for i := 5; i >= 0; i-- {
		n := ds.Node{}
		n.SetValue(i)
		l1.Add(&n)
	}

	// create a non-palindrome list
	for i := 0; i < 10; i++ {
		n := ds.Node{}
		n.SetValue(i)
		l2.Add(&n)
	}
	// we can expand test cases
	var palindromeTests = []palindromeTestPair{
		{l1, true},
		{l2, false},
	}
	return palindromeTests
}

func TestIsPalindrome(t *testing.T) {
	palindromeTests := initializePalindromeTests()
	for _, test := range palindromeTests {
		if elem := IsPalindrome(test.orig); elem != test.result {
			t.Error("For list ", test.orig.GetElements(),
				"Expect", test.result,
				"Got", elem)
		}
	}
}

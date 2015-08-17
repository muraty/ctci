package chapter2

import (
	datastructures "ctci/datastructures"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	l := datastructures.LinkedList{}
	for i := 0; i <= 10; i++ {
		n := &datastructures.Node{}
		n.SetValue(i % 3)
		l.Add(n)
	}

	RemoveDuplicates(&l)

	if l.Len() != 3 {
		t.Error("Wrong length, Len() must be", 3)
	}
	if l.Find(0) == nil || l.Find(1) == nil || l.Find(2) == nil {
		t.Error("Wrong deletion for one of the elements.")
	}
}

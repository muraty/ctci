package ctci

import "testing"

func TestAdd(t *testing.T) {
	l := LinkedList{}
	for i := 0; i < 5; i++ {
		l.Add(&Node{value: i})
	}
	for i := 0; i < 5; i++ {
		if l.Find(i) == nil {
			t.Error("For", i,
				"Expected", true,
				"got", false)
		}
	}
	if l.Len() != 5 {
		t.Error("Wrong length")
	}
}

func TestLength(t *testing.T) {
	l := LinkedList{}
	for i := 0; i < 10; i++ {
		l.Add(&Node{value: i})
	}
	if l.Len() != 10 {
		t.Error("Wrong length")
	}
}

func TestHead(t *testing.T) {
	l := LinkedList{}
	head_elem := Node{value: 11}
	l.Add(&head_elem)
	for i := 0; i < 10; i++ {
		l.Add(&Node{value: i})
	}
	head := l.Head()

	if *head != head_elem {
		t.Error("For head",
			"Expected", head_elem,
			"got", head)
	}
}

func TestTail(t *testing.T) {
	l := LinkedList{}
	for i := 0; i < 10; i++ {
		l.Add(&Node{value: i})
	}
	tail_elem := Node{value: 11}
	l.Add(&tail_elem)
	tail := l.Tail()

	if *tail != tail_elem {
		t.Error("For tail",
			"Expected", tail_elem,
			"got", tail)
	}
}

func TestFind(t *testing.T) {
	l := LinkedList{}
	for i := 10; i < 20; i++ {
		l.Add(&Node{value: i})
	}
	elem3 := Node{value: 3}
	l.Add(&elem3)

	if l.Find(3) == nil {
		t.Error("For", 3,
			"Expected", elem3,
			"got", nil)
	}
	if elem := l.Find(4); elem != nil {
		t.Error("For", 3,
			"Expected", nil,
			"got", elem)
	}
}

func TestDelete(t *testing.T) {
	l := LinkedList{}
	for i := 0; i < 10; i++ {
		l.Add(&Node{value: i})
	}
	l.Delete(5)
	l.Delete(4)
	l.Delete(3)

	if elem := l.Find(5); elem != nil {
		t.Error("For", 5,
			"Expected", nil,
			"got", elem)
	}
	if elem := l.Find(4); elem != nil {
		t.Error("For", 4,
			"Expected", nil,
			"got", elem)
	}
	if elem := l.Find(3); elem != nil {
		t.Error("For", 3,
			"Expected", nil,
			"got", elem)
	}

}

func TestReverse(t *testing.T) {
	l := LinkedList{}
	for i := 0; i < 10; i++ {
		l.Add(&Node{value: i})
	}
	before_reverse := l.GetElements()
	l.Reverse()
	after_reverse := l.GetElements()

	for i := 0; i < 10; i++ {
		if before_reverse[i] != after_reverse[len(after_reverse)-i-1] {
			t.Error("For", i,
				"Expected", before_reverse[i],
				"got", after_reverse[i])
		}
	}
	if len(before_reverse) != len(after_reverse) {
		t.Error("For", 10,
			"Expected", len(before_reverse),
			"got", len(after_reverse))
	}

}

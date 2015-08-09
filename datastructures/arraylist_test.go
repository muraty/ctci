package ctci

import (
	"testing"
)

func TestLen(t *testing.T) {
	l := ArrayList{}
	l.Init(10)

	// test initial size
	if l.Len() != 0 {
		t.Error("For length",
			"Expected", 10,
			"got", l.Len())
	}
	for i := 0; i < 15; i++ {
		l.Set(i)
	}
	// test new size
	if l.Len() != 15 {
		t.Error("For length",
			"Expected", 20,
			"got", l.Len())
	}

	for i := 0; i < 15; i++ {
		l.Set(i)
	}
	// test new size
	if l.Len() != 30 {
		t.Error("For length",
			"Expected", 40,
			"got", l.Len())
	}

}

func TestSet(t *testing.T) {
	l := ArrayList{}
	l.Init(10)
	// set to idx=0
	l.Set(10)
	// set to idx=1
	l.Set(13)
	// set to idx=2
	l.Set(15)

	if l.Get(0) != 10 {
		t.Error("For length",
			"Expected", 10,
			"got", l.Get(0))
	}
	if l.Get(1) != 13 {
		t.Error("For length",
			"Expected", 13,
			"got", l.Get(1))
	}
	if l.Get(2) != 15 {
		t.Error("For length",
			"Expected", 15,
			"got", l.Get(2))
	}

}
func TestGet(t *testing.T) {
	l := ArrayList{}
	l.Init(10)

	// set to idx=0
	l.Set(10)

	if l.Get(0) != 10 {
		t.Error("For length",
			"Expected", 10,
			"got", l.Get(0))
	}

}
func TestGetElements(t *testing.T) {

	l := ArrayList{}
	l.Init(10)
	// set to idx=0
	l.Set(10)
	// set to idx=1
	l.Set(13)

	if elem := l.GetElements()[0]; elem != 10 {
		t.Error("For length",
			"Expected", 10,
			"got", elem)
	}
	if elem := l.GetElements()[1]; elem != 13 {
		t.Error("For length",
			"Expected", 13,
			"got", elem)
	}
	if elem := l.GetElements()[2]; elem != 0 {
		t.Error("For length",
			"Expected", 0,
			"got", elem)
	}

}

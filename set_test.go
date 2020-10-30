package set_test

import (
	"testing"

	"github.com/zeurd/set"
)

func TestAddString(t *testing.T) {
	s := set.New()
	e := "in set"
	added := s.Add(e)
	if !added {
		t.Errorf("could not add element in set: %v", e)
	}
	ok := s.Contains(e)
	if !ok {
		t.Errorf("could not find element in set: %v", e)
	}
}

func TestCannotAddTwice(t *testing.T) {
	s := set.New()
	e := "in set"
	added := s.Add(e)
	if !added {
		t.Errorf("could not element in set: %v", e)
	}
	ok := s.Contains(e)
	if !ok {
		t.Errorf("could not find element in set: %v", e)
	}
	added = s.Add(e)
	if added {
		t.Errorf("same element was added twice: %v", e)
	}
}

func TestDifferentType(t *testing.T) {
	s := set.New()
	e := "inset"
	e2 := 10
	added := s.Add(e)
	if !added {
		t.Errorf("could not element in set: %v", e)
	}
	ok := s.Contains(e)
	if !ok {
		t.Errorf("could not find element in set: %v", e)
	}
	added = s.Add(e2)
	if !added {
		t.Errorf("other element was not added twice: %v", e)
	}
	ok = s.Contains(e2)
	if !ok {
		t.Errorf("could not find element in set: %v", e)
	}
}

func TestDoesNotContain(t *testing.T) {
	s := set.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(5)
	ok := s.Contains(4)
	if ok {
		t.Errorf("found absent element")
	}
	length := len(s)
	if length != 4 {
		t.Errorf("set does not contain all added elements")
	}
}

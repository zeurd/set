package set

import (
	"fmt"
)

//Set is a set
type Set map[interface{}]interface{}

//New creates a string set
func New() Set {
	return make(Set)
}

//Add adds an element to the set and returns true if at least one element has been added
func (s Set) Add(e ...interface{}) bool {
	done := false
	for _, f := range e {
		if _, ok := s[f]; ok {
			continue
		}
		s[f] = nil
		_, ok := s[f]
		done = ok
	}
	return done
}

//Contains returns true if the set contains e
func (s Set) Contains(e interface{}) bool {
	_, ok := s[e]
	return ok
}

func (s Set) String() string {
	str := "{"
	i := 0
	for k := range s {
		str += fmt.Sprintf("%v", k)
		if i != len(s)-1 {
			str += ", "
		}
		i++
	}
	str += "}"
	return str
}

package set

//Set is a set
type Set map[interface{}]interface{}

//New creates a string set
func New() Set {
	return make(Set)
}

//Add adds an element to the set and returns true if the operation was successful
func (s Set) Add(e interface{}) bool {
	if _, ok := s[e]; ok {
		return false
	}
	s[e] = nil
	_, ok := s[e]
	return ok
}

//Contains returns true if the set contains e
func (s Set) Contains(e interface{}) bool {
	_, ok := s[e]
	return ok
}

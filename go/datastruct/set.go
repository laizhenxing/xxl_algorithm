package datastruct

type Set struct {
	m map[interface{}]struct{}
}

func NewSet(items ...interface{}) *Set {
	s := &Set{
		m: make(map[interface{}]struct{}),
	}
	s.Store(items...)
	return s
}

func (s *Set) Store(items ...interface{}) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *Set) Remove(v interface{}) {
	if s.Contains(v) {
		delete(s.m, v)
	}
}

func (s *Set) Contains(v interface{}) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Clear() (members []interface{}) {
	members = s.Members()
	s.m = make(map[interface{}]struct{})
	return
}

func (s *Set) Equal(s1 *Set) bool {
	for key := range s1.m {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) IsSubSet(s1 *Set) bool {
	if s1.Len() > s.Len() {
		return false
	}

	for key := range s1.m {
		if !s.Contains(key) {
			return false
		}
	}

	return true
}

func (s *Set) Restore(s1 *Set) {
	for v := range s1.m {
		s.Store(v)
	}
}

// func (s *Set) LoadAndDelete() interface{} {

// }

func (s *Set) Members() (members []interface{}) {
	for key := range s.m {
		members = append(members, key)
	}
	return
}

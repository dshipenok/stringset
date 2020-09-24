package sset

type StringSet struct {
	m map[string]struct{}
}

// NewStringSet constructor
func NewStringSet(values ...string) *StringSet {
	m := make(map[string]struct{}, len(values))
	for _, val := range values {
		m[val] = struct{}{}
	}
	return &StringSet{
		m: m,
	}
}

func (s *StringSet) Add(values ...string) {
	for _, value := range values {
		s.m[value] = struct{}{}
	}
}

func (s *StringSet) Remove(values ...string) {
	for _, value := range values {
		delete(s.m, value)
	}
}

func (s *StringSet) Merge(toMerge *StringSet) {
	for val := range toMerge.m {
		s.m[val] = struct{}{}
	}
}

func (s *StringSet) Has(values ...string) bool {
	for _, val := range values {
		if _, has := s.m[val]; !has {
			return false
		}
	}
	return true
}

func (s *StringSet) Count() int {
	return len(s.m)
}

func (s *StringSet) Empty() bool {
	return len(s.m) == 0
}

func (s *StringSet) Slice() []string {
	slice := make([]string, 0, len(s.m))
	for val := range s.m {
		slice = append(slice, val)
	}
	return slice
}

func (s *StringSet) Map() map[string]struct{} {
	return s.m
}

package set

import "sort"

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

// NewStringSuperSet groups content of several StringSet's
func NewStringSuperSet(sets ...*StringSet) *StringSet {
	set := NewStringSet()
	for _, s := range sets {
		set.Merge(s)
	}
	return set
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

func (s *StringSet) Has(value string) bool {
	_, has := s.m[value]
	return has
}

func (s *StringSet) HasAnyFrom(o *StringSet) bool {
	for value := range o.m {
		_, has := s.m[value]
		if has {
			return true
		}
	}
	return false
}

func (s *StringSet) Equals(set *StringSet) bool {
	if len(s.m) != len(set.m) {
		return false // fast-path
	}
	// check each key:
	for key := range s.m {
		if !set.Has(key) {
			return false
		}
	}
	return true
}

func (s *StringSet) Intersection(set *StringSet) *StringSet {
	a, b := s, set
	if a.Count() == 0 || b.Count() == 0 {
		return NewStringSet()
	}
	if a.Count() > b.Count() {
		b, a = a, b
	}

	result := NewStringSet()
	// check each key:
	for key := range a.m {
		if b.Has(key) {
			result.Add(key)
		}
	}
	return result
}

func (s *StringSet) Subtract(sub *StringSet) *StringSet {
	if s.Count() == 0 {
		return NewStringSet()
	}
	if sub.Count() == 0 {
		return s.Clone()
	}

	result := NewStringSet()
	// check each key:
	for key := range s.m {
		if sub.Has(key) {
			continue
		}
		result.Add(key)
	}
	return result
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

func (s *StringSet) SortedSlice() []string {
	slice := s.Slice()
	sort.Strings(slice)
	return slice
}

func (s *StringSet) Set() map[string]struct{} {
	return s.m
}

func (s *StringSet) Clone() *StringSet {
	m := make(map[string]struct{}, len(s.m))
	for val := range s.m {
		m[val] = struct{}{}
	}
	return &StringSet{m: m}
}

func (s *StringSet) ForEach(f func(string)) {
	for val := range s.m {
		f(val)
	}
}

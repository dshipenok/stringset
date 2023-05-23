package set

import (
	"sort"
)

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Set[T ordered] struct {
	m map[T]struct{}
}

// NewSet constructor
func NewSet[T ordered](values ...T) *Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, val := range values {
		m[val] = struct{}{}
	}
	return &Set[T]{
		m: m,
	}
}

// NewStringSet constructor
func NewStringSet(values ...string) *Set[string] {
	return NewSet(values...)
}

// NewStringSuperSet groups content of several StringSet's
func NewStringSuperSet(sets ...*Set[string]) *Set[string] {
	set := NewStringSet()
	for _, s := range sets {
		set.Merge(s)
	}
	return set
}

func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		s.m[value] = struct{}{}
	}
}

func (s *Set[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s.m, value)
	}
}

func (s *Set[T]) Merge(toMerge *Set[T]) {
	for val := range toMerge.Set() {
		s.m[val] = struct{}{}
	}
}

func (s *Set[T]) Has(value T) bool {
	if s == nil {
		return false
	}
	_, has := s.m[value]
	return has
}

func (s *Set[T]) HasAnyFrom(o *Set[T]) bool {
	if s == nil {
		return false
	}
	for value := range o.m {
		_, has := s.m[value]
		if has {
			return true
		}
	}
	return false
}

func (s *Set[T]) Equals(set *Set[T]) bool {
	if s == nil || set == nil {
		return s == set
	}
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

func (s *Set[T]) Intersection(set *Set[T]) *Set[T] {
	a, b := s, set
	if a.Count() == 0 || b.Count() == 0 {
		return NewSet[T]()
	}
	if a.Count() > b.Count() {
		b, a = a, b
	}

	result := NewSet[T]()
	// check each key:
	for key := range a.m {
		if b.Has(key) {
			result.Add(key)
		}
	}
	return result
}

func (s *Set[T]) Subtract(sub *Set[T]) *Set[T] {
	if s.Count() == 0 {
		return NewSet[T]()
	}
	if sub.Count() == 0 {
		return s.Clone()
	}

	result := NewSet[T]()
	// check each key:
	for key := range s.m {
		if sub.Has(key) {
			continue
		}
		result.Add(key)
	}
	return result
}

func (s *Set[T]) Count() int {
	if s == nil {
		return 0
	}
	return len(s.m)
}

func (s *Set[T]) Empty() bool {
	return s == nil || len(s.m) == 0
}

func (s *Set[T]) Slice() []T {
	if s == nil {
		return nil
	}
	slice := make([]T, 0, len(s.m))
	for val := range s.m {
		slice = append(slice, val)
	}
	return slice
}

func (s *Set[T]) SortedSlice() []T {
	slice := s.Slice()
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func (s *Set[T]) Set() map[T]struct{} {
	if s == nil {
		return nil
	}
	return s.m
}

func (s *Set[T]) Clone() *Set[T] {
	if s == nil {
		return nil
	}
	m := make(map[T]struct{}, len(s.m))
	for val := range s.m {
		m[val] = struct{}{}
	}
	return &Set[T]{m: m}
}

func (s *Set[T]) ForEach(f func(T)) {
	if s == nil {
		return
	}
	for val := range s.m {
		f(val)
	}
}

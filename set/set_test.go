package set

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringSet_Has(t *testing.T) {
	set := NewStringSet("a", "123", "abc")

	tests := []struct {
		desc      string
		value     string
		expectHas bool
	}{
		{desc: "Empty value",
			value:     "",
			expectHas: false},
		{desc: "Existing value",
			value:     "123",
			expectHas: true},
		{desc: "Non-exisiting value",
			value:     "def",
			expectHas: false},
	}

	for _, tst := range tests {
		t.Run(tst.desc, func(t *testing.T) { // nolint

			has := set.Has(tst.value)

			assert.Equal(t, tst.expectHas, has)
		})
	}
}

func Test_StringSet_Merge(t *testing.T) {
	set1 := NewStringSet("a", "123", "abc")
	set2 := NewStringSet("def", "999", "123")

	set1.Merge(set2)

	assert.Equal(t, 5, set1.Count())
	assert.True(t, set1.Has("123"))
	assert.True(t, set1.Has("999"))
}

func Test_StringSet_Add(t *testing.T) {
	set := NewStringSet("a", "123", "abc")

	set.Add("999")

	assert.Equal(t, 4, set.Count())
	assert.True(t, set.Has("999"))
}

func Test_StringSet_Subtract(t *testing.T) {
	tests := []struct {
		a        *StringSet
		b        *StringSet
		expected []string
	}{
		{
			a:        NewStringSet(),
			b:        NewStringSet(),
			expected: []string{},
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet(),
			expected: []string{"a"},
		},
		{
			a:        NewStringSet(),
			b:        NewStringSet("b"),
			expected: []string{},
		},
		{
			a:        NewStringSet("a", "123", "d", "abc"),
			b:        NewStringSet("b", "123", "c", "d"),
			expected: []string{"a", "abc"},
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			result := tst.a.Subtract(tst.b)

			assert.Equal(t, tst.expected, result.SortedSlice())
		})
	}
}

func Test_StringSet_Intersection(t *testing.T) {
	tests := []struct {
		a        *StringSet
		b        *StringSet
		expected []string
	}{
		{
			a:        NewStringSet(),
			b:        NewStringSet(),
			expected: []string{},
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet(),
			expected: []string{},
		},
		{
			a:        NewStringSet(),
			b:        NewStringSet("b"),
			expected: []string{},
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet("b"),
			expected: []string{},
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet("a"),
			expected: []string{"a"},
		},
		{
			a:        NewStringSet("a", "123", "d", "abc", "def"),
			b:        NewStringSet("b", "123", "c", "d"),
			expected: []string{"123", "d"},
		},
		{
			a:        NewStringSet("b", "123", "c", "d"),
			b:        NewStringSet("a", "123", "d", "abc", "def"),
			expected: []string{"123", "d"},
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			result := tst.a.Intersection(tst.b)

			assert.Equal(t, tst.expected, result.SortedSlice())
		})
	}
}

func Test_StringSet_Remove(t *testing.T) {
	tests := []struct {
		a        *StringSet
		value    string
		expected []string
	}{
		{
			a:        NewStringSet(),
			value:    "",
			expected: []string{},
		},
		{
			a:        NewStringSet("a"),
			value:    "b",
			expected: []string{"a"},
		},
		{
			a:        NewStringSet("a"),
			value:    "a",
			expected: []string{},
		},
		{
			a:        NewStringSet("a", "b", "c"),
			value:    "b",
			expected: []string{"a", "c"},
		},

		{
			a:        NewStringSet("a", "b", "c"),
			value:    "1",
			expected: []string{"a", "b", "c"},
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			tst.a.Remove(tst.value)

			assert.Equal(t, tst.expected, tst.a.SortedSlice())
		})
	}
}

func Test_StringSet_Empty(t *testing.T) {
	tests := []struct {
		a        *StringSet
		expected bool
	}{
		{
			a:        NewStringSet(),
			expected: true,
		},
		{
			a:        NewStringSet("a"),
			expected: false,
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			empty := tst.a.Empty()

			assert.Equal(t, tst.expected, empty)
		})
	}
}

func Test_StringSet_Equals(t *testing.T) {
	tests := []struct {
		a        *StringSet
		b        *StringSet
		expected bool
	}{
		{
			a:        NewStringSet(),
			b:        NewStringSet(),
			expected: true,
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet("a"),
			expected: true,
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet("b"),
			expected: false,
		},
		{
			a:        NewStringSet(),
			b:        NewStringSet("b"),
			expected: false,
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet(),
			expected: false,
		},
		{
			a:        NewStringSet("b", "a"),
			b:        NewStringSet("a", "b"),
			expected: true,
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			equals := tst.a.Equals(tst.b)

			assert.Equal(t, tst.expected, equals)
		})
	}
}

func Test_StringSet_HasAnyFrom(t *testing.T) {
	tests := []struct {
		a        *StringSet
		b        *StringSet
		expected bool
	}{
		{
			a:        NewStringSet(),
			b:        NewStringSet(),
			expected: false,
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet(),
			expected: false,
		},
		{
			a:        NewStringSet(),
			b:        NewStringSet("a"),
			expected: false,
		},
		{
			a:        NewStringSet("a"),
			b:        NewStringSet("a"),
			expected: true,
		},
		{
			a:        NewStringSet("a", "b"),
			b:        NewStringSet("b"),
			expected: true,
		},
		{
			a:        NewStringSet("a", "b", "c"),
			b:        NewStringSet("b", "c"),
			expected: true,
		},
	}
	for i, tst := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			has := tst.a.HasAnyFrom(tst.b)

			assert.Equal(t, tst.expected, has)
		})
	}
}

func Test_StringSet_Map(t *testing.T) {
	set := NewStringSet("a", "123", "abc")

	m := set.Set()

	assert.Equal(t, len(m), set.Count())
	assert.Contains(t, m, "a")
	assert.Contains(t, m, "123")
	assert.Contains(t, m, "abc")
}

func Test_StringSet_ForEach(t *testing.T) {
	set := NewStringSet("a", "123", "abc")
	result := []string{}

	set.ForEach(func(value string) {
		result = append(result, value)
	})

	sort.Strings(result)
	assert.Equal(t, set.SortedSlice(), result)
}

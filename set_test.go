package set

import (
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

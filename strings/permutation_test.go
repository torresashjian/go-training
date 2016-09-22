package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPermutationFalse(t *testing.T) {
	isPerm := IsPermutation("", "")
	assert.False(t, isPerm)
	isPerm = IsPermutation("a", "b")
	assert.False(t, isPerm)
	isPerm = IsPermutation("aa", "a")
	assert.False(t, isPerm)
	isPerm = IsPermutation("aa", "aaa")
	assert.False(t, isPerm)
	isPerm = IsPermutation("ab", "aa")
	assert.False(t, isPerm)
	isPerm = IsPermutation("abc", "acc")
	assert.False(t, isPerm)
}

func TestIsPermutationTrue(t *testing.T) {
	isPerm := IsPermutation("a", "a")
	assert.True(t, isPerm)
	isPerm = IsPermutation("ab", "ab")
	assert.True(t, isPerm)
	isPerm = IsPermutation("ab", "ba")
	assert.True(t, isPerm)
	isPerm = IsPermutation("abcd", "badc")
	assert.True(t, isPerm)
}

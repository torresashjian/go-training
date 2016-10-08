package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompression(t *testing.T) {
	c := Compression("")
	assert.Equal(t, "", c)
	c = Compression("aa")
	assert.Equal(t, "aa", c)
	c = Compression("aaa")
	assert.Equal(t, "a3", c)
	c = Compression("aa aaa")
	assert.Equal(t, "aa aaa", c)
	c = Compression("abccccccc")
	assert.Equal(t, "a1b1c7", c)
	c = Compression("aaaffcccc ")
	assert.Equal(t, "a3f2c4 1", c)

}

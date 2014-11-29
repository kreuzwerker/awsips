package awsips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	assert := assert.New(t)

	s := newSortedSet()

	assert.True(s.IsEmpty())

	s.Add("foo")
	s.Add("foo")
	s.Add("bar")
	s.Add("wee")

	assert.False(s.IsEmpty())

	assert.Equal([]string{"bar", "foo", "wee"}, s.Sorted())

	assert.True(s.IsIncluded("foo"))
	assert.True(s.IsIncluded("bar"))
	assert.False(s.IsIncluded("gee"))

}

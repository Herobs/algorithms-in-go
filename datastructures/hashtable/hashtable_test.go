package hashtable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable(t *testing.T) {
	table := New()

	table.Set("hello", "world")
	assert.True(t, table.Has("hello"))
	assert.False(t, table.Has("foo"))
	assert.Equal(t, "world", table.Get("hello"))

	table.Set("foo", "bar")
	assert.True(t, table.Has("foo"))
	assert.Equal(t, "bar", table.Get("foo"))

	table.Set("oof", "rab")
	assert.True(t, table.Has("oof"))
	assert.Equal(t, "rab", table.Get("oof"))
}

package deque

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeque(t *testing.T) {
	q := New[string]()
	q.PushBack("bar")
	q.PushBack("baz")
	q.PushFront("foo")

	require.Equal(t, 3, q.Len())

	var v string
	v, ok := q.Front()
	require.True(t, ok)
	require.Equal(t, "foo", v)
	require.Equal(t, 3, q.Len())

	v, ok = q.Back()
	require.True(t, ok)
	require.Equal(t, "baz", v)
	require.Equal(t, 3, q.Len())

	v, ok = q.PopFront()
	require.True(t, ok)
	require.Equal(t, "foo", v)
	require.Equal(t, 2, q.Len())

	v, ok = q.PopBack()
	require.True(t, ok)
	require.Equal(t, "baz", v)
	require.Equal(t, 1, q.Len())

	_, ok = q.PopBack()
	require.True(t, ok)
	require.Equal(t, 0, q.Len())

	_, ok = q.PopBack()
	require.False(t, ok)
	require.Equal(t, 0, q.Len())
}

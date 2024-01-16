package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_Int(t *testing.T) {
	s := New[int]()
	s.Push(1)
	require.Equal(t, 1, s.Len())

	v, ok := s.Pop()
	require.True(t, ok)
	require.Equal(t, 1, v)
	require.Equal(t, 0, s.Len())

	v, ok = s.Pop()
	require.False(t, ok)
	require.Equal(t, 0, v)
	require.Equal(t, 0, s.Len())
}

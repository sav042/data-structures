package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Push(value T) bool {
	s.items = append(s.items, value)
	return true
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	l := len(s.items)
	if l == 0 {
		return
	}
	s.items, value = s.items[:l-1], s.items[l-1]
	return value, true
}

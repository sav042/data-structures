package deque

import "github.com/gammazero/deque"

type Deque[T any] struct {
	items *deque.Deque[T]
}

func New[T any](size ...int) *Deque[T] {
	return &Deque[T]{
		items: deque.New[T](size...),
	}
}

func (q *Deque[T]) Len() int {
	return q.items.Len()
}

func (q *Deque[T]) PushBack(value T) {
	q.items.PushBack(value)
}

func (q *Deque[T]) PushFront(value T) {
	q.items.PushFront(value)
}

func (q *Deque[T]) PopBack() (value T, ok bool) {
	l := q.items.Len()
	if l == 0 {
		return
	}
	value = q.items.PopBack()
	return value, true
}

func (q *Deque[T]) PopFront() (value T, ok bool) {
	l := q.items.Len()
	if l == 0 {
		return
	}
	value = q.items.PopFront()
	return value, true
}

func (q *Deque[T]) Back() (value T, ok bool) {
	l := q.items.Len()
	if l == 0 {
		return
	}
	value = q.items.Back()
	return value, true
}

func (q *Deque[T]) Front() (value T, ok bool) {
	l := q.items.Len()
	if l == 0 {
		return
	}
	value = q.items.Front()
	return value, true
}

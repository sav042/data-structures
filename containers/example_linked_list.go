package containers

import (
	"container/list"
)

func ExampleList() {
	l := list.New()
	l.PushBack(2)
	l.PushFront("1")
}

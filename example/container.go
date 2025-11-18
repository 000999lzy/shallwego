package main

import (
	"container/list"
	"container/ring"
)

func containerFunc() {
	ringFunc()
}

func listFunc() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		println(e.Value.(int))
	}
}

func ringFunc() {
	r := ring.New(2)
	s := ring.New(2)

	for i := 1; i < 3; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 3; i < 5; i++ {
		s.Value = i
		s = s.Next()
	}

	rs := r.Link(s)

	rs.Do(func(p any) {
		println(p.(int))
	})
	n := ring.New(5)
	for i := 2; i < 7; i++ {
		n.Value = i
		n = n.Next()
	}

	/*nrs := r.Link(n)
	nrs.Do(func(p any) {
		println(p.(int))
	})*/
}

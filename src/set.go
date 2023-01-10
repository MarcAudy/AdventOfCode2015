package main

type set[T comparable] map[T]struct{}

func NewSet[T comparable](vars ...T) set[T] {
	s := set[T](make(map[T]struct{}))
	for _, v := range vars {
		s.Add(v)
	}
	return s
}

func (s *set[T]) Add(value T) bool {
	prevLen := len(*s)
	(*s)[value] = struct{}{}
	return prevLen != len(*s)
}

func (s *set[T]) Remove(value T) {
	delete(*s, value)
}

func (s *set[T]) Contains(value T) bool {
	_, c := (*s)[value]
	return c
}

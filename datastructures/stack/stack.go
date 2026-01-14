package stack

import "errors"

var ErrEmptyStack = errors.New("empty stack")

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var empty T
		return empty, ErrEmptyStack
	}

	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var empty T
		return empty, ErrEmptyStack
	}

	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex], nil
}

// IsEmpty: 스택이 비었는지 확인합니다.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

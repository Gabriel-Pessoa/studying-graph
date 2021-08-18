package stack

import "errors"

type Stack interface {
	Push(element int)
	Pop() (int, error)
	Clear()
	IsEmpty() bool
	Size() int
}

type stack struct {
	Items map[int]int
	Count int
}

func NewStack() Stack {
	return &stack{
		Items: map[int]int{},
		Count: 0,
	}
}

func (s *stack) Push(element int) {
	s.Items[s.Count] = element
	s.Count++
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("fail to remove element in the stack")
	}

	if element, ok := s.Items[s.Count-1]; ok {
		s.Count--
		delete(s.Items, s.Count)
		return element, nil
	}

	return -1, errors.New("fail to remove element in the stack")
}

func (s *stack) Clear() {
	s.Items = map[int]int{}
	s.Count = 0
}

func (s stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s stack) Size() int {
	return s.Count
}

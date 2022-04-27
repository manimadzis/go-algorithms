package stack

import "errors"

type ListStackItem struct {
	value interface{}
	next  *ListStackItem
}

type ListStack struct {
	head *ListStackItem
	tail *ListStackItem
}

func NewListStack() *ListStack {
	return &ListStack{
		head: nil,
		tail: nil,
	}
}

func (s *ListStack) Push(value interface{}) error {
	node := new(ListStackItem)
	node.value = value

	if s.head == nil {
		s.head = node
	} else {
		s.tail.next = node
	}
	s.tail = node

	return nil
}

func (s *ListStack) Pop() (interface{}, error) {
	var value interface{}

	if s.tail == nil {
		return nil, errors.New("Empty stack")
	} else if s.head.next == nil {
		value = s.head.value
		s.head = nil
		s.tail = nil
	} else {
		value = s.tail.value
		item := s.head
		new_tail := s.head

		for item != s.tail {
			new_tail = item
			item = item.next
		}

		new_tail.next = nil
		s.tail = new_tail
	}

	return value, nil
}

func (s *ListStack) Peek() (interface{}, error) {
	if s.tail == nil {
		return nil, errors.New("Empty stack")
	} else {
		return s.tail.value, nil
	}
}

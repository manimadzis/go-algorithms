package queue

import "errors"

type QueueItem struct {
	value interface{}
	next  *QueueItem
}

type Queue struct {
	head *QueueItem
	tail *QueueItem
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
	}
}

func (q *Queue) Push(value interface{}) {
	new_tail := &QueueItem{
		value: value,
		next:  nil,
	}

	if q.head == nil {
		q.tail = new_tail
		q.head = new_tail
	} else {
		q.tail.next = new_tail
		q.tail = new_tail
	}
}

func (q *Queue) Pop() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("Empty queue")
	}

	var res interface{}
	if q.head.next == nil {
		res = q.head.value
		q.head = nil
		q.tail = nil
	} else {
		res = q.head.value
		q.head = q.head.next
	}

	return res, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("Empty queue")
	}

	return q.head.value, nil
}

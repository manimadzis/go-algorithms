package tests

import (
	"fmt"
	"testing"

	. "github.com/manimadzis/go-algorithms/queue"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()

	value, err := queue.Pop()

	fmt.Println(err)

	queue.Push(100)
	queue.Push(200)
	queue.Push(300)

	for value, err = queue.Pop(); err == nil; value, err = queue.Pop() {
		fmt.Println(value)
	}

	queue.Push(100)
	queue.Push(200)
	queue.Push(300)

	value, err = queue.Peek()
	fmt.Println(value)
	value, err = queue.Peek()
	fmt.Println(value)
}

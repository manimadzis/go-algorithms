package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/manimadzis/go-algorithms/stack"
)

func TestList(t *testing.T) {

	stack := stack.NewListStack()

	for i := 5; i < 10; i++ {
		stack.Push(i)
	}

	var item interface{}
	var err error

	for item, err = stack.Pop(); err == nil; item, err = stack.Pop() {
		fmt.Println("Item: ", item)
		time.Sleep(200 * time.Millisecond)
	}
}

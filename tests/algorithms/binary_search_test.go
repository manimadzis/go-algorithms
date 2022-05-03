package tests

import (
	"fmt"
	"testing"

	"github.com/manimadzis/go-algorithms/algorithms"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 5, 6, 9, 10}

	fmt.Println(algorithms.BinarySearch(a, 5))
	fmt.Println(algorithms.BinarySearch(a, 0))
	fmt.Println(algorithms.BinarySearch(a, 14))
	fmt.Println(algorithms.BinarySearch(a, 8))
	fmt.Println(algorithms.BinarySearch(a, 9))
}

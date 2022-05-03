package tests

import (
	"fmt"
	"sort"
	"testing"

	"github.com/manimadzis/go-algorithms/sorts"
)

type MyInts struct {
	array []int
}

func (m *MyInts) Len() int {
	return len(m.array)
}

func (m *MyInts) Less(i int, j int) bool {
	return m.array[i] < m.array[j]
}

func (m *MyInts) Swap(i int, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

func TestBubbleSort(t *testing.T) {
	a := MyInts{
		array: []int{5, 4, 3, 2, 1},
	}

	b := MyInts{
		array: []int{5, 4, 3, 2, 1},
	}

	sorts.BubbleSort(&a)

	fmt.Println(a)

	sort.Sort(&b)
	fmt.Println(b)
}

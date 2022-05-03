package tests

import (
	"fmt"
	"testing"

	"github.com/manimadzis/go-algorithms/sorts"
)

func TestCountSort(t *testing.T) {

	a := []int{8, 1, 6, 9, 1, 7, 765}

	sorts.CountSort(a)

	fmt.Print(a)

}

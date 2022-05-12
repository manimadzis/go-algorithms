package BST_test

import (
	"fmt"
	"testing"

	"github.com/manimadzis/go-algorithms/BST"
)

func TestBinaryTree(t *testing.T) {
	bintree := BST.New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	bintree.Add(15)
	bintree.Add(14)
	bintree.Add(13)

	bintree.Print()
	fmt.Println("Expected: 15 -> 14 -> 13")
	fmt.Println()

	bintree.Del(14)
	bintree.Print()
	fmt.Println("Expected: 15 -> 13")
	fmt.Println()

	bintree.Del(13)
	bintree.Print()
	fmt.Println("Expected: Nothing")
	fmt.Println()

	bintree.Add(20)
	bintree.Add(16)
	bintree.Add(14)

	bintree.Print()
	fmt.Println("                  /-> 16")
	fmt.Println("Expected: 15 -> 20")
	fmt.Println("            \\-> 14")
}

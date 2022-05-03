package main

import "fmt"

func main() {
	a := []int{4, 1, 5}

	fmt.Print(a)
	zeros(a)
	fmt.Print(a)

}

func zeros(data []int) {
	data[0] = 0
	data[1] = 0
	data[2] = 0
}

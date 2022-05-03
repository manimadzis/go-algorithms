package algorithms

func BinarySearch(data []int, value int) int {
	left, right := -1, len(data)

	for right-left > 1 {
		mid := (right + left) / 2
		if data[mid] >= value {
			right = mid
		} else {
			left = mid
		}
	}

	index := right

	if index < 0 || index >= len(data) || value != data[index] {
		index = -1
	}

	return index
}

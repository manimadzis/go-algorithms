package sorts

func BubbleSort(data Sortable) {
	len := data.Len()
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}
}

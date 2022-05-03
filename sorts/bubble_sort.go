package sorts

func BubbleSort(data Sortable) {
	len := data.Len()
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}
}

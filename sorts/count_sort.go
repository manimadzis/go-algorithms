package sorts

func CountSort(data []int) {
	min, max := data[0], data[0]
	n := len(data)

	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}

		if data[i] > max {
			max = data[i]
		}
	}

	counts := make([]int, max-min+1)

	for i := 0; i < n; i++ {
		counts[data[i]-min]++
	}

	index := 0
	for i := 0; i < len(counts); i++ {
		for ; counts[i] > 0; counts[i]-- {
			data[index] = i + min
			index++
		}
	}
}

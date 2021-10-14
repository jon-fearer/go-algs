package algorithms

func QuickSort(data []int) []int {
	return quickSortSubset(data, 0, len(data) - 1)
}

func quickSortSubset(data []int, left int, right int) []int {
	j := partition(data, left, right)
	if left < j - 1 {
		quickSortSubset(data, left, j - 1)
	}
	if j + 1 < right {
		quickSortSubset(data, j + 1, right)
	}
	return data
}

func partition(data []int, left int, right int) int {
	pivot := data[left]
	i := left
	j := right
	for i < j {
		for i < len(data) - 1 && data[i] <= pivot {
			i++
		}
		for j > 0 && data[j] > pivot {
			j--
		}
		if i < j {
			swap(data, i, j)
		}
	}
	swap(data, left, j)
	return j
}

func swap(data []int, i int, j int) {
	iVal := data[i]
	jVal := data[j]
	data[i] = jVal
	data[j] = iVal
}

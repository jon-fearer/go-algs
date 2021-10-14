package algorithms

func MergeSort(data []int) ([]int, error) {
	left := 0
	right := len(data) - 1
	work := make([]int, len(data))
	mergeSortSubset(&data, &work, left, right)
	return data, nil
}

func mergeSortSubset(data *[]int, work *[]int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		mergeSortSubset(data, work, left, middle)
		mergeSortSubset(data, work, middle + 1, right)
		merge(data, work, left, middle, right)
	}
}

func merge(data *[]int, work *[]int, left int, middle int, right int) {
	for i := left; i <= right; i++ {
		(*work)[i] = (*data)[i]
	}

	workLeft := left
	workRight := middle + 1
	currentIndex := left

	for workLeft <= middle && workRight <= right {
		if (*work)[workLeft] <= (*work)[workRight] {
			(*data)[currentIndex] = (*work)[workLeft]
			workLeft++
		} else {
			(*data)[currentIndex] = (*work)[workRight]
			workRight++
		}
		currentIndex++
	}

	remaining := middle - workLeft
	for i := 0; i <= remaining; i++ {
		(*data)[currentIndex+ i] = (*work)[workLeft+ i]
	}
}

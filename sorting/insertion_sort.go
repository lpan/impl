package sorting

// mutates slice
func InsertionSort(slice []int) {
	if len(slice) < 2 {
		return
	}

	for i := 1; i < len(slice); i++ {
		// assume that 0 -> i are sorted
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				swap(slice, j, j-1)
			}
		}
	}

	return
}

// mutates slice
func swap(slice []int, i0, i1 int) {
	slice[i0], slice[i1] = slice[i1], slice[i0]
}

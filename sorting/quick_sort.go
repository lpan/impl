package sorting

// does not mutate slice
// space complexity - O(n)
func QuickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	p := pivot(slice)
	left, right := partition(slice, p)

	return append(
		append(QuickSort(left), p),
		QuickSort(right)...,
	)
}

func partition(slice []int, pivot int) (left []int, right []int) {
	for _, v := range slice {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			continue
		}
	}
	return
}

// pick first element as the pivot
func pivot(slice []int) int {
	return slice[0]
}

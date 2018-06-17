package sorting

// QuickSort does not mutate slice
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

// out of place
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

// in-place partition, return the pivot index
// elements smaller than pivot will be on the left of pivot
func partition1(slice []int) int {
	pivot := slice[len(slice)-1]

	// Hoare's in-place partition
	i := -1
	for j := range slice {
		if slice[j] < pivot {
			i++
			tmp := slice[i]
			slice[i] = slice[j]
			slice[j] = tmp
		}
	}

	// insert pivot
	for j := len(slice) - 1; j > i+1; j-- {
		slice[j] = slice[j-1]
	}
	slice[i+1] = pivot

	return i + 1
}

// pick first element as the pivot
func pivot(slice []int) int {
	return slice[0]
}

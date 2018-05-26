package sorting

func merge2(slice []int, leftLow, leftHigh, rightLow, rightHigh int) {
	result := make([]int, rightHigh-leftLow)

	left, right := leftLow, rightLow
	for k, _ := range result {
		if left == leftHigh {
			result[k] = slice[right]
			right++
		} else if right == rightHigh {
			result[k] = slice[left]
			left++
		} else if slice[left] < slice[right] {
			result[k] = slice[left]
			left++
		} else {
			result[k] = slice[right]
			right++
		}
	}

	for k, _ := range result {
		slice[leftLow+k] = result[k]
	}
}

// MergeSort2 mutates the parameter
// ensures - slice range low:high are sorted
func MergeSort2(slice []int, low, high int) {
	if low == high-1 {
		return
	}

	mid := (low + high) / 2
	MergeSort2(slice, low, mid)
	MergeSort2(slice, mid, high)
	merge2(slice, low, mid, mid, high)
}

// --------------------------------------------

func merge(a0, a1 []int) []int {
	result := make([]int, len(a0)+len(a1))

	var i0, i1 int
	for k, _ := range result {
		if i0 == len(a0) {
			result[k] = a1[i1]
			i1++
		} else if i1 == len(a1) {
			result[k] = a0[i0]
			i0++
		} else if a0[i0] < a1[i1] {
			result[k] = a0[i0]
			i0++
		} else {
			result[k] = a1[i1]
			i1++
		}
	}

	return result
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])
	sorted := merge(left, right)

	return sorted
}

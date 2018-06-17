package sorting

import "fmt"

// CountSortLeft uses the left boundary of the position
// Complexity: O(n + K)
//   - n: number of elements
//   - K: range of the element
// out-of place
// no need to allocate memory on heap
// stable
func CountSort(slice []int) []int {
	// find the range
	// O(n)
	l := slice[0]
	h := slice[0]
	for _, el := range slice {
		if el < l {
			l = el
		}

		if el > h {
			h = el
		}
	}

	// count each element
	// O(n)
	counts := make([]int, h-l+1)
	for _, el := range slice {
		counts[el-l]++
	}

	// find left boundary
	// O(K)
	var pos []int
	for i := range counts {
		if i == 0 {
			pos = append(pos, 0)
		} else {
			pos = append(pos, pos[i-1]+counts[i-1])
		}
	}

	result := make([]int, len(slice))
	for _, el := range slice {
		idx := pos[el-l]
		pos[el-l]++
		result[idx] = el
	}

	return result
}

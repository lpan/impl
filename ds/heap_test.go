package ds

import (
	"testing"

	st "github.com/lpan/impl/sorting/sortingtest"
)

func prepend(slice []int, i int) []int {
	return append([]int{i}, slice...)
}

func Test_BuildHeapFixUp(t *testing.T) {
	nums := st.Input(15)
	h := NewHeapBad(nums)

	var result []int
	for {
		n, ok := h.DeleteMax()
		if !ok {
			break
		}

		result = prepend(result, n)
	}

	st.AssertSorted(t, result)
}

func Test_BuildHeapHeapify(t *testing.T) {
	nums := st.Input(7)
	h := NewHeap(nums)

	var result []int
	for {
		n, ok := h.DeleteMax()
		if !ok {
			break
		}

		result = prepend(result, n)
	}

	st.AssertSorted(t, result)
}

package ds

import (
	"testing"

	st "github.com/lpan/interviews/sorting/sortingtest"
)

func prepend(slice []int, i int) []int {
	return append([]int{i}, slice...)
}

func Test_HeapSort(t *testing.T) {
	h := NewHeap()
	nums := st.Input(15)

	for _, n := range nums {
		h.Insert(n)
	}

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

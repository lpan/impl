package sorting

import (
	"testing"

	st "github.com/lpan/interviews/sorting/sortingtest"
)

func Test_MergeSort(t *testing.T) {
	result := MergeSort(st.Input(10))
	st.AssertSorted(t, result)
}

func Test_MergeSort2(t *testing.T) {
	input := st.Input(10)
	MergeSort2(input, 0, len(input))
	st.AssertSorted(t, input)
}

package sorting

import (
	"testing"

	st "github.com/lpan/interviews/sorting/sortingtest"
)

func Test_QuickSort(t *testing.T) {
	input := st.Input(10)
	result := QuickSort(input)
	st.AssertSorted(t, result)
}

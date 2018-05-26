package sorting

import (
	"testing"

	st "github.com/lpan/interviews/sorting/sortingtest"
)

func Test_InsertionSort(t *testing.T) {
	input := st.Input(10)
	InsertionSort(input)
	st.AssertSorted(t, input)
}

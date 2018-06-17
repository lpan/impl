package sorting

import (
	"fmt"
	"testing"

	st "github.com/lpan/impl/sorting/sortingtest"
)

func Test_QuickSort0(t *testing.T) {
	input := st.Input(10)
	result := QuickSort(input)
	st.AssertSorted(t, result)
}

// in-place partition algo
func Test_partition1(t *testing.T) {
	input := st.Input(8)
	fmt.Println(input)
	fmt.Println(partition1(input))
}

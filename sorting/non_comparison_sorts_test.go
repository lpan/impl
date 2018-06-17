package sorting

import (
	"fmt"
	"testing"

	st "github.com/lpan/impl/sorting/sortingtest"
)

func Test_CountSort(t *testing.T) {
	result := CountSort([]int{23, 87, 82, 82, 82, 83, 1, 3})
	fmt.Println(result)
	st.AssertSorted(t, result)
}

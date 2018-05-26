package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

func genInput(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(100)
	}
	return result
}

func Test_MergeSort(t *testing.T) {
	result := MergeSort(genInput(10))
	if s := sort.IntsAreSorted(result); !s {
		t.Errorf("input: %v is not sorted", result)
	}
}

func Test_MergeSort2(t *testing.T) {
	input := genInput(10)
	MergeSort2(input, 0, len(input))
	if s := sort.IntsAreSorted(input); !s {
		t.Errorf("input: %v is not sorted", input)
	}
}

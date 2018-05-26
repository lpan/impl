package sortingtest

import (
	"math/rand"
	"sort"
	"testing"
)

func Input(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(100)
	}
	return result
}

func AssertSorted(t *testing.T, slice []int) {
	if s := sort.IntsAreSorted(slice); !s {
		t.Errorf("input: %v is not sorted", slice)
	}
}

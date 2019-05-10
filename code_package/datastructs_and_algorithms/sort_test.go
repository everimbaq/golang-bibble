package datastructs_and_algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {

	var sortCases = []struct {
		input  []int
		expect []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 5, 2, 4, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, v := range sortCases {
		assert.Equal(t, QuickSort(v.input), v.expect)
	}
}

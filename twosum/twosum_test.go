package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For nums %v and target %d, expected %v but got %v", test.nums, test.target, test.output, result)
			}
		})
	}
}

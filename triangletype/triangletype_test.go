package main

import (
	"reflect"
	"testing"
)

func TestTriangleType(t *testing.T) {
	tests := []struct {
		nums   []int
		output string
	}{

		{
			nums:   []int{3, 3, 3},
			output: "equilateral",
		},
		{
			nums:   []int{3, 4, 5},
			output: "scalene",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := triangleTypes(test.nums)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For nums %v, expected %v but got %v", test.nums, test.output, result)
			}
		})
	}
}

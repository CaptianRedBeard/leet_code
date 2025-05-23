package main

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs      []string
		output string
		}{

		{
			strs:   []string{"flower","flow","flight"},
			output: "fl",
		},
		{
			strs:      []string{"dog","racecar","car"},
			output: "",
		},

	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := longestCommonPrefix(test.strs)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For nums %v, expected %v but got %v", test.strs, test.output, result)
			}
		})
	}
}
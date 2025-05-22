package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x      int
		output bool
	}{

		{
			x:      120,
			output: true,
		},
		{
			x:      -121,
			output: false,
		},
		{
			x:      10,
			output: false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := isPalindrome(test.x)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For nums %v, expected %v but got %v", test.x, test.output, result)
			}
		})
	}
}

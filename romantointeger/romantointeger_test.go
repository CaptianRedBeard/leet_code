package main

import (
	"reflect"
	"testing"
)

func TestTriangleType(t *testing.T) {
	tests := []struct {
		s   string
		output int
	}{

		{
			s:   "III",
			output: 3,
		},
		{
			s:   "LVIII",
			output: 58,
		},
		{
			s: "MCMXCIV",
			output: 1994,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := romanToInt(test.s)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For s %v, expected %v but got %v", test.s, test.output, result)
			}
		})
	}
}

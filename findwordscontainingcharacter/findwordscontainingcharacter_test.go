package main

import (
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	tests := []struct {
		words  []string
		x      byte
		output []int
	}{
		{
			words:  []string{"leet", "code"},
			x:      'e', // Use a byte literal for 'e'
			output: []int{0, 1},
		},
		{
			words:  []string{"abc", "bcd", "aaaa", "cbc"},
			x:      'a', // Use a byte literal for 'a'
			output: []int{0, 2},
		},
		{
			words:  []string{"abc", "bcd", "aaaa", "cbc"},
			x:      'z', // Use a byte literal for 'z'
			output: []int{},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := findWordsContaining(test.words, test.x)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("For words %v and byte %v, expected %v but got %v", test.words, test.x, test.output, result)
			}
		})
	}
}

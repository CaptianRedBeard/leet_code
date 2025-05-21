package main

import (
	"strconv"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for len(str) > 1 {
		if str[0] != str[len(str)-1] {
			return false
		}
		str = str[1 : len(str)-1]
	}

	return true

}

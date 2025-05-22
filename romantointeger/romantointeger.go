package main

func romanToInt(s string) int {

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	count := 0
	prevValue := 0

	for index := len(s) -1 ; index >= 0; index-- {
		currentValue := romanMap[rune(s[index])]

		if currentValue < prevValue {
			count -= currentValue
		} else {
			count += currentValue
		}

		prevValue = currentValue
	}

	return count
    
}
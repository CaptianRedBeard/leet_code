package main

func findWordsContaining(words []string, x byte) []int {
	foundInstances := []int{}

	if len(words) == 0 {
		return foundInstances
	}

	for index, word := range words {
		if containsByte(word, x) {
			foundInstances = append(foundInstances, index)
		}

	}

	return foundInstances
}

func containsByte(word string, x byte) bool {

	for i := 0; i < len(word); i++ {
		if word[i] == x {
			return true
		}
	}
	return false
}

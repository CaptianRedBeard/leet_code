package main

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := ""
    index := 0

    for {
        if index >= len(strs[0]) { 
            break
        }
        letter := rune(strs[0][index])
        for _, word := range strs {
            if index >= len(word) || letter != rune(word[index]) {
                return prefix 
            }
        }

        prefix += string(letter)
        index++
    }

    return prefix
}
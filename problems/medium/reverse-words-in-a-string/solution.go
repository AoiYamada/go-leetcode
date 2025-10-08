package main

import "strings"

// Medium: ReverseWordsInAString
// Solution for reverse-words-in-a-string (medium)
func reverseWords(s string) string {
	words := strings.Fields(s)
	reversedWords := make([]string, len(words))

	for pos, word := range words {
		reversedWords[len(words)-pos-1] = word
	}

	return strings.Join(reversedWords, " ")
}

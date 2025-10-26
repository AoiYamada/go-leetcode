package main

// Medium: DetermineIfTwoStringsAreClose
// Solution for determine-if-two-strings-are-close (medium)
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// Letter Counts for words
	lc1 := countLetters(word1)
	lc2 := countLetters(word2)

	if len(lc1) != len(lc2) {
		return false
	}

	for letter := range lc1 {
		if _, ok := lc2[letter]; !ok {
			return false
		}
	}

	// Count occurence of count for letter counts
	cc1 := countCounts(lc1)
	cc2 := countCounts(lc2)

	if len(cc1) != len(cc2) {
		return false
	}

	for count, occurence := range cc1 {
		if cc2[count] != occurence {
			return false
		}
	}

	return true
}

// countLetters returns a frequency map of runes in a string
func countLetters(s string) map[rune]int {
	count := make(map[rune]int)

	for _, letter := range s {
		count[letter]++
	}

	return count
}

func countCounts(lc map[rune]int) map[int]int {
	count := make(map[int]int)

	for _, val := range lc {
		count[val]++
	}

	return count
}

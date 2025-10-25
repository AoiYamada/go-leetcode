package main

type void struct{}

// Easy: UniqueNumberOfOccurrences
// Solution for unique-number-of-occurrences (easy)
func uniqueOccurrences(arr []int) bool {
	counterMap := make(map[int]int)

	for _, num := range arr {
		counterMap[num]++
	}

	countSet := make(map[int]void)

	for _, count := range counterMap {
		if _, ok := countSet[count]; ok {
			// count exists
			return false
		} else {
			countSet[count] = void{}
		}
	}

	return true
}

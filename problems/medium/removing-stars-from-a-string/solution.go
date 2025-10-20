package main

// Medium: RemovingStarsFromAString
// Solution for removing-stars-from-a-string (medium)
func removeStars(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '*' {
			stack = stack[:len(stack)-1]
		} else {
			// Push the character
			stack = append(stack, char)
		}
	}

	return string(stack)
}

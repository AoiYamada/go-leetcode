package main

// Medium: RemovingStarsFromAString
// Solution for removing-stars-from-a-string (medium)
func removeStars2(s string) string {
	result := make([]byte, len(s))
	idx := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			idx--
		} else {
			result[idx] = s[i]
			idx++
		}
	}

	return string(result[:idx])
}

// removeStarsB uses []rune, which means each character is treated as a Unicode code point (up to 4 bytes). When you range over a string with for _, char := range s, Go internally decodes UTF-8 bytes into runes on each iteration. This decoding overhead doesn't exist in removeStarsA. Plus, converting []rune back to a string requires re-encoding those runes back to UTF-8.
// func removeStarsB(s string) string {
// 	result := make([]rune, len(s))
// 	idx := 0

// 	for _, char := range s {
// 		if char == '*' {
// 			idx--
// 		} else {
// 			result[idx] = char
// 			idx++
// 		}
// 	}

// 	return string(result[:idx])
// }

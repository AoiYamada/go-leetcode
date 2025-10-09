package main

// Easy: CanPlaceFlowers
// Solution for can-place-flowers (easy)
func canPlaceFlowers(flowerbed []int, n int) bool {
	avaliableCount := 0
	plantFailed := false

	for pos := 0; pos < len(flowerbed); pos++ {
		lookForward := 2
		if plantFailed {
			lookForward = 3
		}
		target := pos + lookForward

		for checkPos := pos; checkPos < target; checkPos++ {
			current := 0
			if checkPos < len(flowerbed) {
				current = flowerbed[checkPos]
			}

			if current == 1 || checkPos > len(flowerbed) {
				plantFailed = true
				pos = checkPos
				break
			}

			if checkPos == pos+lookForward-1 {
				pos = checkPos
				plantFailed = false
				avaliableCount++
			}
		}

		if avaliableCount >= n {
			return true
		}
	}

	return false
}

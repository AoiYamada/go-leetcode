package main

// Easy: CanPlaceFlowers
// Solution for can-place-flowers (easy)
func canPlaceFlowers2(flowerbed []int, n int) bool {
	for pos := 0; pos < len(flowerbed); {
		if n == 0 {
			return true
		}

		if flowerbed[pos] == 1 {
			pos += 2
			continue
		}

		prevAvaliable := pos-1 < 0 || flowerbed[pos-1] == 0
		nextAvaliable := pos+1 >= len(flowerbed) || flowerbed[pos+1] == 0

		if prevAvaliable && nextAvaliable {
			n--
			pos += 2
			continue
		}

		if nextAvaliable {
			pos += 1
			continue
		}

		pos += 3
	}

	return n <= 0
}

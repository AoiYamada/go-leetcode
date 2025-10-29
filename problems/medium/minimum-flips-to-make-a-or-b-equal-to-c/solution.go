package main

// Medium: MinimumFlipsToMakeAOrBEqualToC
// Solution for minimum-flips-to-make-a-or-b-equal-to-c (medium)
func minFlips(a int, b int, c int) int {
	count := 0

	// 0 >> 1 == 0
	for a > 0 || b > 0 || c > 0 {
		aBit := a & 1
		bBit := b & 1
		cBit := c & 1

		if cBit == 0 {
			// both aBit and bBit need to flip to 0
			count += aBit + bBit
		} else if aBit == 0 && bBit == 0 {
			// cBit is 1 here
			// only 1 bit is 1 ensures `|` results 1
			count++
		}

		// simultaneously shift 1 bit
		a >>= 1
		b >>= 1
		c >>= 1
	}

	return count
}

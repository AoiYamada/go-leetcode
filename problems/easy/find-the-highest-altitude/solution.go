package main

// Easy: FindTheHighestAltitude
// Solution for find-the-highest-altitude (easy)
func largestAltitude(gain []int) int {
	altitude := 0
	highest := 0

	for _, diff := range gain {
		altitude += diff

		if highest < altitude {
			highest = altitude
		}
	}

	return highest
}

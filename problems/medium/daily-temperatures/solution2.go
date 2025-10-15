package main

// Medium: DailyTemperatures
// Solution for daily-temperatures (medium)
func dailyTemperatures2(temperatures []int) []int {
	monoStack := []int{} // unhandled idx of days
	results := make([]int, len(temperatures))

	for pos, temperature := range temperatures {
		for len(monoStack) > 0 &&
			// this means the temperature of a unhandled day is lower then the temperature of the day at `pos`
			temperatures[monoStack[len(monoStack)-1]] < temperature {
			idx := monoStack[len(monoStack)-1]
			results[idx] = pos - idx
			monoStack = monoStack[:len(monoStack)-1]
		}

		monoStack = append(monoStack, pos)
	}

	return results
}

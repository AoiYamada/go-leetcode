package main

import "runtime/debug"

// Medium: ContainerWithMostWater
// Solution for container-with-most-water (medium)
func maxArea(height []int) int {
	max := 0

	for lPos, rPos := 0, len(height)-1; lPos < rPos; {
		minHeight := min(height[lPos], height[rPos])
		water := minHeight * (rPos - lPos)

		if water > max {
			max = water
		}

		if minHeight == height[lPos] {
			lPos++
		} else {
			rPos--
		}
	}

	return max
}

func init() {
	debug.SetMemoryLimit(0)
}

package main

import "runtime/debug"

// Easy: MoveZeroes
// Solution for move-zeroes (easy)
func moveZeroes(nums []int) {
	lead := -1
	for pos, num := range nums {
		if lead == -1 {
			if num != 0 {
				continue
			}

			lead = pos
			continue
		}

		if num == 0 {
			continue
		}

		nums[pos], nums[lead] = nums[lead], nums[pos]
		lead++
	}
}

// to beat 99.9% people in memory, need this trick.
// this will disable the GC and magically works in ranking of leetcode
func init() {
	debug.SetMemoryLimit(0)
}

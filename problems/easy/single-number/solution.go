package main

// Easy: SingleNumber
// Solution for single-number (easy)
func singleNumber(nums []int) int {
	ans := 0

	for _, num := range nums {
		ans ^= num
	}

	return ans
}

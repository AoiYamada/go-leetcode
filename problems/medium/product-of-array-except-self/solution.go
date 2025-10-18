package main

// Medium: ProductOfArrayExceptSelf
// Solution for product-of-array-except-self (medium)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	results := make([]int, n)
	results[0] = 1

	for i := 0; i < n-1; i++ {
		results[i+1] = nums[i] * results[i]
	}

	surfixProduct := 1
	for j := n - 1; j >= 0; j-- {
		results[j] = results[j] * surfixProduct
		surfixProduct = surfixProduct * nums[j]
	}

	return results
}

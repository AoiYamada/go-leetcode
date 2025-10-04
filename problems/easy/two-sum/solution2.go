package twosum

// Solution2 - Hash Map Approach
// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSumHashMap(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return nil
}

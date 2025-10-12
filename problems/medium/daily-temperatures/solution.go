package main

type stack [][2]int

func (s *stack) push(v [2]int) {
	*s = append(*s, v)
}

func (s *stack) pop() [2]int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res
}

// Medium: DailyTemperatures
// Solution for daily-temperatures (medium)
func dailyTemperatures(temperatures []int) []int {
	monoStack := make(stack, 0)
	results := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(monoStack) > 0 {
			if monoStack[len(monoStack)-1][0] <= temperatures[i] {
				monoStack.pop()
				continue
			}

			break
		}

		if len(monoStack) > 0 {
			results[i] = monoStack[len(monoStack)-1][1] - i
		}

		monoStack.push([2]int{temperatures[i], i})
	}

	return results
}

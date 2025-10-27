package main

// Easy: NumberOfRecentCalls
// Solution for number-of-recent-calls (easy)
type RecentCounter struct {
	queue []int
	// used in Ping2 only
	cursor int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

// this approach is slow even added `debug.SetMemoryLimit(0)`, probably `counter.queue[pos:]` is an heavy operation
func (counter *RecentCounter) Ping(t int) int {
	counter.queue = append(counter.queue, t)

	if t < 3000 {
		return len(counter.queue)
	}

	for pos, time := range counter.queue {
		if time < t-3000 {
			continue
		}

		counter.queue = counter.queue[pos:]
		break
	}

	return len(counter.queue)
}

func (counter *RecentCounter) Ping2(t int) int {
	counter.queue = append(counter.queue, t)

	for pos := counter.cursor; pos < len(counter.queue); pos++ {
		if counter.queue[pos] < t-3000 {
			continue
		}

		counter.cursor = pos
		break
	}

	return len(counter.queue) - counter.cursor
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

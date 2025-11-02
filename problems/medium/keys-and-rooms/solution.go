package main

// Medium: KeysAndRooms
// Solution for keys-and-rooms (medium)
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visted := make([]bool, n)
	vistedCount := 0
	keys := []int{0}

	for len(keys) > 0 {
		key := keys[0]
		if visted[key] {
			keys = keys[1:]
			continue
		}

		visted[key] = true
		vistedCount++

		if vistedCount == n {
			return true
		}

		newKeys := rooms[key]
		// for doing dfs
		// slices.Reverse(newKeys)
		keys = append(keys[1:], newKeys...)
	}

	return false
}

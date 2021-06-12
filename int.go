package split

// IntSlice splits []int by value of element itself.
func IntSlice(intSlice []int, nPart int, greedy bool) []int {
	return Slice(intSlice, func(i int) int { return intSlice[i] }, nPart, greedy)
}

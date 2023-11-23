package lc

// MaxArea (11 | Medium)
func MaxArea(height []int) int {
	lo, hi := 0, len(height)-1
	res := 0
	for lo < hi {
		h := min(height[lo], height[hi])
		res = max(res, h*(hi-lo))
		// Move the pointers based on which line is shorter
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return res
}

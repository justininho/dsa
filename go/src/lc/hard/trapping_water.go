package hard

// trap (42 | Hard)
// Approach 4: Using 2 pointers
func trap(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

package easy

/*
(66 | Easy) Plus One

You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
*/
func plusOne(digits []int) []int {
	carryover := 1
	for i := len(digits) - 1; i >= 0 && carryover == 1; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			carryover = 0
		}
	}
	if carryover == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

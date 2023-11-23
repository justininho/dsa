package lc

import "sort"

// threeSums (15 | Medium)
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSums(i, nums, &result)
		}
	}
	return result
}

func twoSums(i int, nums []int, result *[][]int) {
	lo := i + 1
	hi := len(nums) - 1
	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		} else {
			*result = append(*result, []int{nums[i], nums[lo], nums[hi]})
			lo++
			hi--
			// to avoid duplicates
			for lo < hi && nums[lo-1] == nums[lo] {
				lo++
			}
		}
	}
}

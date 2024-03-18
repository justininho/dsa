class Solution {
    /// #26 Easy
    /// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
    func removeDuplicates(_ nums: inout [Int]) -> Int {
        var i = 0
        for j in 1..<nums.count {
            if nums[i] != nums[j] {
                i += 1
                nums[i] = nums[j]
            }
        }
        return i + 1
    }
}


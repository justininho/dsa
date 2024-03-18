
class Solution {
    /// #27 Easy
    /// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
    func removeElement(_ nums: inout [Int], _ val: Int) -> Int {
        var i = 0
        var j = nums.count - 1
        while i <= j {
            if nums[i] == val {
                nums[i] = nums[j]
                j -= 1
            } else {
                i += 1
            }
        }
        return j + 1
    }
}
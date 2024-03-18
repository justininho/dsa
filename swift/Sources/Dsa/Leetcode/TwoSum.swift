class Solution {
    /// 1. Two Sum
    /// https://leetcode.com/problems/two-sum/
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        let dict = [Int: Int]()
        for (i, _) in nums.enumerated() {
            let compliment = target - nums[i]
            if let complimentIndex = dict[compliment] {
                return [complimentIndex, i]
            } else {
                dict[nums[i]] = i
            }
        }
        return [-1, -1]
    }
}
from typing import List

def binary_search(target: int, nums: List[int]) -> int:
  high = len(nums)
  low = 0
  while low < high:
    mid = low + (high - low) // 2
    if target > nums[mid]:
      low = mid + 1
    elif target < nums[mid]:
      high = mid
    else:
      return mid

  return -1

print(binary_search(3, [1, 2, 2, 3, 4, 5, 6, 7, 10]))
export function binarySearch(target: number, nums: number[]): number {
  let high = nums.length;
  let low = 0;
  while (low < high) {
    const mid = Math.floor((low + (high - low)) / 2) // prevents int overflow
    if(target > nums[mid]) low = mid + 1;
    else if (target < nums[mid]) high = mid;
    else return mid;
  }
  return -1;
}

console.log(binarySearch(3, [1, 2, 2, 3, 4, 5, 6, 7, 10]))


export function quicksort(nums: number[]) {
  if(nums.length <= 1) return nums;
  const p = partition(nums);
  return quicksort(nums.slice(0, p)).concat(quicksort(nums.slice(p)))
}

function partition(nums: number[]): number {
  const pivot = nums[nums.length-1];
  let p = nums.length;
  let i = 0;
  while(i < p) {
    if(nums[i] > pivot) {
      --p;
      const temp = nums[p];
      nums[p] = nums[i];
      nums[i] = temp;
    }
    i++;
  }
  return p;
}

console.log(quicksort([4,2,6,3,8,5]))
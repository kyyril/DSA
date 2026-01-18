function twoSum(nums: number[], target: number): number[] {
  // Create a map to store numbers we've seen so far and their indices
  const numMap = new Map<number, number>();

  // Iterate through the array
  for (let i = 0; i < nums.length; i++) {
    // Calculate the complement we need to find
    const complement = target - nums[i];

    // Check if the complement exists in our map
    if (numMap.has(complement)) {
      // If found, return the indices of the complement and current number
      return [numMap.get(complement)!, i];
    }

    // If not found, store the current number and its index
    numMap.set(nums[i], i);
  }

  // This should never be reached given the problem constraints
  return [];
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// console.log(twoSumm(nums, 10));

// brute force
function twoSumm(nums: number[], target: number) {
  for (let i = 0; i <= nums.length - 1; i++) {
    for (let j = i + 1; j <= nums.length - 1; j++) {
      if (nums[i] + nums[j] == target) {
        return [nums[i], nums[j]];
      }
    }
  }
}

const nums: number[] = [1, 2, 3, 4, 5, 6, 7, 8];

// sum number
let sum = 0;
for (let num of nums) {
  sum += num;
}
console.log(sum);

// find max number
let max = nums[0];

for (let num of nums) {
  if (num > max) {
    max = num;
  }
}
console.log(max);

// reverse array
let left = 0;
let right = nums.length - 1;

while (left < right) {
  const temp = nums[left];
  nums[left] = nums[right];
  nums[right] = temp;

  left += 1;
  right -= 1;
}
console.log(nums);

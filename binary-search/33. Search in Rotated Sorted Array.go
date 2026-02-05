func search(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        // Left part is sorted
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // Right part is sorted
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}



// 1. Check mid
// 2. Check which side is sorted
// 3. Check target in that side
// 4. Move left or right

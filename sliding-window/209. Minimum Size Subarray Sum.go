func minSubArrayLen(target int, nums []int) int {
    left := 0
    sum := 0
    minLen := len(nums)+1//paling panjang

    for right:=0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            currentLen := right - left + 1

            if currentLen < minLen {
                minLen = currentLen
            }
            sum -= nums[left]
            left++
        }
    }
    //if minLen tidak pernah berubah, berarti tidak ada subarray yang cocok
    if minLen > len(nums) {
        return 0
    }
    return minLen
}
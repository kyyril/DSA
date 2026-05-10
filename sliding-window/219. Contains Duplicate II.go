func containsNearbyDuplicate(nums []int, k int) bool {
    window := make(map[int] bool)
    for i:=0; i < len(nums); i++ {
        if window[nums[i]] {
            return true
        }

        window[nums[i]] = true

        // buang angka yang paling kiri sebelum pindah ke angka berikutnya
        if len(window) > k {
            delete(window, nums[i-k])
        }
    }
    return false
}
// Boyer-Moore Voting Algorithm

func majorityElement(nums []int) int {
    candidate := 0
    count := 0

    for _, num := range nums {
        // Jika count nol, kita pilih kandidat baru
        if count == 0 {
            candidate = num
        }
        
        // Jika angka sama dengan kandidat, tambah poin. Jika beda, kurangi.
        if num == candidate {
            count++
        } else {
            count--
        }
    }

    return candidate
}
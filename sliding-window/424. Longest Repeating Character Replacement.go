func characterReplacement(s string, k int) int {
    left := 0
    freq := make(map[byte] int)
    max_count := 0
    result := 0
    for right:=0; right < len(s); right++ {
        freq[s[right]]++ //add value each of right
        //check if freq[s[right]] > max_count -> update max
        if freq[s[right]] > max_count {
            max_count = freq[s[right]]
        }

        for (right - left + 1) - max_count > k {
            freq[s[left]]--
            left++
        }
        window := right - left + 1

        if window > result {
            result = window
        }
    }
    return result
}


// Expand until valid
// Shrink while valid
// Save best answer
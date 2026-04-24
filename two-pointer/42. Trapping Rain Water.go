func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0

	for left < right {
		// Update Max di kedua sisi terlebih dahulu
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		// Bandingkan mana yang lebih rendah untuk menentukan air
		if leftMax < rightMax {
			// Sisi kiri lebih rendah, hitung air di sisi kiri
			res += leftMax - height[left]
			left++
		} else {
			// Sisi kanan lebih rendah, hitung air di sisi kanan
			res += rightMax - height[right]
			right--
		}
	}

	return res
}
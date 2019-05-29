package wrr

// gcd greatest common divisor
func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}

	return a
}

// max return the max num and its index in int slice
func max(nums []int) (int, int) {
	i, max := 0, nums[0]
	for idx, num := range nums {
		if num > max {
			i = idx
			max = num
		}
	}

	return i, max
}


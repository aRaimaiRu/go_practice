package maximumproductsubarray

func maxProduct(nums []int) int {
	var i, largest, current, currentPositive int

	if len(nums) == 1 {
		return nums[0]
	}
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			current = nums[i]
			currentPositive = nums[i]
			largest = Max(largest, current)
			i++
			break
		}
	}

	for ; i < len(nums); i++ {

		if nums[i] == 0 {
			current = 1
			currentPositive = 1
			largest = Max(largest, nums[i])
			continue
		}

		current = current * nums[i]
		currentPositive = Max(currentPositive*nums[i], nums[i])
		largest = Max(largest, current)
		if currentPositive < 0 {
			currentPositive = 1
			continue
		}
		largest = Max(largest, currentPositive)

	}

	for i = len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			current = nums[i]
			currentPositive = nums[i]
			i--
			break
		}
	}

	for ; i >= 0; i-- {
		if nums[i] == 0 {
			current = 1
			currentPositive = 1
			largest = Max(largest, nums[i])
			continue
		}

		current = current * nums[i]
		currentPositive = Max(currentPositive*nums[i], nums[i])
		largest = Max(largest, current)
		if currentPositive < 0 {
			currentPositive = 1
			continue
		}
		largest = Max(largest, currentPositive)

	}

	return largest
}

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	lastest := nums[0]
	for i := 1; i < len(nums); i++ {
		if lastest == nums[i] {
			return true

		}
		lastest = nums[i]

	}
	return false
}

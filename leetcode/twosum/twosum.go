package twosum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if m[v] > 0 {
			return []int{m[v] - 1, i}
		}
		m[target-v] = i + 1

	}
	return []int{0, 0}
}

package exercise

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			if i == j {
				continue
			}
			return []int{i, j}
		}
	}
	return nil
}

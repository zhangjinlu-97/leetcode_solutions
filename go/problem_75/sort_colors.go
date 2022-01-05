package problem_75

// 类似于快排中的partition
func sortColors(nums []int) {
	redEnd, blueStart := -1, len(nums)
	index := 0
	for index < blueStart {
		if nums[index] < 1 {
			redEnd++
			swap(nums, index, redEnd)
			index++
		} else if nums[index] > 1 {
			blueStart--
			swap(nums, index, blueStart)
		} else {
			index++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

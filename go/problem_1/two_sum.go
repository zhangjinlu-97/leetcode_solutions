package problem_1

// 直接使用hash表即可
func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := idxMap[target-num]; ok {
			return []int{i, index}
		}
		idxMap[num] = i
	}
	return []int{}
}

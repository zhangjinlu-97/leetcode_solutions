package problem_78

func subsets(nums []int) [][]int {
	var set []int
	var ret [][]int
	process(0, nums, &set, &ret)
	return ret
}

func process(i int, nums []int, set *[]int, ret *[][]int) {
	if i == len(nums) {
		*ret = append(*ret, append([]int{}, *set...))
		return
	}
	*set = append(*set, nums[i])
	process(i+1, nums, set, ret) // select nums[i]
	*set = (*set)[:len(*set)-1]
	process(i+1, nums, set, ret) // not select nums[i]
}

package problem_189

// 循环替换 时间：O(N) 空间：O(1)
func rotate1(nums []int, k int) {
	n := len(nums)
	for i, cnt := 0, n; cnt > 0; i++ {
		val, cur := nums[i], i
		for ok := true; ok; ok = cur != i {
			next := (cur + k) % n
			nums[next], val, cur = val, nums[next], next
			cnt--
		}
	}
}

// 反转数组 时间：O(N) 空间：O(1)
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 额外空间 时间：O(N) 空间：O(N)
func rotate3(nums []int, k int) {
	n := len(nums)
	ans := make([]int, n)
	for i := range nums {
		ans[(i+k)%n] = nums[i]
	}
	for i := range nums {
		nums[i] = ans[i]
	}
}

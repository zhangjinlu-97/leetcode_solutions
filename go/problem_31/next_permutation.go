package problem_31

//1. 从后向前遍历，找到第一个升序对(i, j);
//2. 再从后向前找到第一个比nums[i]大的位置k，并交换i和k上的数;
//3. 将i+1到最后的数组逆序;
//注意：如果第一步始终找不到升序对，说明数组整体降序，直接将数组整体逆序即可。
func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums, 0, n-1)
		return
	}
	j := n - 1
	for ; j > i; j-- {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

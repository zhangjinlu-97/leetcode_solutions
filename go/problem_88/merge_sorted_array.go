package problem_88

// 倒着遍历，这样额外空间复杂度为O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, index := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[index] = nums1[p1]
			p1--
		} else {
			nums1[index] = nums2[p2]
			p2--
		}
		index--
	}
	for p2 >= 0 {
		nums1[index] = nums2[p2]
		p2--
		index--
	}
}

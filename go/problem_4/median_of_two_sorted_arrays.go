package problem_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	up, down := (m+n+1)/2, (m+n+2)/2 // 计算上中位数与下中位数是第几个数
	return float64(kthMin(nums1, nums2, 0, m-1, 0, n-1, up)+
		kthMin(nums1, nums2, 0, m-1, 0, n-1, down)) / 2
}

// 求两数组中第K小的数
func kthMin(nums1, nums2 []int, l1, r1, l2, r2, k int) int {
	m, n := r1-l1+1, r2-l2+1
	if m > n {
		return kthMin(nums2, nums1, l2, r2, l1, r1, k) // 始终让nums1比nums2短，这样最先为0的一定是nums1
	}
	if m == 0 { // nums1长度为0，因此整体的第k个，就是nums2的第k个
		return nums2[l2+k-1]
	}
	if k == 1 { // 当k为1时，两数组第一个位置上较小的那个即为结果
		return min(nums1[l1], nums2[l2])
	}
	kth1, kth2 := l1+min(m, k/2)-1, l2+min(n, k/2)-1
	if nums1[kth1] <= nums2[kth2] {
		return kthMin(nums1, nums2, kth1+1, r1, l2, r2, k-(kth1-l1+1))
	} else {
		return kthMin(nums1, nums2, l1, r1, kth2+1, r2, k-(kth2-l2+1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

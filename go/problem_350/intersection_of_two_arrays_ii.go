package problem_350

import "sort"

// 方法1 哈希表
func intersect(nums1 []int, nums2 []int) []int {
	var bigger, smaller []int
	if len(nums1) > len(nums2) {
		bigger = nums1
		smaller = nums2
	} else {
		bigger = nums2
		smaller = nums1
	}

	numMap := make(map[int]int, len(smaller))
	ans := make([]int, 0)
	for _, num := range smaller {
		numMap[num] = numMap[num] + 1
	}
	for _, num := range bigger {
		if cnt, ok := numMap[num]; ok && cnt > 0 {
			ans = append(ans, num)
			numMap[num] = numMap[num] - 1
		}
	}
	return ans
}

// 假定数组已经排好序
func intersect1(nums1 []int, nums2 []int) []int {
	// 模拟数组排序
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	ans := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ans = append(ans, nums1[i])
			i++
			j++
		}
	}
	return ans
}

// 内存限制
// todo: 外排序

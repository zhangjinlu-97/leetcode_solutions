package problem_215

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func findKthLargest(nums []int, k int) int {
	return process(nums, 0, len(nums)-1, k-1)
}

func process(nums []int, l, r, index int) int {
	if l >= r {
		return nums[index]
	}
	pivot := nums[l+rand.Intn(r-l)]
	area := partition(nums, l, r, pivot)
	if index >= area[0] && index <= area[1] {
		return nums[index]
	} else if index <= area[0] {
		return process(nums, l, area[0]-1, index)
	} else {
		return process(nums, area[1]+1, r, index)
	}
}

func partition(nums []int, l, r, pivot int) []int {
	less := l - 1
	more := r + 1
	index := l
	for index < more {
		if nums[index] > pivot {
			less++
			swap(nums, less, index)
			index++
		} else if nums[index] < pivot {
			more--
			swap(nums, more, index)
		} else {
			index++
		}
	}
	return []int{less + 1, more - 1}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

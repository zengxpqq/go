package algorithm

import "fmt"

// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	if len(nums) == 1 {
		fmt.Println(nums)
		return
	}
	realRotate := len(nums) - (k % len(nums)) // 实际需要轮转的次数
	start := append([]int{}, nums[realRotate:]...)
	end := append([]int{}, nums[0:realRotate]...)
	result := append([]int{}, start...)
	result = append(result, end...)
	copy(nums, result)
}

func reverse(nums []int) {
	for i, n := 0, len(nums)-1; i < len(nums)/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}
}

// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func rotate1(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k %= len(nums)
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

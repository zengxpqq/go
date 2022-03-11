package sort

import "fmt"

func sortedSquares(nums []int) []int {
	var (
		squares []int
		mid     int
	)

	mid = -1
	for i, num := range nums {
		if num < 0 {
			mid = i
		}
		squares = append(squares, num*num)
	}

	// 全是正数
	if mid == -1 {
		return squares
	}

	// 全是负数
	if mid == len(nums)-1 {
		var result []int
		for i := len(squares) - 1; i >= 0; i-- {
			result = append(result, squares[i])
		}
		return result
	}

	// 1 3 5
	// 4 3 2
	left := append([]int{}, squares[0:mid+1]...) // 负数部分
	right := append([]int{}, squares[mid+1:]...) // 正数部分

	squares = []int{}
	var i, j int
	i = 0
	j = len(left) - 1

	for i < len(right) && j >= 0 {
		if right[i] <= left[j] {
			squares = append(squares, right[i])
			i += 1
		} else {
			squares = append(squares, left[j])
			j -= 1
		}
	}

	for i < len(right) {
		squares = append(squares, right[i])
		i += 1
	}

	for j >= 0 {
		squares = append(squares, left[j])
		j -= 1
	}

	return squares
}

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

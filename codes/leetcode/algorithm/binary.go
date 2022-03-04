package algorithm

func BinarySearch1(nums []int, target int) int {
	numsLen := len(nums)
	startIndex := 0
	endIndex := numsLen

	for i := 0; i < numsLen; i++ {
		if startIndex == endIndex {
			break
		}
		midIndex := (startIndex + endIndex) / 2
		midValue := nums[midIndex]
		if midValue == target {
			return midIndex
		} else if midValue > target {
			endIndex = midIndex
		} else {
			startIndex = midIndex
		}
	}

	return -1
}

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		midIndex := (left + right) / 2
		midValue := nums[midIndex]
		if midValue == target {
			return midIndex
		} else if midValue > target {
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}
	return -1
}

package algorithm

func moveZeroes(nums []int) {
	var (
		countZeros int
		index      int
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZeros += 1
		} else {
			nums[index] = nums[i]
			index += 1
		}
	}
	for i := 0; i < countZeros; i++ {
		nums[len(nums)-1-i] = 0
	}
}

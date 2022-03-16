package 数据结构

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsDuplicate1(nums []int) bool {
	numsMap := make(map[int]struct{}, 0)
	for i := 0; i < len(nums); i++ {
		if _, exist := numsMap[nums[i]]; exist {
			return true
		} else {
			numsMap[nums[i]] = struct{}{}
		}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	numsMap := make(map[int]struct{}, 0)
	for i := 0; i < len(nums); i++ {
		if _, exist := numsMap[nums[i]]; exist {
			return true
		} else {
			numsMap[nums[i]] = struct{}{}
		}
	}

	return false
}

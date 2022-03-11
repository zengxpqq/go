package algorithm

//给你一个下标从 1 开始的整数数组numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数target 的两个数。
//如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
//
//以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
//
//你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
//
//你所设计的解决方案必须只使用常量级的额外空间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func twoSum1(numbers []int, target int) []int {
	for i := range numbers {
		remain := target - numbers[i]

		left := i + 1
		right := len(numbers) - 1

		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == remain {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < remain {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nil
}

func twoSum2(numbers []int, target int) []int {
	var (
		left  = 0
		right = len(numbers) - 1
	)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left += 1
		} else {
			right += 1
		}
	}
	return nil
}

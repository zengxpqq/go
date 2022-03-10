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

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

// 第一个错误的版本
// 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-bad-version
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 输入：n = 5, bad = 4
//输出：4
//解释：
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5) -> true
//调用 isBadVersion(4) -> true
//所以，4 是第一个错误的版本。
//
// 输入：n = 1, bad = 1
//输出：1
func firstBadVersion(n int) int {
	left := 0
	right := n
	for {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			if isBadVersion(mid - 1) {
				right = mid - 1
			} else {
				return mid
			}
		} else {
			if !isBadVersion(mid + 1) {
				left = mid + 1
			} else {
				return mid + 1
			}
		}
	}
}

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 在一个有序数组中找第一个大于等于 target 的下标
func findInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	index := len(nums)
	for left <= right {
		midIndex := left + (right-left)/2
		if target <= nums[midIndex] {
			index = midIndex
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}

	return index
}

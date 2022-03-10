package sort

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

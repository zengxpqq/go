package algorithm

func maxAreaOfIsland(grid [][]int) int {
	dx := []int{1, 0, 0, -1}
	dy := []int{0, 1, -1, 0}

	// 岛屿长宽
	landN := len(grid)
	landM := len(grid[0])

	// 当前最大面积
	maxArea := 0

	for n := 0; n < landN; n++ {
		for m := 0; m < landM; m++ {
			if grid[n][m] == 1 {
				grid[n][m] = 0
				currArea := 1
				areaList := append([][]int{}, []int{n, m})

				for i := 0; i < len(areaList); i++ {
					for j := 0; j < 4; j++ {
						dn, dm := areaList[i][0]+dx[j], areaList[i][1]+dy[j]
						if dn >= 0 && dn < landN && dm >= 0 && dm < landM && grid[dn][dm] == 1 {
							areaList = append(areaList, []int{dn, dm})
							grid[dn][dm] = 0
							currArea += 1
						}
					}
				}

				if currArea > maxArea {
					maxArea = currArea
				}
			}
		}
	}

	return maxArea
}

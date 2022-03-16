package algorithm

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

// 广度优先算法
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currentColor := image[sr][sc]
	if currentColor == newColor {
		return image
	}
	image[sr][sc] = newColor

	n := len(image)
	m := len(image[0])
	fillColor := append([][]int{}, []int{sr, sc})

	for i := 0; i < len(fillColor); i++ {
		for j := 0; j < 4; j++ {
			dn, dm := fillColor[i][0]+dx[j], fillColor[i][1]+dy[j]
			if dn >= 0 && dn < n && dm >= 0 && dm < m && image[dn][dm] == currentColor {
				fillColor = append(fillColor, []int{dn, dm})
				image[dn][dm] = newColor
			}
		}
	}

	return image
}

// 深度优先算法 dfs
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	currentColor := image[sr][sc]
	if currentColor != newColor {
		dfs(image, sr, sc, currentColor, newColor)
	}
	return image
}

func dfs(image [][]int, sr int, sc int, currentColor, newColor int) {
	if image[sr][sc] == currentColor {
		image[sr][sc] = newColor
		for j := 0; j < 4; j++ {
			dn, dm := sr+dx[j], sc+dy[j]
			if dn >= 0 && dn < len(image) && dm >= 0 && dm < len(image[0]) {
				dfs(image, dn, dm, currentColor, newColor)
			}
		}
	}

}

package algorithm

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 1)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 4)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 5)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 6)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 7)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 8)

	rotate([]int{1, 2}, 1)
	rotate([]int{1}, 1)
}

func TestRotate1(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(data, 8)
	//rotate1(data, 1)
	//rotate1(data, 2)
	//rotate1(data, 3)
	//rotate1(data, 4)
	//rotate1(data, 5)
	//rotate1(data, 6)
	//rotate1(data, 7)
	//rotate1(data, 8)
	//rotate1([]int{1, 2}, 1)
	//rotate1([]int{1}, 1)
	fmt.Println(data)
}

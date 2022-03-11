package algorithm

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	data := []int{0, 1, 0, 3, 12}
	moveZeroes(data)
	fmt.Println(data)
}

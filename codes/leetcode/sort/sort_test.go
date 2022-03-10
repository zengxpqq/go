package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{-1}
	result := sortedSquares(data)
	fmt.Println(result)
}

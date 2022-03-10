package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{-1, 0, 3, 5, 9, 12}
	assert.Equal(t, -1, BinarySearch(data, -2))
	assert.Equal(t, 0, BinarySearch(data, -1))
	assert.Equal(t, 3, BinarySearch(data, 5))
	assert.Equal(t, 5, BinarySearch(data, 12))
	assert.Equal(t, -1, BinarySearch(data, 91))
}

func TestBinaryInsert(t *testing.T) {
	var index int
	index = findInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println("index: ", index)

	//index = findInsert([]int{1, 2, 4}, 3)
	//fmt.Println("index: ", index)
}

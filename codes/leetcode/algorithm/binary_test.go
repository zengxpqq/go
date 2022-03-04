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
	data := []int{1}
	target := 10
	left := 0
	right := len(data) - 1

	for left <= right {
		midIndex := (left + right) / 2
		midValue := data[midIndex]
		if midValue == target {

			return
		} else if midValue > target {
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}

	fmt.Println(left, right)
}

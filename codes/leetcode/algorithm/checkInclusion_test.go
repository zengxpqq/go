package algorithm

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	result := checkInclusion("abc", "ccccbbbbaaaa")
	fmt.Println(result)
}

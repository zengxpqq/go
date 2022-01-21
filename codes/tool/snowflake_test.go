package tool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSnow(t *testing.T) {
	gen := InitGenerator(1, 1)
	snowID, err := gen.GetNextID()
	assert.NoError(t, err)
	assert.Equal(t, 19, len(strconv.Itoa(int(snowID))))
	fmt.Println(snowID)
}

func BenchmarkSnow(b *testing.B) {
	b.StartTimer()
	gen := InitGenerator(1, 1)
	for i := 0; i < b.N; i++ {
		_, _ = gen.GetNextID()
	}
	b.StopTimer()
}

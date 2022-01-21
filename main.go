package main

import (
	"fmt"
	"gocook/codes/tool"
)

func main() {
	fmt.Println("========")
	gen := tool.InitGenerator(1, 1)
	snowID, err := gen.GetNextID()
	if err != nil {
		return
	}
	fmt.Println(snowID)
}

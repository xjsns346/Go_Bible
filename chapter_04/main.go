package main

import (
	"fmt"
	"gobible/chapter_04/Slice"
)

func main() {
	x := []int{1, 5, 9, 6, 0, 0}

	Slice.Rotate(x, 2)
	fmt.Printf("%v", x)
}

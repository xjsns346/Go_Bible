package main

import (
	"fmt"
	"gobible/chapter_04/Slice"
)

func main() {
	x := [6]int{1, 5, 9, 6, 0, 0}
	ptr := &x
	Slice.Reverseint(ptr)
	fmt.Printf("%v", x)
}

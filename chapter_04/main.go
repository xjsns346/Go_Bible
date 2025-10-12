package main

import (
	"fmt"
	"gobible/chapter_04/Slice"
)

func main() {
	x := []int{1, 5, 9, 6}

	b := Slice.Reverseint(x)
	fmt.Printf("%v", b)
}

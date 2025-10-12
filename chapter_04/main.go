package main

import (
	"fmt"
	"gobible/chapter_04/Slice"
	"time"
)

func main() {
	strings := []string{"a", "b", "b", "b", "c", "d", "e", "e"}

	strings = Slice.CompactStrings(strings)
	fmt.Printf("%v", strings)
	time.Sleep(10 * time.Second)
}

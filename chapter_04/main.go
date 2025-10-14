package main

import (
	"fmt"
	"gobible/chapter_04/Slice"
	"time"
)

func main() {
	nums := "123   456   789"
	bytes := []byte(nums)
	s := "he   llo,  world  !"
	sbytes := []byte(s)
	sbytes = Slice.RemoveExtraSpaces(sbytes)
	bytes = Slice.RemoveExtraSpaces(bytes)
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%v\n", sbytes)
	fmt.Printf("%v\n", string(sbytes))
	fmt.Printf("%v\n", string(bytes))
	time.Sleep(10 * time.Second)
}

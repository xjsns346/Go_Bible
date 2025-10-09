package main

import (
	"fmt"
	"gobible/chapter_02/Package/popcount"
)

func main() {

	c := popcount.PopCount(0x123456789ABCDEF0)
	fmt.Printf("所含有的1bit的个数为%v", c)

}

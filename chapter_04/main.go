package main

import (
	"crypto/sha256"
	"fmt"
	"gobible/chapter_04/Array"
)

func main() {
	a := sha256.Sum256([]byte("Hello,World!"))
	b := sha256.Sum256([]byte("hello,world!"))
	fmt.Printf("%v,%v", a, b)
	count := Array.DifCount(a, b)

	fmt.Printf("不同的bit位数为%v", count)

}

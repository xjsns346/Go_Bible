package main

import "fmt"

func main() {

	s1 := "你好，世界！"
	fmt.Printf("%d\n", len(s1))
	for i := range len(s1) {
		fmt.Println(s1[i])

	}

}

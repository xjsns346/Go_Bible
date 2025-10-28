package main

import (
	"fmt"
	funcvalue "gobible/chapter_05/FuncValue"
)

func main() {
	a := "126678"
	b := "66"
	c, err := funcvalue.Expand(a, b, funcvalue.F_example)
	if err != nil {
		fmt.Printf("报错:%v", err)
	}
	fmt.Printf("%s", c)
}

//flag包的基本使用，用于简便对命令行参数的使用。

package main

import (
	"flag"
	"fmt"
	"strings"
)

// n,sep都是指针类型，需要使用*n来调用。
var n = flag.Bool("n", false, "默认换行")
var sep = flag.String("s", " ", "以空格连接")

func main() {
	flag.Parse() //flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）。
	fmt.Print(strings.Join(flag.Args(), *sep))
	//默认n是false，所以会执行println,执行换行，否则不会换行。
	if !*n {
		fmt.Println()
	}
}

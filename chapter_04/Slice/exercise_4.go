// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回。
package slice

import (
	"fmt"
	"unicode"
)

func RemoveExtraSpaces(b []byte) []byte {
	if len(b) == 0 {
		fmt.Printf("输入了一个空切片")
	}

	i := 0
	for j := 1; j < len(b); j++ {
		//这个代码块实现了函数目的，对于多个连续空格，仅保留一个空格。
		if !unicode.IsSpace(rune(b[j])) || !unicode.IsSpace(rune(b[i])) {
			i++
			b[i] = b[j]
		}
	}
	fmt.Println(i)
	return b[:i+1]
}

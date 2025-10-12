//练习 4.3： 重写reverse函数，使用数组指针代替slice。

package Slice

import "unsafe"

func Reverseint(x []int) []int {

	b := []int{}

	for i := len(x) - 1; i > (-1); i-- {
		a := (*int)(unsafe.Pointer(&x[i]))
		b = append(b, *a)
	}

	return b

}

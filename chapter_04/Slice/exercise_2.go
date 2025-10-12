//练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。

package Slice

func Rotate(s []int, n int) {
	n = n % len(s)                   // 防止越界
	tmp := append([]int{}, s[:n]...) // 复制前n个元素
	copy(s, s[n:])                   // 将后面的元素前移
	copy(s[len(s)-n:], tmp)          // 将前n个元素放到末尾
}

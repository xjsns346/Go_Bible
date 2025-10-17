// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
// 题目的要求是不能破坏UTF-8编码，如果不额外分配内存，那就要使用双重反转，先按照字节进行反转，在根据UTF-8的规则，找到每一个字符所对应的字节范围，在进行反转。

package slice

func ReverseByte(x []byte) {
	i := 0
	j := len(x) - 1
	for i != j {
		x[i], x[j] = x[j], x[i]
		i++
		j--
	}

} //还没有完成。

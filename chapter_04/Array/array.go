// 练习 4.1  编写一个函数，计算两个SHA256哈希码中不同bit的数目。
package Array

//import "gobible/chapter_02/Package/popcount"

func DifCount(a, b [32]byte) (count int) {
	//第一种方法，逐个取出数组中的byte,并逐个比较最后一位bit位是否相等，来计算，较为繁琐。
	for i, s := range a {
		g := b[i]
		for bit := range 8 {
			//通过移位和与运算，来取出最后一位，并进行比较。
			if (s>>bit)&1 != (g>>bit)&1 {
				count++
			}
		}
	}
	//第二种方法，逐个取出其中的byte，并进行异或运算，然后通过popcount函数，查表来统计1的个数，就可得出结果，简单方便。
	/*for i := range 32 {
		count += popcount.PopCount8(a[i] ^ b[i])
	}
	*/
	return
}

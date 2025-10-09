//popcount函数实现，用于统计64位无符号数中，1bit的个数，采用的查表法。

package popcount

//匿名函数+立即执行的形式，等同于init函数。
var pc [256]uint8 = func() (pc [256]uint8) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}() //在最后加上()来实现立即执行。

//传入一个uint64的参数，返回含有的1bit的个数
func PopCount(x uint64) (count int) {
	//这里是右移0，8，16，... ，56位，并取出末尾的8位，根据pc数组来得出含有的1bit的个数，最后相加得出全部的个数。
	count = int(pc[uint8(x>>(0*8))] +
		pc[uint8(x>>(1*8))] +
		pc[uint8(x>>(2*8))] +
		pc[uint8(x>>(3*8))] +
		pc[uint8(x>>(4*8))] +
		pc[uint8(x>>(5*8))] +
		pc[uint8(x>>(6*8))] +
		pc[uint8(x>>(7*8))])
	return

}

//练习 3.13

package Const

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB //到这里已经超出了uint64的范围，所以这些常量为无类型常量。
	YB = 1000 * ZB
)

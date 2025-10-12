// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package Slice

func CompactStrings(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			//由于slice当形参时，是值传递类型，所以切片的len ，cap是一个副本，所以需要返回。
			strings = Delete(strings, i)
			//为了避免多个重复的数据，需要进行i--。
			i--
		}
	}
	return strings

}

func Delete(strings []string, i int) []string {
	copy(strings[i:], strings[i+1:])
	strings = strings[:len(strings)-1]
	//由于slice当形参时，是值传递类型，所以切片的len ，cap是一个副本，所以需要返回。
	return strings
}

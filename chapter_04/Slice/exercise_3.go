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

//AI版本
//代码更加精简，并且思路独特。
//demo :  []string{"a", "b", "b", "b", "c", "d", "e"}转变为{"a", "b", "c", "d", "e","d", "e"}最后取出{"a","b","c","d","e"}

func eliminateAdjacentDuplicates(s []string) []string {
	if len(s) == 0 {
		return s
	}
	//这个思路可以学习，轻松解决了多个值重复的问题，并实现了原地修改，没有分配新的内存。
	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}

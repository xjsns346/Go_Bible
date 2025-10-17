// 练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

//这个程序是有缺陷的，由于input.Split(bufio.ScanWords)函数是按照空格来分割，在一篇实际的文章中无法分割标点符号，可以继续优化。

package wordfreq

import (
	"bufio"
	"fmt"
	slice "gobible/chapter_04/Slice"
	"os"
	"sort"
)

func Wordfreq(filePath string) {

	countsWord := make(map[string]int) //存储单词出现的次数。
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "报错 :%v", err)
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords) //以单词来分割
	for input.Scan() {
		countsWord[input.Text()]++ //统计单词出现次数
	}

	//通过新建一个Slice来实现对map的排序打印。
	s := []int{}
	for _, v := range countsWord {
		s = append(s, v)
	}
	sort.Ints(s) //进行从小到大的排序。

	//这里不可以使用反向map表，因为不知道countsWord中有多少value相等的数据，简单进行反向map，会造成数据丢失。

	slice.EliminateAdjacentDuplicates(s) //去除多余的相等元素，仅保留一个。
	//按照单词出现的次数，从小到大依次打印，但由于map的随机特性，对于多个出现次数相同的单词，不能保证每次打印顺序相同。
	for _, i := range s {
		for k, v := range countsWord {
			if v == i {
				fmt.Printf("%v的出现次数为%v\n", k, v)
			}
		}

	}

}

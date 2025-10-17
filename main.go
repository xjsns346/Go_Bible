//练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	countsLetter := 0
	countsDigit := 0
	countsSpace := 0
	countsPunct := 0

	file, err := os.Open("C:/Users/陈子鸣/Desktop/test.txt.txt") //获取了一个File类型的指针，实现了io.Reader接口。
	if err != nil {
		fmt.Fprintf(os.Stderr, "报错:%v", err)
	}
	in := bufio.NewReader(file)
	//循环结束条件在循环内部，当读到文件结尾时，break退出循环。
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		//当读到文件结尾时退出。
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		//判断是否为无效字符，并统计个数。
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			countsLetter++
		}
		if unicode.IsDigit(r) {
			countsDigit++
		}
		if unicode.IsSpace(r) {
			countsSpace++
		}
		if unicode.IsPunct(r) {
			countsPunct++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Printf("字母个数 :%d\n", countsLetter)
	fmt.Printf("数字个数 :%d\n", countsDigit)
	fmt.Printf("空白字符个数 :%d\n", countsSpace)
	fmt.Printf("标点符号个数 :%d\n", countsPunct)
}

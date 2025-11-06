// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
package countfun

import (
	"bufio"
	"bytes"
)

// 统计字节数
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil

}

// 统计单词数，(使用scan.Split(bufio.ScanWords)一般会有些许误差，具体可以再了解)
type WordCount int

func (w *WordCount) Write(p []byte) (int, error) {
	//bytes.NewReader(p)返回一个*bytes.Reader ,满足io.Reader接口
	//bufio.NewScanner()接受一个满足io.Reader接口的实例，返回一个*bufio.Scanner(这是一个结构体)
	scan := bufio.NewScanner(bytes.NewReader(p))
	//定义如何分割数据
	scan.Split(bufio.ScanWords)
	//根据分割规则，逐个取出token
	for scan.Scan() {
		//根据取出的token个数,来增加w的值
		*w++
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}
	//len(p)返回的是字节数
	return len(p), nil
}

// 这里与上面几乎同理
type LineCount int

func (l *LineCount) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		*l++
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

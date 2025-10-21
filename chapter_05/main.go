package main

import (
	"fmt"
	recursion "gobible/chapter_05/Recursion"

	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 定义一个map，来存储所有节点的元素，以及出现的次数，如果大于1，就输出，实现了练习5.2
var DataSame = make(map[string]int)

func main() {
	//定义一个URL
	url := "https://golangstar.cn/go_series/introduction.html"

	//获取他的返回体，传给html.Parse,从而获取HTML文本的根节点。
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "报错为%v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "返回失败，状态码为:%v", err)
		os.Exit(1)
	}

	//获取根节点，依次进行递归，获取页面中的链接。
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//visit函数返回一个字符串切片，每一个字符串都是一个链接。
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	fmt.Printf("------------------------------------------------\n")
	//简单的判断，输出同名元素以及出现的次数
	for k, v := range DataSame {
		if v > 1 {
			fmt.Printf("元素%q出现次数%d\n", k, v) //%q会给输出的字符串带上双引号，并且对其中的特殊字符进行转义。
		}
	}

	fmt.Printf("------------------------------------------------\n")

	s := recursion.CountText(doc)
	for i, v := range s {
		fmt.Printf("第%d个text节点的内容为%q\n", i, v)
	}

}

func visit(links []string, n *html.Node) []string {
	DataSame[n.Data]++

	//这个判断语句是根据HTML中存储链接的语句来设计的。
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//使用递归来搜索HTML树，隐藏着一个非常经典的「深度优先遍历（DFS）」过程。
	/*for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}*/

	// 练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
	// 递归访问第一个子节点
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	// 递归访问下一个兄弟节点
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

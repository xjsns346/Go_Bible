package main

import (
	"fmt"
	funcvalue "gobible/chapter_05/FuncValue"

	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 定义一个map，来存储所有节点的元素，以及出现的次数，如果大于1，就输出，实现了练习5.2
var DataSame = make(map[string]int)

func main() {
	//定义一个URL
	url := "https://www.runoob.com/markdown/md-tutorial.html"

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

	//获取根节点.
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//输出DOM结构。
	vist(doc, funcvalue.StartElement, funcvalue.EndElement)

	fmt.Printf("---------------------------------------------------")
	//练习5.8实现

	funcvalue.ForEachNode(doc, "pull", funcvalue.ElementByID)

}

func vist(n *html.Node, start, end func(*html.Node)) {
	if start != nil {
		start(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		vist(c, start, end)
	}

	if end != nil {
		end(n)
	}

}

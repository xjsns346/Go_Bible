// 练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，
// 返回与这些标签名匹配的所有元素。

//易错点:标签名在doc.Data部分。

package variable

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, n ...string) []*html.Node {

	result := make([]*html.Node, 0)
	//通过创建一个map，利用查表的方式来寻找符合要求的节点，更加便捷。
	seen := make(map[string]bool)
	for _, v := range n {
		seen[v] = true
	}

	//将查找与递归的逻辑代码放在一个匿名函数在中，避免对seen与result的重复定义，并且不需要有返回值直接将节点添加进result。
	var find func(doc *html.Node)
	find = func(doc *html.Node) {
		if doc.Type == html.ElementNode {
			if seen[doc.Data] {
				result = append(result, doc)
			}
		}

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}

	}
	//对初始节点调用
	find(doc)

	return result

}

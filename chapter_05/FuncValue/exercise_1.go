//练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，
// 文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。编写测试，验证程序输出的格式正确。

package funcvalue

//对HTML没有系统性的学习，借助AI完成的，之后进行温习。

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func StartElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		//在 Go 的 html 包中：对于无属性的元素，n.Attr 通常是一个长度为 0 的切片（不是 nil）。
		//在GO中零值可用，进行range不会报错，会直接退出循环。
		if n.FirstChild == nil {
			//空元素，直接输出简写格式
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=%q", a.Key, a.Val)
			}
			fmt.Println("/>")
		} else {
			// 普通元素，正常输出
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=%q", a.Key, a.Val)
			}
			fmt.Println(">")
			depth++
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}

	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)

	}

}
func EndElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

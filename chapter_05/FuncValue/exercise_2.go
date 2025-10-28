//练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。
// 使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

package funcvalue

import (
	"fmt"

	"golang.org/x/net/html"
)

var bo = new(bool)

func init() {
	*bo = true
}

// 使用ElementByID函数，来查找id，如果找到了，将标志*bo改为false,就会提前退出，不会继续遍历后面的节点。
func ForEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool) {
	if *bo == false {
		return
	}
	if pre != nil {
		pre(n, id)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, id, pre)
	}

}

// 输入一个HTML节点，判断其属性中是否含有id="string",如果有就输出该节点，否则输出nil。
func ElementByID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				fmt.Printf("查找拥有该id元素的HTML元素如下\n")
				fmt.Printf("<%s %v>", n.Data, n.Attr)
				*bo = false
				return *bo
			}

		}
	}
	return *bo
}

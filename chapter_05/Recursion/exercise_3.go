//练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的。
//<script>和<style>元素的子节点为text节点，但是存储的是javascript代码以及CSS，所以要排除。
//当遇到 <script> 或 <style> 就直接 return（不遍历其子树），因为一般情况下，他们的子节点就只有textnode。

package recursion

import "golang.org/x/net/html"

var TextString = make([]string, 0)

//这个函数实现了对所有text节点文本的提取，
func CountText(n *html.Node) []string {

	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			n = n.NextSibling //将n赋值为他的兄弟节点，直接跳过了该子树。
		}
	}

	if n.Type == html.TextNode {
		TextString = append(TextString, n.Data)
	}

	//深度优先搜索，递归调用函数。
	if n.FirstChild != nil {
		CountText(n.FirstChild)
	}
	if n.NextSibling != nil {
		CountText(n.NextSibling)
	}

	return TextString
}

//练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
//很简单，加上判断条件即可，不写了。

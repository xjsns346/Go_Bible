// 练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）
package revalues

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("打开网址%s失败,错误信息为%v", url, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("报错，状态码返回为:%d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {

	if n.Type == html.TextNode {
		ws := strings.Fields(n.Data)
		words += len(ws)
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i

	}
	return
}

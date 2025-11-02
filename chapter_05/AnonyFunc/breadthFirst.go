// 匿名函数章节中，广度优先搜索的实现细节。
// 以及匿名函数章节中多个函数的实现。
package anonyfunc

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 广度优先搜索函数
func BreadthFirst(f func(item string, file *os.File) []string, worklist []string, file *os.File) {
	seen := make(map[string]bool)
	//实现了一个动态判断，可以多次进行判断，如果使用if则只进行一次判断。
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			//使用一个map来避免对同一个string重复调用函数。
			if !seen[item] {
				seen[item] = true
				get := f(item, file)
				worklist = append(worklist, get...)
			}

		}

	}

}

//练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。

//抓取器实现

func Crawl(url string, file *os.File) []string {

	fmt.Println(url)
	seen := make(map[string]bool)
	seen[url] = true
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, item := range list {
		if !seen[item] {
			seen[item] = true
			fmt.Fprintln(file, item)
		}

	}
	return list
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				//将HTML属性中的相对路径或绝对路径转换成完整的URL
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(doc *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(doc)

	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(doc)

	}

}

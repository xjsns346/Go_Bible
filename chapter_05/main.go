package main

import (
	"fmt"
	variable "gobible/chapter_05/Variable"

	"golang.org/x/net/html"

	"net/http"
)

func main() {
	url := "https://xiaolincoding.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("打开URL失败\n")
		return
	}
	defer resp.Body.Close()
	doc, _ := html.Parse(resp.Body)
	result := variable.ElementsByTagName(doc, "img")

	for _, node := range result {
		fmt.Printf("%q", node.Data)
	}
}

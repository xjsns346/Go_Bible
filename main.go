package main

import (
	"fmt"
	revalues "gobible/chapter_05/Revalues"
)

func main() {
	w, i, err := revalues.CountWordsAndImages("https://golangstar.cn/go_series/introduction.html")
	if err != nil {
		fmt.Printf("统计失败，报错为%v\n", err)
	}
	fmt.Printf("文字数目为%d,图片数目为%d\n", w, i)

}

//练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

// 有关发送HTTP请求时的查询参数不了解，有关time包的使用也要学习，而且总个数与实际打印出来的issue总数不同，需要进一步了解。
package main

import (
	"fmt"
	github "gobible/chapter_04/Json"
	"log"
	"os"
	"time"
)

func main() {
	//获取解码后的结果。
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//获取当前时间
	t1 := time.Now()
	//这两个变量用于更新循环，来减少循环次数。
	issueOne := github.IssuesSearchResult{}
	issueTwo := github.IssuesSearchResult{}

	//通过time包来比较时间，输出符合要求的issue，同时更新issueOne用于下一次循环。后面同理
	fmt.Printf("创建时间在一个月之内的issue如下:\n")
	for i, item := range result.Items {
		if t1.Before(item.CreatedAt.AddDate(0, 1, 0)) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		} else {
			//在range循环中不要对切片进行剪切操作，采用创建一个新变量来储存不符合要求的item，在之后的循环中就减少了循环次数。后面同理
			issueOne.Items = append(issueOne.Items, result.Items[i])
		}
	}
	fmt.Printf("共%d个\n", len(result.Items)-len(issueOne.Items)) //输出个数

	fmt.Printf("创建时间在一年之内的issue如下:\n")
	for i, item := range issueOne.Items {
		if t1.Before(item.CreatedAt.AddDate(1, 0, 0)) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		} else {
			issueTwo.Items = append(issueTwo.Items, result.Items[i])
		}
	}
	fmt.Printf("共%d个\n", len(issueOne.Items)-len(issueTwo.Items))

	fmt.Printf("创建时间超过一年的issue如下:\n")
	for _, item := range issueTwo.Items {
		if t1.After(item.CreatedAt.AddDate(1, 0, 0)) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
	fmt.Printf("共%d个\n", len(issueTwo.Items))

}

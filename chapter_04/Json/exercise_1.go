package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// 定义Issue结构体
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// 传入查询参数， 返回一个结构体指针和错误。
func SearchIssues(terms []string) (*IssuesSearchResult, error) {

	//QueryEscape函数用于对 URL 查询字符串中的特殊字符进行百分号编码，防止因为特殊符号（如 &、=、+）导致参数解析错误。
	//strings.Join函数用于将一个字符串切片，使用第二个参数作为分隔符连接成一个字符串。
	q := url.QueryEscape(strings.Join(terms, " "))

	//向指定的网址加上查询字符串，发送http请求。
	resp, err := http.Get(IssuesURL + "?q=" + q)

	//常用的返回无效结果与报错的结构，nil表示返回无效结果，err表示 返回的错误。
	if err != nil {
		return nil, err
	}

	//如果返回的状态码不是OK说明请求失败，关闭请求体，返回状态码。
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	//对数据进行了解码，并检查是否报错。
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	//返回解码结果。
	return &result, nil
}

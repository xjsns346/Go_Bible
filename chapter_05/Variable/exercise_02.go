// 练习5.16：编写多参数版本的strings.Join。
package variable

import (
	"fmt"
	"strings"
)

// 用于将多个字符串切片相连
func Joins(sep string, strs ...[]string) (s string, err error) {
	if len(strs) < 2 {
		return "", fmt.Errorf("至少需要两个字符串。")
	}
	for i := range strs {
		t := strings.Join(strs[i], "")
		s += t
	}
	return s, nil
}

//用于将多个字符串相连

func JoinStrings(strs ...string) (s string, err error) {
	if len(strs) < 2 {
		return "", fmt.Errorf("至少需要两个字符串。")
	}
	for _, v := range strs {
		s += v

	}
	return s, nil
}

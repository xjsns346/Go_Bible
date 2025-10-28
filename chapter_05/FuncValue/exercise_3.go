// 练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
package funcvalue

import (
	"fmt"
	"strings"
)

func Expand(s, substr string, f func(string) string) (string, error) {
	if strings.Contains(s, substr) {
		fmt.Printf("在字符串%q中存在字串%q\n", s, substr)
		index := strings.Index(s, substr)
		s = s[:index] + f(substr) + s[index+len(substr):]
		return s, nil
	} else {
		return s, fmt.Errorf("未找到字串%q,无法更改。", substr)

	}

}

func F_example(s string) string {
	a := make([]rune, len(s))
	for i := range a {
		a[i] = rune('a' + i)
	}
	return string(a)
}

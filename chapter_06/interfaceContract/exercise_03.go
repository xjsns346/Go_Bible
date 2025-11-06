// 练习 7.3： 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。
package countfun

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//这个String接口的实现只是简单的通过fmt.printf打印单个*tree的·内容,可以优化成DOM树的结构，懒得搞了。
func (t *tree) String() string {
	if t == nil {
		return "nil"
	}

	return fmt.Sprintf("{value : %d\nleft : %v\nright : %v}", t.value, t.left, t.right)

}

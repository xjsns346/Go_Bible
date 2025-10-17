// 使用一个二叉树来实现一个插入排序
// 有函数的迭代，需要多看看
package btsort

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}
	back(values[:0], root)
}

//复用add函数来实现构建二叉搜索树，需要返回树指针，来更新节点。
func add(r *tree, v int) *tree {
	if r == nil {
		r = new(tree)
		r.value = v
		return r
	}
	if r.value > v {
		r.left = add(r.left, v)
	} else {
		r.right = add(r.right, v)
	}
	return r
}

//通过中序遍历来将二叉树转换成切片，进而输出。
func back(values []int, t *tree) []int {

	if t != nil {
		values = back(values, t.left)
		values = append(values, t.value)
		values = back(values, t.right)
	}
	return values
}

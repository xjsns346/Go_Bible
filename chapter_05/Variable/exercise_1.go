//练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

package variable

import "fmt"

func max(items ...int) (maxnum int, err error) {
	if len(items) == 0 {
		fmt.Printf("没有参数输入\n")
		return 0, fmt.Errorf("至少需要一个参数。")
	}
	maxnum = items[0]
	for _, v := range items {
		if maxnum < v {
			maxnum = v
		}
	}
	return maxnum, nil
}

func min(items ...int) (minnum int, err error) {
	if len(items) == 0 {
		fmt.Printf("没有参数输入\n")
		return 0, fmt.Errorf("至少需要一个参数。")
	}
	minnum = items[0]
	for _, v := range items {
		if minnum > v {
			minnum = v
		}
	}
	return minnum, nil
}

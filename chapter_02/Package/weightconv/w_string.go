// 定义输出的string方法。
package weightconv

import "fmt"

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (k Kg) String() string {
	return fmt.Sprintf("%g kg", k)
}

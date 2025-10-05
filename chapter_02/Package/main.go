//练习2.2 获取输入，并进行数制转换。

package main

import (
	"fmt"
	"gobible/chapter_02/Package/weightconv"
	"os"
	"strconv"
)

func main() {
	var p weightconv.Pound
	var k weightconv.Kg
	var a float64
	if len(os.Args[1:]) < 1 {
		//从标准输入获取内容，使用fmt.Scanln()。
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("输入错误，请输入一个数字，例如 10 或 3.5")
			return
		}
		p = weightconv.Pound(a)
		k = weightconv.Kg(a)
		fmt.Printf("%.3f kg = %.3f lb\n", k, weightconv.KtoP(k))
		fmt.Printf("%.3f lb = %.3f kg\n", p, weightconv.PtoK(p))

	} else {
		for i, s := range os.Args[1:] {
			//由于获取的a是切片类型的数据，无法直接转成浮点数，所以要使用strconv包。
			a, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "转换第%v个参数时失败,报错:%s ", i, err)
			}
			p = weightconv.Pound(a)
			k = weightconv.Kg(a)
			fmt.Printf("%.3f kg = %.3f lb\n", k, weightconv.KtoP(k))
			fmt.Printf("%.3f lb = %.3f kg\n", p, weightconv.PtoK(p))

		}
	}

}

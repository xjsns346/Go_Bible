// 练习4.2实现   编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
package Array

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

// 调用flag包。
var useSHA384 = flag.Bool("sha384", false, "use SHA384 instead of SHA256")
var useSHA512 = flag.Bool("sha512", false, "use SHA512 instead of SHA256")

func Mult_sha() {
	//解析命令行参数。
	flag.Parse()
	data, _ := io.ReadAll(os.Stdin)
	//通过switch来实现算法之间的选择。
	switch {
	case *useSHA384:
		fmt.Printf("%x\n", sha512.Sum384(data))
	case *useSHA512:
		fmt.Printf("%x\n", sha512.Sum512(data))
	default:
		fmt.Printf("%x\n", sha256.Sum256(data))
	}
}

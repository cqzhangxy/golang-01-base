// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}

	// fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(os.Args[2:])
}

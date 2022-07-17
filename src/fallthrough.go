// 如果想在执行完某个分支的代码后继续执行后续分支的代码，可以使用fallthrough关键字来达到目的。

// fallthrough 需要是 case 分支中的最后一条语句。否则编译器会抛出错误:fallthrough statement out of place。
package main

import (
	"fmt"
)

func main() {
	switch i := 45; {
	case i < 10:
		fmt.Printf("%d is less than 10\n", i)
		fallthrough
	case i < 50:
		fmt.Printf("%d is less than 50\n", i)
		fallthrough
	case i < 100:
		fmt.Printf("%d is less than 100\n", i)
	}
}

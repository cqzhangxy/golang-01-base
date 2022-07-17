// else 后面可以添加 if 语句，可以处理三种以上分支逻辑。
// else if 分支数量没有限制。但为了代码可读性，在 if 语句后不要添加太多 else-if 分支。
package main

import "fmt"

func main() {
	var str string = "12312332"
	if len(str) == 0 {
		fmt.Println("This string is empty")
	} else if len(str) <= 5 {
		fmt.Println("string length <5")
	} else {
		fmt.Println("This string length >5")
	}
}

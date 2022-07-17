// if 语句
// if 语句具有以下格式。如果条件为真，则执行大括号内的语句。
// if 条件表达式 {
//     // do something
// }

// if condition-expression {
//     // do something
// }

package main

import "fmt"

func main() {
	var str string
	if str == "" {
		fmt.Println("The string is empty.")
	}

	// 条件表达式前可以有一个初始化语句。例如给一个变量赋值。
	// 初始化语句后需要添加分号
	Abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if a := Abs(-7); a > 5 {
		fmt.Println("The absolute value of -7 is greater than 5")
	}
}

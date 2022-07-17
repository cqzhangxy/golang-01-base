// switch 一般格式如下：
// switch 变量 {
// case 可能值1:
// 	// do something
// case 可能值2:
// 	// do something
// default:
// 	// do something
// }

// switch 关键字之后有一个变量，case 关键字之后是变量的可能值。可选的 default 分支一般放在最后。相当于 if-else 语句中的最后一个 else 分支。

// case 语句从上到下逐一测试，直到匹配成功时停止。

// 一旦一个分支匹配成功，整个switch代码块在执行相应代码后就会退出，这意味着你不需要使用 break 语句来表示结束。

package main

import (
	"fmt"
)

func main() {
	animal := "wangliang"
	switch animal {
	case "dog":
		fmt.Println("Animal is ", animal)
	case "cat":
		fmt.Println("Animal is ", animal)
	case "sheep":
		fmt.Println("Animal is ", animal)
	default:
		fmt.Println("Other Animal(Maybe)")
	}

	// 没有条件的switch
	// switch 语句的第二种形式是不提供任何要判断的值，然后在每个 case 分支中测试不同的条件。

	// 当任一分支的测试结果为真时，将执行该分支的代码。

	// 这看起来很像链式的 if/else 语句，但是比前者可读性更好。
	fmt.Println("==================================================")
	var num1 int = 7

	switch {
	case num1 < 0:
		fmt.Println("Num < 0")
	case num1 > 0 && num1 < 100:
		fmt.Println("num < 100")
	default:
		fmt.Println("Num >= 100")
	}

	// 包含初始化语句的switch
	fmt.Println("====================================")
	switch num2 := 1; num2 {
	case 1:
		fmt.Println("num2=", num2)
	case 2:
		fmt.Println("num2=", num2)
	case 3:
		fmt.Println("num2=", num2)
	}
}

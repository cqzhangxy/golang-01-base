// for 循环是 Go 中唯一的循环结构。

// 基本的 for 循环由三部分组成，它们用分号隔开：

// 初始化语句（init statement）
// 在第一次迭代前执行且仅执行一次。通常为一句短变量声明
// 条件表达式（condition expression）
// 每次迭代之前执行。如果条件为 false 则退出循环迭代，否则循环继续迭代。
// 修饰语句（modification statement）
// 每次迭代之后执行。通常为更新计数器。
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("This is the end %d interation\n", i)
	}

	fmt.Println("=====================")

	// for 循环结构中的初始化和修改部分是可选的。
	// 两个分号也可以省略。
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	for sum < 2000 {
		sum += sum
	}

	fmt.Println(sum)

	fmt.Println("=============================")
	// for 条件表达式可以省略，没有条件表达式的for循环本质上是无限循环。
	a := 0
	for {
		fmt.Println(a)
		a++
		time.Sleep(time.Second * 1)
	}
}

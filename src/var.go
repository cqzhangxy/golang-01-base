// 变量声明
// var 语句用于声明一个或多个变量，多个变量用逗号分隔，变量的类型在最后。如下面的结构所示
package main

import "fmt"

var a, b, c int

func main() {
	// var d string
	// var e, f bool

	// 声明多个不同类型的变量
	var (
		d string
		e bool
		f bool
	)
	fmt.Println(a, b, c, d, e, f)

	// 变量赋值
	// 变量在声明时可以指定一个初始值。
	// 使用赋值运算符（=）给变量赋值，多个值使用逗号分隔并和变量一一对应，格式如下：
	var m string = "string"
	var n bool = true
	var i bool = false
	fmt.Println(m, n, i)

	var g string = "gg"
	var j, k bool = true, false
	fmt.Println(g, j, k)

	// 短变量声明
	// 在函数中，短变量声明语句 := 可以代替 var 声明
	// 它的语法很简洁，快速实现变量的声明与赋值。不过短变量声明语句仅可以在函数内使用。
	jayAge, zhangsanAge := 24, 25
	isAdult := true
	fmt.Println(jayAge, zhangsanAge, isAdult)

	// 仅声明没有初始化的变量的值是零值（zero-valued）
	// 不同数据类型的零值不同。例如：
	// 字符串（string）类型的零值是""
	// 整数（int）的零值是 0
	// 浮点数（float32、float64）的零值是 0.0
	// 布尔值（bool）的零值是 false
	// 映射（map）的零值是 nil

	var (
		q string
		w int
		r float32
		t bool
	)
	fmt.Println(q, w, r, t)

	// 变量类型转换
	// Go 语言中没有隐式转换，类型的转换需要显式声明。
	// 表达式 T(v) 将值 v 转换为类型 T。
	// int、 uint、float 可以使用 T(v) 表达式相互转换。
	var z int = 21

	fmt.Printf("int64(z) type is %T\n", int64(z))
	fmt.Printf("float64(i) type is %T\n", float64(z))
	fmt.Printf("uint(float64(z)) type is %T\n", uint(float64(z)))

	var x []byte = []byte("golang")

	fmt.Printf("string(x) type is %T\n", string(x))

	// 在声明一个变量而不指定其类型时,变量的类型由右值推导得出。
	v := 1
	fmt.Printf("v type is %T\n", v)

	v1 := 1.1
	fmt.Printf("v1 type is %T\n", v1)

	v2 := "golang"
	fmt.Printf("v2 type is %T\n", v2)

	// 常量的声明与变量类似，只不过是使用 const 关键字。常量的定义格式和变量的声明语法类似
	const const1 int = 1
	const world = "世界"
	fmt.Println("hello", world)

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

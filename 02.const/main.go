package main

import "fmt"

/**
## 02.const 常量
- 常量的定义是相对于变量的，它在整个声明周期中是不允许进行更改的
- 常量的定义采用const
- 常量如果被重新赋值会报错 Cannot assign
- 常量如果定义了，没有使用是不会报错的
- 常量在编译过程中已经计算好的值进行存储，变量是在运行中分配的值
- 常量支持的类型，bool、string、 数字
*/
func main() {

	// 常量的定义采用const
	const PI = 3.1415926
	fmt.Println(PI)
	//  Cannot assign
	// PI = 3.14

	// 常量如果定义了，没有使用是不会报错的
	const TEST = 999
	const IS = true
	const STRING = "A"
	const NAME = "TACKS"
	fmt.Println(IS, STRING, NAME)

}
